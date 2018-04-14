package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxLoadEntries ... pulls all entries for a given array of entry ids and fprints it back as a json array
func AjaxLoadEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	entryIDsString := r.FormValue("postIndex")
	entryIDsArr := uses.ConvertStrToStrArr(entryIDsString)

	fmt.Println("debug text ajaxloadentries line 25")
	fmt.Println(entryIDsArr)
	fmt.Println(len(entryIDsArr[0]))

	entries, err := uses.LoadEntries(client.Eclient, entryIDsArr)
	if err != nil {
		fmt.Println("err middleware/profile/ajaxloadentries line 30")
		fmt.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("err middleware/profile/ajaxloadentries line 37")
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
