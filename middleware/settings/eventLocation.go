package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/event"
	uses "github.com/sea350/ustart_go/uses"
)

//EventLocation ...
//For Events Location
func EventLocation(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	street := r.FormValue("street")
	zip := r.FormValue("zip")
	//   fmt.Println(blob)

	evnt, err := get.EventByID(eclient, r.FormValue("eventID"))
	//fmt.Println(reflect.TypeOf(blob))
	//TODO: DocID
	err = uses.ChangeEventLocation(eclient, r.FormValue("eventID"), country, state, city, street, zip)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/EventSettings/"+evnt.URLName, http.StatusFound)
	return

}
