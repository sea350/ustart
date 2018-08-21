package project

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//ProjectsPage ... Displays the projects page
func ProjectsPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	fmt.Println("project", test1)

	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	project, err := uses.AggregateProjectData(client.Eclient, r.URL.Path[10:], test1.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	widgets, errs := uses.LoadWidgets(client.Eclient, project.ProjectData.Widgets)
	if len(errs) > 0 {
		log.Println("there were one or more errors loading widgets")
		for _, eror := range errs {
			fmt.Println(eror)
		}
	}
	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), Project: project, Widgets: widgets}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "projectsF", cs)
}

//MyProjects ... ManageProject
func MyProjects(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	var heads []types.FloatingHead

	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	for _, projectInfo := range userstruct.Projects {
		head, err := uses.ConvertProjectToFloatingHead(client.Eclient, projectInfo.ProjectID)
		if err != nil {
			panic(err)
		}
		heads = append(heads, head)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string), ListOfHeads: heads}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "manageprojects-Nil", cs)
}

//CreateProjectPage ...
func CreateProjectPage(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	userstruct, err := get.UserByID(client.Eclient, session.Values["DocID"].(string))
	if err != nil {
		panic(err)
	}
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string), Username: session.Values["Username"].(string)}

	title := r.FormValue("project_title")
	description := []rune(r.FormValue("project_desc"))
	category := r.FormValue("category")
	college := r.FormValue("universityName")
	customURL := r.FormValue("curl")

	if title != `` {
		//proper URL
		if !uses.ValidUsername(customURL) {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println("Invalid custom project URL")
			cs.ErrorStatus = true
			cs.ErrorOutput = errors.New("Invalid custom project URL")
			client.RenderSidebar(w, r, "template2-nil")
			client.RenderSidebar(w, r, "leftnav-nil")
			client.RenderTemplate(w, r, "createProject-Nil", cs)
			return

		}
		url, err := uses.CreateProject(client.Eclient, title, description, session.Values["DocID"].(string), category, college, customURL)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
		} else {
			http.Redirect(w, r, "/Projects/"+url, http.StatusFound)
			return
		}
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "createProject-Nil", cs)
}
