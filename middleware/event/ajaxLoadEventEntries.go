package event

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/event"
	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadEventEntries ... pulls all entries for a given event and fprints it back as a json array
func AjaxLoadEventEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	wallID := r.FormValue("userID")
	evnt, err := get.EventByID(client.Eclient, wallID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	/*
		entries, err := uses.LoadEntries(client.Eclient, evnt.EntryIDs)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, err)
		}
	*/

	res, entries, err := scrollpkg.ScrollPageEvents(client.Eclient, evnt.EntryIDs, "")
	if err != nil {
		fmt.Println(res)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	fmt.Fprintln(w, string(data))
}
