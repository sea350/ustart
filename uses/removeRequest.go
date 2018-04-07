package uses

import (
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//RemoveRequest ...
func RemoveRequest(eclient *elastic.Client, projectID string, userID string) (error, int) {
	var newNumRequests int
	err := userPost.DeleteSentProjReq(eclient, userID, projectID)
	if err != nil {
		return err, newNumRequests
	}
	err, newNumRequests = projPost.DeleteMemberReqReceived(eclient, projectID, userID)
	return err, newNumRequests
}
