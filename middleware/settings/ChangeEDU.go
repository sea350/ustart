package settings

import (
	"fmt"
	"html"
	
	"net/http"
	
	"strconv"

	"github.com/microcosm-cc/bluemonday"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Major ...
type Major struct {
	List []string
}

//ChangeEDU ...
func ChangeEDU(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	typeAcc := p.Sanitize(r.FormValue("type_select"))
	i, err2 := strconv.Atoi(typeAcc)
	if err2 != nil {
		fmt.Println(err2)
	}

	highschoolName := p.Sanitize(r.FormValue("schoolname"))
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"HS:", highschoolName)
	highschoolGrad := p.Sanitize(r.FormValue("highSchoolGradDate"))
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"HS Grad Date:", highschoolGrad)
	// uniName := p.Sanitize(r.FormValue("universityName"))
	uniName := r.FormValue("universityName")
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"UNI:", uniName)
	var major []string

	var m Major

	for i := range m.List {
		m.List[i] = p.Sanitize(m.List[i])
		m.List[i] = html.EscapeString(m.List[i])
	}

	major = append(major, r.FormValue("majors"))
	//	Year := r.FormValue("year")
	gradDate := p.Sanitize(r.FormValue("uniGradDate"))
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"UNI Grad Date:", gradDate)

	var minor []string

	err := uses.ChangeEducation(client.Eclient, session.Values["DocID"].(string), i, highschoolName, highschoolGrad, uniName, gradDate, major, minor)
	if err != nil {
		

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	http.Redirect(w, r, "/Settings/#educollapse", http.StatusFound)
	return
}
