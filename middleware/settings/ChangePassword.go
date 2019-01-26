package settings

import (
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangePassword ...
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()
	oldp := r.FormValue("oldpass")
	newp := r.FormValue("confirmpass")

	if oldp == `` && newp == `` {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Critical data not passed in")
		return
	}
	oldpb := []byte(oldp)
	newpb := []byte(newp)
	err := uses.ChangePassword(client.Eclient, session.Values["DocID"].(string), oldpb, newpb)
	if err != nil {
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err, "| ABOUT TO RETURN")
		return
	}
	client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err, "| BROKE OUT OF RETURN ")

	http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	return

}
