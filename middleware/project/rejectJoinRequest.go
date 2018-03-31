package project

import (
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//RejectJoinRequest ...
func RejectJoinRequest(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	projID := r.FormValue("projectID")
	newMemberID := r.FormValue("userID")

	err := uses.RemoveRequest(client.Eclient, projID, newMemberID)
	if err != nil {
		fmt.Println("err middleware/project/acceptjoinrequest line 27")
		fmt.Println(err)
	}

	fmt.Fprintln(w, projID)
}
