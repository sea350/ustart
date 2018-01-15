package main

import (
	"html/template"
	"net/http"
	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
//	types "github.com/sea350/ustart_go/types"
	"fmt"
	"github.com/gorilla/sessions"
	"time"
//	"golang.org/x/crypto/bcrypt"
)


type ClientSide struct {
	//Session SessionInfo
	FirstName string
    LastName string
	
}



var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var templates = template.Must(template.ParseFiles("home/rr2396/www/ustart.tech/new-reg-nil.html",
	 "home/rr2396/www/ustart.tech/payment-nil.html",
	 "home/rr2396/www/ustart.tech/templateNoUser2.html",
	 "home/rr2396/www/ustart.tech/profile-nil.html",
	 "home/rr2396www/ustart.tech/template2-nil.html",
	 "home/rr2396/www/ustart.tech/template-footer-nil.html",
	 "home/rr2396/www/ustart.tech/nil-index2.html"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code 

func LoggedIn (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	
	fmt.Println(session.Values["username"].(string))
	cs := &ClientSide{FirstName:session.Values["username"].(string)}
	
	session.Save(r, w)
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"profile-nil",cs)


}


func Signup (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	
	store.MaxAge( 8640 * 7)
	test1, _ := session.Values["username"]
	cs := &ClientSide{}
     if (test1 != nil){
     //	fmt.Println(test1)
   //  	cs = &ClientSide{FirstName:session.Values["username"].(string)}
        http.Redirect(w, r, "/profile/", http.StatusFound)
        }

     session.Save(r, w)

	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"new-reg-nil",cs)

}


func Paypal(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["username"]
     if (test1 != nil){ http.Redirect(w, r, "/profile/", http.StatusFound) }
	fname := r.FormValue("firstName")
	lname := r.FormValue("lastName")
	email := r.FormValue("inputEmail")
//	var password []byte
	password := r.FormValue("inputPassword")
//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	school := r.FormValue("universityName")
	var major []string
	major = append(major,r.FormValue("majors"))
	bday := r.FormValue("dob")
	country := r.FormValue("country")
	state := r.FormValue("state")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	currYear := r.FormValue("year")
	if (err != nil){ fmt.Println(err) } // associated with eclient 
	err2 :=  uses.SignUpBasic(eclient, email, password, fname,lname, country, state, city, zip, school, major, bday, currYear)
 	if (err2 != nil){ fmt.Println(err2) }

    expiration := time.Now().Add((30) * time.Hour)
    cookie := http.Cookie{Name: fname, Value: "user", Expires: expiration, Path:"/"}
    http.SetCookie(w, &cookie)

    session.Values["username"] = r.FormValue("firstName")
    cs := &ClientSide{FirstName:session.Values["username"].(string)}
    fmt.Println("we are on payment here ")
    session.Save(r, w)
    renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"profile-nil",cs)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
//	session, _ := store.Get(r, "session_please")
//	test1, _ := session.Values["username"]
  //   if (test1 != nil){
     //	fmt.Println(test1)
  //   http.Redirect(w, r, "/profile/", http.StatusFound)
   //     }
//	u.FirstName = r.FormValue("firstName")
	email := r.FormValue("inputEmail")
//	var password []byte
	password := r.FormValue("inputPassword")
//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	successful,_,err2 :=  uses.Login(eclient, email, password)
 	if (err2 != nil){
	//	fmt.Println(err2)
	
	}

	if (successful == true){
		fmt.Println("login successful")
		http.Redirect(w, r, "/signup/", http.StatusFound)
	}
	if (successful == false){
		fmt.Println("did not login successful")
		http.Redirect(w, r, "/signup/", http.StatusFound)
	}
	
	


// renderTemplate takes in the ResponseWriter, a string referring to the file name, and a struct containing data we want to pass onto pages
func renderTemplate(w http.ResponseWriter, tmpl string, cs *ClientSide) {
  	err := templates.ExecuteTemplate(w, tmpl+".html", cs)
  	if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError) }
  }

