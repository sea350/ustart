package settings

import (
	"log"
	"net/http"
	"time"

	get "github.com/sea350/ustart_go/get/project"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectBannerUpload ... pushes a new banner image into ES
func ProjectBannerUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-banner")

	//Get the project and member
	proj, member, err1 := get.ProjAndMember(client.Eclient, r.FormValue("projectID"), test1.(string))
	if err1 != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err, "Project or Member not found")
		http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
		return
	}

	switch err {
	case nil:
		blob := r.FormValue("banner-data")
		if uses.HasPrivilege("banner", proj.PrivilegeProfiles, member) {
			buffer := make([]byte, 512)
			_, _ = clientFile.Read(buffer)
			defer clientFile.Close()
			if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
				//Update the project banner
				url, err := uses.UploadToS3(blob, r.FormValue("projectID")+"-banner")
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
				}
				err = post.UpdateProject(client.Eclient, r.FormValue("projectID"), "Banner", url)
				if err != nil {
					log.SetFlags(log.LstdFlags | log.Lshortfile)
					log.Println(err)
				}
			} else {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println("Invalid file upload")
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("You do not have permission to change event banner")
		}
	case http.ErrMissingFile:
		blob := r.FormValue("banner-data")
		if uses.HasPrivilege("banner", proj.PrivilegeProfiles, member) {
			//Update the project banner
			url, err := uses.UploadToS3(blob, r.FormValue("projectID")+"-banner")
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			}
			err = post.UpdateProject(client.Eclient, r.FormValue("projectID"), "Banner", url)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
			} else {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err, "Invalid file upload")
			}
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err, "You do not have permission to change event banner")
		}

	default:
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	time.Sleep(2 * time.Second)
	http.Redirect(w, r, "/ProjectSettings/"+proj.URLName, http.StatusFound)
}
