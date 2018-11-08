package settings

import (
	"log"
	"net/http"
	"strings"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	uses "github.com/sea350/ustart_go/uses"
)

//LeaveProject ... lets a user leave a project
//If Rol
func LeaveProject(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	leavingUser := r.FormValue("leaverID")
	leavingUser = strings.Trim(leavingUser, "/")

	projID := r.FormValue("projectID")
	if projID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Project ID not passed")
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}
	newCreator := r.FormValue("newCreator")

	proj, err := get.ProjectByID(client.Eclient, projID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	if leavingUser == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Leaver not specified")
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	var canLeave = false
	if leavingUser == test1.(string) {
		//if the current active user wants to leave, they can
		canLeave = true
	} else {
		for _, mem := range proj.Members {
			if mem.MemberID == test1.(string) && mem.Role == 0 {
				//if the current acessing user is creator, they can do whatever they want
				canLeave = true
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println("checkpoint 2b")
				break
			}
		}
	}
	if !canLeave {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("User attempting to leave was not permitted, check variables and try again.")
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	if newCreator == `` {
		err = post.DeleteMember(client.Eclient, projID, leavingUser)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	} else {
		err = uses.NewProjectLeader(client.Eclient, projID, leavingUser, newCreator)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		err = post.DeleteMember(client.Eclient, projID, leavingUser)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return

}
