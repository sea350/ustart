package settings

import (
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
)

//ImageUpload ...
func ImageUpload(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	blob := r.FormValue("image-data")

	err := uses.ChangeAccountImagesAndStatus(eclient, session.Values["DocID"].(string), blob, true, "hello", "Avatar")
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/Settings/#avatarcollapse", http.StatusFound)

}
