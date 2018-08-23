package notification

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/notification"
	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/uses"
)

//AjaxNotificationLoad ... crawling in the 90s
//Designed for ajax
func AjaxNotificationLoad(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		return
	}

	var notifs []map[string]interface{}

	proxy, err := get.ProxyNotificationByUserID(client.Eclient, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	count := 0
	for _, id := range proxy.NotificationCache {
		notif, err := get.NotificationByID(client.Eclient, id)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}
		if notif.Invisible {
			continue
		}

		msg, url, err := uses.GenerateNotifMsgAndLink(client.Eclient, notif)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}

		notifAggregate := make(map[string]interface{})
		notifAggregate["Data"] = notif
		notifAggregate["Message"] = msg
		notifAggregate["URL"] = URL
		notifs = append(notifs, notifAggregate)
		count++
		if count == 5 {
			break
		}

	}

	sendData := make(map[string]interface{})
	sendData["notifications"] = notifs

	data, err := json.Marshal(sendData)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
	fmt.Fprintln(w, string(data))
}
