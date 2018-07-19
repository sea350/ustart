package chat

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

var chatClients = make(map[string](*websocket.Conn))
var chatNotifRadio = make(chan types.FloatingHead)

//HandleChatClients ...
func HandleChatClients(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	chatURL := r.URL.Path[4:]

	//security checks before socket is opened
	valid, actualChatID, dmTargetUserID, err := uses.ChatVerifyURL(client.Eclient, chatURL, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	if !valid {
		return
	}

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	_, exists := chatroom[chatURL]
	if !exists {
		temp := make(map[*websocket.Conn]bool)
		temp[ws] = true
		chatroom[chatURL] = temp
	} else {
		temp := chatroom[chatURL]
		temp[ws] = true
		chatroom[chatURL] = temp
	}

	for {
		var msg types.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		msg = types.Message{SenderID: docID.(string), TimeStamp: time.Now(), Content: msg.Content, ConversationID: actualChatID}
		if actualChatID == `` {
			newConvoID, err := uses.ChatFirst(client.Eclient, msg, docID.(string), dmTargetUserID)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
			}
			actualChatID = newConvoID
		} else {
			_, err = uses.ChatSend(client.Eclient, msg)
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
			}
		}

		//send notification here

		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleChatAlert() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		/*
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			dir, _ := os.Getwd()
			log.Println(dir, "debug text")
			log.Println("channel #" + msg.ChatID)
			log.Printf("message: %v \n", msg)
			log.Println(chatroom[msg.ChatID])
		*/
		for client := range chatroom[msg.ConversationID] {
			err := client.WriteJSON(msg)
			if err != nil {
				//log.Printf("error: %v", err)
				client.Close()
				delete(chatroom[msg.ConversationID], client)
			}
		}
	}
}
