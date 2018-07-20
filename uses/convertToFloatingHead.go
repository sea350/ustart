package uses

import (
	getChat "github.com/sea350/ustart_go/get/chat"
	getEvent "github.com/sea350/ustart_go/get/event"
	getProject "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ConvertUserToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertUserToFloatingHead(eclient *elastic.Client, userDocID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	usr, err := getUser.UserByID(eclient, userDocID)
	if err != nil {
		return head, err
	}

	head.FirstName = usr.FirstName
	head.LastName = usr.LastName
	head.Image = usr.Avatar
	head.Username = usr.Username
	head.DocID = userDocID

	head.Interface = usr.Tags
	head.Bio = usr.Description

	return head, err
}

//ConvertProjectToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertProjectToFloatingHead(eclient *elastic.Client, projectID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	proj, err := getProject.ProjectByID(eclient, projectID)
	if err != nil {
		panic(err)
	}

	head.FirstName = proj.Name
	head.Bio = proj.Description
	head.Image = proj.Avatar
	head.Username = proj.URLName
	head.DocID = projectID
	head.Notifications = len(proj.MemberReqReceived)
	head.Interface = proj.Tags

	return head, err
}

//ConvertEventToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertEventToFloatingHead(eclient *elastic.Client, eventID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	evnt, err := getEvent.EventByID(eclient, eventID)
	if err != nil {
		panic(err)
	}

	head.FirstName = evnt.Name
	head.Bio = evnt.Description
	head.Image = evnt.Avatar
	head.Username = evnt.URLName
	head.DocID = eventID
	head.Notifications = len(evnt.MemberReqReceived)
	head.Interface = evnt.Tags

	return head, err
}

//ConvertChatToFloatingHead ... pulls latest version of chat and converts relevent data into floating head
func ConvertChatToFloatingHead(eclient *elastic.Client, conversationID string, viewerID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	convo, err := getChat.ConvoByID(eclient, conversationID)
	if err != nil {
		return head, err
	}

	if convo.Class == 3 {
		head, err = ConvertProjectToFloatingHead(eclient, convo.ReferenceID)
		if err != nil {
			return head, err
		}
	}

	if convo.Class == 4 {
		head, err = ConvertEventToFloatingHead(eclient, convo.ReferenceID)
		if err != nil {
			return head, err
		}
	}

	if len(convo.MessageIDCache) > 0 {
		msg, err := getChat.MsgByID(eclient, convo.MessageIDCache[len(convo.MessageIDCache)-1])
		head.Bio = []rune(msg.Content)
		if err != nil {
			return head, err
		}
	}

	for key, eaver := range convo.Eavesdroppers {
		if key == viewerID {
			head.Notifications = len(convo.MessageIDArchive) - eaver.Bookmark - 1
		} else if convo.Class == 1 {
			usr, err := getUser.UserByID(client.Eclient, key)
			if err != nil {
				return head, err
			}
			head.Username = usr.Username
		}
	}

	head.DocID = conversationID

	return head, err
}
