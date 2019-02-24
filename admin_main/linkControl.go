package main

import (
	getUser "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"

	// post "github.com/sea350/ustart_go/post/user"

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	"github.com/sea350/ustart_go/globals"
	elastic "github.com/olivere/elastic"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	usr, err := getUser.UserByUsername(eclient, "support")
	usrID, err := getUser.IDByUsername(eclient, "support")
	// usr2, err := getUser.UserByUsername(eclient, "min")
	usrID2, err := getUser.IDByUsername(eclient, "min")

	fmt.Println(len(usr.QuickLinks))
	var emp []types.Link{}
	err = post.UpdateUser(eclient, usrID, "QuickLinks", emp)
	if err != nil {
		fmt.Println("LINE 24,", err)
	}

	err = post.UpdateUser(eclient, usrID2, "QuickLinks", emp)
	if err != nil {
		fmt.Println("LINE 24,", err)
	}
}
