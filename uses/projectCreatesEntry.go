package uses

import (
	"time"

	postEntry "github.com/sea350/ustart_go/post/entry"
	post "github.com/sea350/ustart_go/post/project"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectCreatesEntry ... creates a new entry for projects and handles logic/parallel arrays
func ProjectCreatesEntry(eclient *elastic.Client, projID string, posterID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 0
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true

	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, projID, entryID)

	return err

}

//ProjectCreatesReply ... creates a new reply entry for projects and handles logic/parallel arrays
func ProjectCreatesReply(eclient *elastic.Client, projID string, replyID string, posterID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 1
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true

	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, projID, entryID)
	if err != nil {
		return err
	}

	err = postEntry.AppendReplyID(eclient, entryID, replyID)

	return err

}

//ProjectCreatesShare ... creates a new share entry for projects and handles logic/parallel arrays
func ProjectCreatesShare(eclient *elastic.Client, projID string, replyID string, posterID string, newContent []rune) error {
	createdEntry := types.Entry{}
	createdEntry.PosterID = posterID
	createdEntry.Classification = 2
	createdEntry.Content = newContent
	createdEntry.TimeStamp = time.Now()
	createdEntry.Visible = true

	//usr, err := get.GetUserByID(eclient,userID)

	entryID, err := postEntry.IndexEntry(eclient, createdEntry)
	if err != nil {
		return err
	}

	err = post.AppendEntryID(eclient, projID, entryID)
	if err != nil {
		return err
	}

	err = postEntry.AppendShareID(eclient, entryID, replyID)

	return err

}
