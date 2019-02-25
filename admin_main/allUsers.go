package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

//Jv63yWgBN3Vvtvdiu5YP

func main() {

	ctx := context.Background()

	maq := elastic.NewMatchAllQuery()
	res, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(500).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, id := range res.Hits.Hits {
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data.FirstName, "  ", data.LastName, "  ", data.Username)
	}

}

// func main() {

// 	usr, _ := getUser.UserByID(eclient, "Jv63yWgBN3Vvtvdiu5YP")
// 	fmt.Println(usr.FirstName, usr.LastName)

// }
