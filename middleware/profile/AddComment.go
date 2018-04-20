package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//AddComment ... Iunno
func AddComment(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("followstat")
	postActual := postID[1:]
	comment := r.FormValue("commentz")
	id := r.FormValue("id") // userID
	fmt.Println("ID IS " + id)
	contentArray := []rune(comment)
	//username := r.FormValue("username")
	fmt.Println(postActual + "is the post ID? ")
	err4 := uses.UserReplyEntry(client.Eclient, id, postActual, contentArray)
	if err4 != nil {
		fmt.Println(err4)
	}

	_, cmts, err := uses.LoadComments(client.Eclient, postID, 0, -1)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(cmts)
	fmt.Println("DATA NEXT:", string(data))
	fmt.Fprintln(w, string(data))

	//http.Redirect(w, r, "/profile/"+username, http.StatusFound)

}

//AddComment2 ... Iunno
func AddComment2(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	fmt.Println("WE ARE IN ADDCOMMENT.GO")
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("postID")
	fmt.Println(postID)
	// postActual := postID[1:]
	comment := r.FormValue("body")
	contentArray := []rune(comment)
	fmt.Println(session.Values["DocID"].(string) + "IS PIKA")
	err4 := uses.UserReplyEntry(client.Eclient, session.Values["DocID"].(string), postID, contentArray)
	if err4 != nil {
		fmt.Println(err4)
	}

	_, cmts, err := uses.LoadComments(client.Eclient, postID, 0, -1)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(cmts)

	fmt.Fprintln(w, string(data))

	//http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	//return
}
