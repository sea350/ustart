package main

import (
	"html/template"
	"net/http"
	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"fmt"
	"github.com/gorilla/sessions"
	"time"
	"golang.org/x/crypto/bcrypt"
//	"github.com/gorilla/mux"
//	"log"
	//types"github.com/sea350/ustart_go/types"
)


type ClientSide struct {
	DOCID string 
	FirstName string
    LastName string
    Username string
    ErrorR bool 
    ErrorLogin bool 
    UserInfo types.User
    Class string 
}



var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var templates = template.Must(template.ParseFiles("../../../../www/ustart.tech/new-reg-nil.html","../../../../www/ustart.tech/templateNoUser2.html","../../../../www/ustart.tech/regcomplete-nil.html"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code 


func RegistrationComplete (w http.ResponseWriter, r *http.Request){
	cs := ClientSide{}
	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"regcomplete-nil",cs)

}	


func Signup (w http.ResponseWriter, r *http.Request){
	store.MaxAge( 8640 * 7)
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	var errorreg bool
	errorreg = false
	
     if (test1 != nil){
     //	fmt.Println(test1)
   //  	cs = &ClientSide{FirstName:session.Values["username"].(string)}
        http.Redirect(w, r, "/profile/"+test1.(string), http.StatusFound)
        }

     session.Save(r, w)
	cs := ClientSide{ErrorR: errorreg, ErrorLogin:errorreg}
	//fmt.Println("hello buddy signup")
	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"new-reg-nil",cs)

}


func Registration(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session, _ := store.Get(r, "session_please")
	// check DOCID instead 
	test1, _ := session.Values["DocID"]
     if (test1 != nil){
     //	fmt.Println(test1)
     	// REGISTRATION SHOULD NOT LOG YOU IN 
     http.Redirect(w, r, "/profile/", http.StatusFound)
        }
//	u.FirstName = r.FormValue("firstName")
	fname := r.FormValue("firstName")
	lname := r.FormValue("lastName")
	email := r.FormValue("inputEmail")

	password := r.FormValue("inputPassword")
	passwordb := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordb, bcrypt.DefaultCost)
	school := r.FormValue("universityName")
	var major []string
	major = append(major,r.FormValue("majors"))
	bday := time.Now()//r.FormValue("dob")
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	currYear := r.FormValue("year")
	if (err != nil){
		fmt.Println(err)

		
	}
	err2 :=  uses.SignUpBasic(eclient, email, hashedPassword, fname,lname, country, state, city, zip, school, major, bday, currYear)
 	if (err2 != nil){
		fmt.Println(err2)
				cs := ClientSide{ErrorR:true}
	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"new-reg-nil",cs)

	
	}

	if (err2 == nil){
		      http.Redirect(w, r, "/registrationcomplete/", http.StatusFound)
	}

}


func renderTemplate(w http.ResponseWriter, tmpl string, cs ClientSide) {
//  	fmt.Println("rT called")
  	err := templates.ExecuteTemplate(w, tmpl+".html", cs)
  	if err != nil {
  		http.Error(w, err.Error(), http.StatusInternalServerError)
  	}
  }


