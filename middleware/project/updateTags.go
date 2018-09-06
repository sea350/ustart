package project

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/project"
)

type TagStruct struct {
	Tags []string
}

//UpdateTags ...
func UpdateTags(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	p := bluemonday.UGCPolicy()

	ID := r.FormValue("skillArray")
	log.Println("TAGS:", ID)
	var ts TagStruct
	err := json.Unmarshal([]byte(ID), &ts.Tags)

	if err != nil {
		log.Println("Could not unmarshal")
		return
	}

	for t := range ts.Tags {
		ts.Tags[t] = p.Sanitize(ts.Tags[t])
		ts.Tags[t] = html.EscapeString(ts.Tags[t])
	}

	log.Println("TS:", ts.Tags)
	err = post.UpdateProject(client.Eclient, ID, "Tags", ts.Tags)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
