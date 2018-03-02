package main

import (
	"net/http"

	inbox "github.com/sea350/ustart_go/middleware/inbox"
	login "github.com/sea350/ustart_go/middleware/login"
	profile "github.com/sea350/ustart_go/middleware/profile"
	project "github.com/sea350/ustart_go/middleware/project"
	registration "github.com/sea350/ustart_go/middleware/registration"
	settings "github.com/sea350/ustart_go/middleware/settings"
	widget "github.com/sea350/ustart_go/middleware/widget"
)

func main() {
	/*
		Lines 18-19 handle the static file locating
		If we wanted to reorganize file/folder locations, this is one of 3 things that would have to change
		In executeTemplates you will need to make the same changes
		The other being the relative link on the actual html pages
	*/
	// fs := http.FileServer(http.Dir("/home/rr2396/www/"))
	fs := http.FileServer(http.Dir("/ustart/ustart_front/"))
	// http.Handle("/www/", http.StripPrefix("/www/", fs))
	http.Handle("/ustart_front/", http.StripPrefix("/ustart_front/", fs))
	/*
		The following are all the handlers we have so far.
	*/

	http.HandleFunc("/Inbox/", inbox.Inbox)
	http.HandleFunc("/loginerror/", login.Error)
	http.HandleFunc("/", login.Home)
	http.HandleFunc("/profilelogin/", login.Login)
	http.HandleFunc("/logout/", login.Logout)

	http.HandleFunc("/callme/", profile.Follow)
	http.HandleFunc("/Like", profile.Like)
	http.HandleFunc("/getComments/", profile.GetComments)
	http.HandleFunc("/shareComments/", profile.ShareComments)
	http.HandleFunc("/ShareComment", profile.ShareComment2)
	http.HandleFunc("/AddComment", profile.AddComment)
	http.HandleFunc("/loadWall/", profile.WallLoad)
	http.HandleFunc("/addPost/", profile.WallAdd)
	http.HandleFunc("/profile/", profile.ViewProfile)
	http.HandleFunc("/addSkill/", profile.AddTag)
	http.HandleFunc("/deleteSkill/", profile.DeleteTag)
	http.HandleFunc("/addLink/", profile.AddQuickLink)
	http.HandleFunc("/deleteLink/", profile.DeleteQuickLink)

	http.HandleFunc("/addWidget/", widget.AddWidget)
	http.HandleFunc("/deleteWidget/", widget.DeleteWidgetProfile)
	http.HandleFunc("/deleteLinkFromWidget/", widget.EditWidgetDataDelete)

	http.HandleFunc("/Projects/", project.ProjectsPage)
	http.HandleFunc("/MyProjects/", project.MyProjects)
	http.HandleFunc("/CreateProject/", project.CreateProject)
	http.HandleFunc("/Settings/", settings.Settings)
	http.HandleFunc("/ImageUpload/", settings.ImageUpload)
	http.HandleFunc("/changeName/", settings.ChangeName)
	http.HandleFunc("/changePass/", settings.ChangePassword)
	http.HandleFunc("/changeLoc/", settings.ChangeLocation)
	http.HandleFunc("/changeEDU/", settings.ChangeEDU)
	http.HandleFunc("/UpdateDescription/", settings.ChangeContactAndDescription)

	http.HandleFunc("/Signup/", registration.Signup)
	http.HandleFunc("/Registration/Type/", registration.RegisterType)
	http.HandleFunc("/registrationcomplete/", registration.Complete)
	http.HandleFunc("/welcome/", registration.Registration)
	http.ListenAndServe(":5000", nil)
}
