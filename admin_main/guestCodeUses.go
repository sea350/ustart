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

func main() {

	ctx := context.Background()

	maq := maq.Must(elastic.NewMatchQuery("FirstName", "joaquin"))

	res, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(500).
		// Sort("_score", true).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, id := range res.Hits.Hits {
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		fmt.Println(data.Email+","+data.FirstName+","+data.LastName, +",    "+id.Id)
		guestID := id.Id
		break
		if err != nil {
			fmt.Println(err)

		}
	}

	maq = maq.Must(elastic.NewMatchQuery("_id", guestID))

	res, err = eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(500).
		// Sort("_score", true).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, id := range res.Hits.Hits {
		data := types.GuestCode{}
		err = json.Unmarshal(*id.Source, &data)
		fmt.Println(data.Codel + "," + data.Users + "," + data.Expiration)
		if err != nil {
			fmt.Println(err)

		}
	}

}
