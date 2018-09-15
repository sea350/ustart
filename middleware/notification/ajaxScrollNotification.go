package notification

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sea350/ustart_go/middleware/client"
	properloading "github.com/sea350/ustart_go/properloading"
	"github.com/sea350/ustart_go/uses"
)

//AjaxScrollNotification ...
//Scrolls notifications
func AjaxScrollNotification(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	scrollID := r.FormValue("scrollID")
	sID, notifMap, hits, err := properloading.ScrollNotifications(client.Eclient, docID.(string), scrollID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	var notifs []map[string]interface{}

	for notifID, notif := range notifMap {

		msg, url, err := uses.GenerateNotifMsgAndLink(client.Eclient, notif)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}

		notifAggregate := make(map[string]interface{})
		notifAggregate["ID"] = notifID
		notifAggregate["Data"] = notif
		notifAggregate["Message"] = msg
		notifAggregate["URL"] = url
		notifs = append(notifs, notifAggregate)
		// count++
		// if count == 5 {
		// 	break
		// }

	}

	sendData := make(map[string]interface{})
	sendData["notifications"] = notifs
	// sendData["numUnread"] = proxy.NumUnread

	data, err := json.Marshal(sendData)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(data))
}
