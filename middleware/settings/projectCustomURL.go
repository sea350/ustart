package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	uses "github.com/sea350/ustart_go/uses"
)

//ProjectCustomURL ... pushes a new banner image into ES
func ProjectCustomURL(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	newURL := r.FormValue("purl")

	projID := r.FormValue("projectID")
	proj, err := get.ProjectByID(eclient, projID)
	if err != nil {
		fmt.Println(err)
	}

	err = uses.ChangeProjectURL(eclient, projID, newURL)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)

}
