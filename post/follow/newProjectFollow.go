package post

import (
	"context"
	"errors"
	"fmt"
	"strings"

	getFollow "github.com/sea350/ustart_go/get/follow"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewProjectFollow ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
//Field can be Followers or Following
func NewProjectFollow(eclient *elastic.Client, projID string, field string, newKey string, isBell bool) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.FollowIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	follID, foll, err := getFollow.ByID(eclient, projID)
	if err != nil {
		return err
	}

	//  vFollowerLock.Lock()

	var followMap = make(map[string]bool)
	var bellMap = make(map[string]bool)
	switch strings.ToLower(field) {
	case "followers":
		FollowerLock.Lock()
		defer FollowerLock.Unlock()
		if len(foll.UserFollowers) == 0 {
			var newMap = make(map[string]bool)
			newMap[newKey] = isBell
			followMap = newMap
			if isBell {
				var newBell = make(map[string]bool)
				newBell[newKey] = isBell
				bellMap = newBell
			}
		} else {
			foll.UserFollowers[newKey] = isBell
			followMap = foll.ProjectFollowers

			//modify user bell map if bell follower
			if isBell {
				foll.ProjectBell[newKey] = isBell
				bellMap = foll.ProjectBell
			}
		}

	case "following":
		FollowingLock.Lock()
		defer FollowingLock.Unlock()
		if len(foll.ProjectFollowing) == 0 {
			var newMap = make(map[string]bool)
			newMap[newKey] = isBell
			followMap = newMap
		} else {
			foll.ProjectFollowing[newKey] = isBell
			followMap = foll.ProjectFollowing
		}
	default:
		return errors.New("Invalid field")
	}
	var theField string
	if strings.ToLower(field) == "followers" {
		theField = "UserFollowers"
	} else if strings.ToLower(field) == "following" {
		theField = "UserFollowing"
	}

	fmt.Println("THE FIELD:", theField)
	newFollow := eclient.Update().
		Index(globals.FollowIndex).
		Type(globals.FollowType).
		Id(follID).
		Doc(map[string]interface{}{theField: followMap}) //field = Followers or Following, newContent =

	//only executes when there is a new bell follower
	if isBell && strings.ToLower(field) == "followers" {
		newFollow.Doc(map[string]interface{}{"UserBell": bellMap})
	}
	_, err = newFollow.Do(ctx)
	return err
}
