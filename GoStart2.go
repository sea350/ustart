package main

import (
	"html/template"
	"net/http"
	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	get "github.com/sea350/ustart_go/get"
	"fmt"
	"github.com/gorilla/sessions"
	"time"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var templates = template.Must(template.ParseFiles("../../../../www/ustart.tech/followerlist-nil.html","../../../../www/ustart.tech/emTee.html","../../../../www/ustart.tech/wallttt.html","../../../../www/ustart.tech/wallload-nil.html","../../../../www/ustart.tech/testimage.html","../../../../www/ustart.tech/ajax-nil.html","../../../../www/ustart.tech/Membership-Nil.html","../../../../www/ustart.tech/settings-Nil.html","../../../../www/ustart.tech/inbox-Nil.html","../../../../www/ustart.tech/createProject-Nil.html","../../../../www/ustart.tech/manageprojects-Nil.html","../../../../www/ustart.tech/projects-Nil.html","../../../../www/ustart.tech/new-reg-nil.html","../../../../www/ustart.tech/loginerror-nil.html","../../../../www/ustart.tech/test.html", "../../../../www/ustart.tech/payment-nil.html","../../../../www/ustart.tech/templateNoUser2.html","../../../../www/ustart.tech/profile-nil.html","../../../../www/ustart.tech/template2-nil.html","../../../../www/ustart.tech/template-footer-nil.html","../../../../www/ustart.tech/nil-index2.html","../../../../www/ustart.tech/regcomplete-nil.html"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code 

type ClientSide struct {
	DOCID string 
	FirstName string
    LastName string
    Username string
    ErrorR bool 
    ErrorLogin bool 
    UserInfo types.User
    Class string 
    Birthday string
    ImageCode string
    Description string
    Followers int
    Following int
    Page string
    FollowingStatus string 
    Wall []types.JournalEntry 
}
/*
func WallTest (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
//	fmt.Println(session.Values["FirstName"].(string))
	var WallEntry types.JournalEntry
//	var newEntry types.Entry 
//	var newJournal types.JournalEntry

	WallEntry.FirstName = "Nil"
	WallEntry.LastName = "Patel"
//	WallEntry.Element = newEntry
//	WallEntry.RepliesArray = _
	WallEntry.NumLikes = 0
	WallEntry.NumReplies = 2
	WallEntry.NumShares = 100000000000000
	fmt.Println(WallEntry)

	var WallEntries []types.JournalEntry
	WallEntries[1] = WallEntry
	cs := ClientSide{FirstName:session.Values["DocID"].(string),Wall:WallEntries}
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"wallttt",cs)
}
*/
func LoggedIn (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	//	fmt.Println(session.Values["FirstName"].(string))
	cs := ClientSide{FirstName:session.Values["FirstName"].(string)}
	session.Save(r, w)
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"profile-nil",cs)
}

func Home (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 != nil){
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound) }
	session.Save(r, w)
	cs := ClientSide{}
	//fmt.Println("helllo")
	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"nil-index2",cs)
	renderTemplate(w,"template-footer-nil",cs)
}

func Follow (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 != nil){
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound) }
	session.Save(r, w)
	cs := ClientSide{}
	//fmt.Println("helllo")
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"followerlist-nil",cs)

}

func ViewProfile (w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 == nil){
    	http.Redirect(w, r, "/~", http.StatusFound) }
	fmt.Println("------------------------")
	fmt.Println(r.URL.Path[9:])
	fmt.Println("SESSIONS ID IS "+session.Values["DocID"].(string))
	userstruct,_, followbool,_ := uses.UserPage(eclient,r.URL.Path[9:],session.Values["DocID"].(string))
	fmt.Println(userstruct.EntryIDs)
	jEntries, err5 := uses.LoadEntries(eclient,userstruct.EntryIDs)
	if (err5 != nil){
		fmt.Println(err5);
	}
	jEntries2 := jEntries 
//	counter := len(jEntries)-1;
	/*
	for i := range jEntries{
      //  fmt.Println(jEntries[len(jEntries)-1-i].Element.TimeStamp) // Suggestion: do last := len(s)-1 before the loop
        jEntries2[counter] = jEntries[len(jEntries)-1-i]
        counter--

}
*/


	
	fmt.Println(userstruct.FirstName)
	fmt.Println(userstruct.LastName)
	fmt.Println(userstruct.Username)
	followingState := "no"
	if (followbool == true){
		followingState = "yes"
		fmt.Println("is following "+followingState)	
	}
	if (followbool == false){
		fmt.Println("is not following "+followingState)
	}
	for i := 0; i < len(jEntries); i++ {
		fmt.Println(jEntries2[i].Element.TimeStamp)
	}

	var ClassYear string 
	if (userstruct.Class == 1){
		ClassYear = "Freshman"
	}
	if (userstruct.Class == 2){
		ClassYear = "Sophomore"
	}
	if (userstruct.Class == 3){
		ClassYear = "Junior"
	}
	if (userstruct.Class == 4){
		ClassYear = "Senior"
	}
	if (userstruct.Class == 5){
		ClassYear = "Graduate"
	}
	if (userstruct.Class == 6){
		ClassYear = "Post-Graduate"
	}
	bday := userstruct.Dob.String()
//	fmt.Println(bday)
	month := bday[5:7]
	day := bday[8:10]
	year := bday[0:4]
	fmt.Println(month)
	fmt.Println(day)
	fmt.Println(year)
	birthdayline := month+"/"+day+"/"+year
	cs := ClientSide{UserInfo:userstruct, DOCID: session.Values["DocID"].(string),Birthday: birthdayline,Class:ClassYear} //Class:ClassYear}
	//fmt.Println("email is "+cs.UserInfo)
	viewingDOC, errID := get.GetIDByUsername(eclient, r.URL.Path[9:])
	if (errID != nil){
		fmt.Println(errID);
	}
	fmt.Println("viewing "+viewingDOC)
	fmt.Println("description is "+string(cs.UserInfo.Description))
	temp := string(cs.UserInfo.Description) 

	numberFollowing,errnF := uses.NumFollow(eclient, session.Values["DocID"].(string),true)
	if (errnF != nil){
		fmt.Println(errnF);
	}
	numberFollowers,errnF2 := uses.NumFollow(eclient, session.Values["DocID"].(string),false)
	if (errnF2 != nil){
		fmt.Println(errnF2);
	}

		test123 := "hello"
		test1245 := []rune(test123)
		postactual := "AV7T7n8C22dVORxe2i9O"
		id := session.Values["DocID"].(string)
		fmt.Println(id+"is docid 1234")
		err4 := uses.UserNewReplyEntry(eclient,id,test1245,postactual)
		if (err4 != nil){
		fmt.Println(err4)
	}

	cs = ClientSide{UserInfo:userstruct, Wall: jEntries, DOCID: session.Values["DocID"].(string),Birthday: birthdayline,Class:ClassYear, Description:temp,Followers:numberFollowers,Following:numberFollowing, Page:viewingDOC,FollowingStatus:followingState}


	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"profile-nil",cs)
}


func RegistrationComplete (w http.ResponseWriter, r *http.Request){
	cs := ClientSide{}
	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"regcomplete-nil",cs)
}	


func RegisterType (w http.ResponseWriter, r *http.Request){
	cs := ClientSide{}
	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"Membership-Nil",cs)
}	



func Signup (w http.ResponseWriter, r *http.Request){
	store.MaxAge( 8640 * 7)
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	var errorreg bool
	errorreg = false
	if (test1 != nil){
		http.Redirect(w, r, "/profile/"+test1.(string), http.StatusFound)}
	session.Save(r, w)
	cs := ClientSide{ErrorR: errorreg, ErrorLogin:errorreg}
	renderTemplate(w,"templateNoUser2",cs)
	renderTemplate(w,"new-reg-nil",cs)
}

func Test (w http.ResponseWriter, r *http.Request){

	fmt.Println("hello buddy test")
	cs := ClientSide{}
	renderTemplate(w,"test",cs)

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
	email = strings.ToLower(email)

	password := r.FormValue("inputPassword")
	passwordb := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordb, bcrypt.DefaultCost)
	school := r.FormValue("universityName")
	var major []string
	major = append(major,r.FormValue("majors"))
	fmt.Println(r.FormValue("dob"))
	bday := time.Now()//r.FormValue("dob")
	month,_ := strconv.Atoi(r.FormValue("dob")[0:2])
	day,_ := strconv.Atoi(r.FormValue("dob")[3:5])
	year,_ := strconv.Atoi(r.FormValue("dob")[6:10])
	bday = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	fmt.Println(bday.Date())
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


	// if registration unsuccessful, return to new-reg-nil and make .RegistrationError == true

 //   expiration := time.Now().Add((30) * time.Hour)
 //   cookie := http.Cookie{Name: fname, Value: "user", Expires: expiration, Path:"/"}
 //   http.SetCookie(w, &cookie)

  //  session.Values["DOCID"] = r.FormValue("firstName")
  //  cs := &ClientSide{FirstName:session.Values["DocID"].}
  //  fmt.Println("we are on payment here ")
  //  session.Save(r, w)
 //   renderTemplate(w,"template2-nil",cs)
	//	renderTemplate(w,"profile-nil",cs)

	// <---     --->
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
     if (test1 != nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
       }
	//	u.FirstName = r.FormValue("firstName")
	email := r.FormValue("email")
	email = strings.ToLower(email)
	fmt.Println(email)
	//	var password []byte
	password := r.FormValue("password")
	fmt.Println(password)
	//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	passwordb := []byte(password)
	successful,sessionInfo,err2 :=  uses.Login(eclient, email, passwordb)

	// doc ID can be retrieved here! 
	//cs := &ClientSide{}
 	if (err2 != nil){
		fmt.Println(err2)
	
	}

	if (successful == true){
		fmt.Println("login successful")
		session.Values["DocID"] = sessionInfo.DocID
		session.Values["FirstName"] = sessionInfo.FirstName
		session.Values["LastName"] = sessionInfo.LastName
		session.Values["Email"] = sessionInfo.Email
		session.Values["Username"] = sessionInfo.Username 
    	expiration := time.Now().Add((30) * time.Hour)
    	fmt.Println("Doc id is "+sessionInfo.DocID)
    	cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", Expires: expiration, Path:"/"}
    	http.SetCookie(w, &cookie)
		session.Save(r,w)
    	http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)	
	}

	if (successful == false){
		fmt.Println("did not login successful")
		var errorL bool
		errorL = true
		cs := ClientSide{ErrorLogin: errorL}
		fmt.Println("errorL is ")
		fmt.Print(errorL)
		renderTemplate(w,"templateNoUser2",cs)
		renderTemplate(w,"loginerror-nil",cs)
		
		

		
	}
}

func ProjectsPage(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
     if (test1 == nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/~", http.StatusFound)
       }
    cs := ClientSide{} 
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"projects-Nil",cs)
}



func MyProjects(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 == nil){
		http.Redirect(w, r, "/~", http.StatusFound) }
	userstruct, _, _,_ := uses.UserPage(eclient,session.Values["Username"].(string),session.Values["DocID"].(string))
	cs := ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)} 	
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"manageprojects-Nil",cs)
}

/*
func Follow(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
     if (test1 == nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/~", http.StatusFound)
       }
 //   userstruct, _, _, _,_ := uses.UserPage(eclient,session.Values["Username"].(string),session.Values["DocID"].(string))
 //   cs = ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)} 

	err := UserFollow(eclient,session.Values['Username'],r.URL.Path[9:])
	if (err){
		fmt.Println(err);
	}

}
*/


func CreateProject(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 == nil){
		http.Redirect(w, r, "/~", http.StatusFound)
	}
    userstruct, _, _,_ := uses.UserPage(eclient,session.Values["Username"].(string),session.Values["DocID"].(string))
    cs := ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)} 
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"createProject-Nil",cs)
}


func call(w http.ResponseWriter, r *http.Request){
	// If followingStatus = no 
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
    http.Redirect(w, r, "/~", http.StatusFound)
    }

	r.ParseForm()
	fmt.Println(r.Form)
	fname := r.FormValue("userID")
	fmt.Println(fname)
	following := r.FormValue("Following")
	fmt.Println(following)

	isFollowed, err4 := uses.IsFollowed(eclient, session.Values["DocID"].(string),fname)
	if (err4 != nil){
		fmt.Println(err4)
	}
	if (isFollowed == true){
	fmt.Println("called unfollow in ajax button")
	err := uses.UserUnfollow(eclient,session.Values["DocID"].(string),fname)
	if (err != nil){
		fmt.Println(err);
	}
	}else{
	fmt.Println("called follow in ajax button")
	err := uses.UserFollow(eclient,session.Values["DocID"].(string),fname)
	if (err != nil){
		fmt.Println(err);
	}	
	}
	//params := r.URL.Query()
	//params.Get('testing123')
	//	hello := "<div style='color:red;'>hello {{.UserInfo.FirstName}} do you understand the power of the http protocol</div>"
	//	fmt.Fprintln(w, hello) 
	// LINE 430 FEEELS SOOO GOOOOD !!!!!!!!!!!!!!!!!!!!!
}



func Like(w http.ResponseWriter, r *http.Request){
	// If followingStatus = no 
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
    http.Redirect(w, r, "/~", http.StatusFound)
    }

	r.ParseForm()
	fmt.Println(r.Form)
	fname := r.FormValue("userID")
	fmt.Println(fname)
	following := r.FormValue("Following")
	fmt.Println(following)

	isLiked, err4 := uses.IsFollowed(eclient, session.Values["DocID"].(string),fname)
	if (err4 != nil){
		fmt.Println(err4)
	}
	if (isLiked == true){
	fmt.Println("called unfollow in ajax button")
	err := uses.UserUnfollow(eclient,session.Values["DocID"].(string),fname)
	if (err != nil){
		fmt.Println(err);
	}
	}else{
	fmt.Println("called follow in ajax button")
	err := uses.UserFollow(eclient,session.Values["DocID"].(string),fname)
	if (err != nil){
		fmt.Println(err);
	}	
	}

}
func AddComment(w http.ResponseWriter, r *http.Request){
	// If followingStatus = no 
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
    if (test1 == nil){
     	fmt.Println(test1)
     	fmt.Println("^ is the username?")
    http.Redirect(w, r, "/~", http.StatusFound)
    }

    r.ParseForm()
	//postaid := r.FormValue("followstat")
	postid := r.FormValue("followstat")
	postactual := postid[1:]
	fmt.Println(postid+"is the postid"+postactual)
	commentz := r.FormValue("commentz")
	id := r.FormValue("id")
	fmt.Println(commentz+" is the input")
	fmt.Println("MADE IT HERE &&&&&&&&&&&&&&")
	contentarray := []rune(commentz)
	username := r.FormValue("username")
	fmt.Println("USERNAME IN ADD COMMENTS IS "+username)
	// journal entry, err 
	fmt.Println(contentarray)
	fmt.Println(id+" is doc id")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~``")
	err4 := uses.UserNewReplyEntry(eclient,id,contentarray,postactual)
		if (err4 != nil){
		fmt.Println(err4)
	}

   http.Redirect(w, r, "/profile/"+username, http.StatusFound)


}
func getComments(w http.ResponseWriter, r *http.Request){
	// If followingStatus = no 
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
    http.Redirect(w, r, "/~", http.StatusFound)
    }

	r.ParseForm()
	postid := r.FormValue("PostID")
	postaid := postid[9:]
	postactual := postid[10:]
	fmt.Println(postaid+" is the post id ")
	fmt.Println(postactual+" is the actual post id ")
	pika := r.FormValue("Pikachu")
	fmt.Println(pika+" is the pika value");
	// journal entry, err 
	parentPost, arrayofComments, err4 := uses.LoadComments(eclient, postactual, 0, -1)
	if (err4 != nil){
		fmt.Println(err4)
	}

	fmt.Println("hello get comments called")
	fmt.Println(parentPost.FirstName+" is parentpost first name")
	var sum int 
	var output string 
	for i := 0; i < len(arrayofComments); i++ {
		fmt.Println(arrayofComments[i].FirstName)
		sum += i
	}
	fmt.Println(sum)
//	id := session.Values["DOCID"].(string)
	username := session.Values["Username"].(string)
	fmt.Println("username is "+session.Values["Username"].(string))
	output += `
	 <div class="modal fade" id=main-moda`+postaid+` role="dialog">
                                <div class="modal-dialog">
                                    <!-- Modal content-->
                                    <div class="modal-content">
                                        <div class="modal-header">
                                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                                            <div class="media">
                                                <a class="pull-left" href="#">
                                                    <img class="media-object img-rounded" src="https://scontent-lga3-1.xx.fbcdn.net/v/t31.0-8/12514060_499384470233859_6798591731419500290_o.jpg?oh=329ea2ff03ab981dad7b19d9172152b7&oe=5A2D7F0D">
                                                </a>
                                                <div class="media-body">
                                                    <h6 class="pull-right text-muted time">3 hours ago</h6>
                                                    <h5 class="mt-0" style="color:cadetblue;">Ryan Rozbiani</h5>
                                                    <p>Hey guys! We're launching UStart! Watch out for us!</p>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="modal-body">
                                            <div class="input-group">
                                                <form id="commentform" method="POST" action="/AddComment">
                                                    <input name="commentz" placeholder="Add a comment" type="text">
                                                      <input type="hidden" name="followstat" value=`+postaid+`>
                                                      <input type="hidden" name = "id" value=`+pika+`>
                                                      <input type ="hidden" name="username" value=`+username+`>
                                                </form>
                                                <span class="input-group-addon">
                                                    <a onclick="document.getElementById('commentform').submit();">
                                                    <script>
                                                    console.log('inside the its not gonna work because it's just hml stuff so put inside script')
                                                    </script>
                                                        <i class="fa fa-edit"></i>
                                                    </a>
                                                </span>
                                            </div>
                                            <br>


	`

	//params := r.URL.Query()
	//params.Get('testing123')
	fmt.Fprintln(w, output) 

}


func call2(w http.ResponseWriter, r *http.Request){
	// If followingStatus = yes 
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
    http.Redirect(w, r, "/~", http.StatusFound)
    }
	fmt.Println("called unfollow")
	r.ParseForm()
	fmt.Println(r.Form)
	fname := r.FormValue("userID")
	fmt.Println(fname)

	err := uses.UserUnfollow(eclient,session.Values["DocID"].(string),fname)
	if (err != nil){
		fmt.Println(err);
	}
	//params := r.URL.Query()
	//params.Get('testing123')

}

func Inbox(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
     if (test1 == nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/~", http.StatusFound)
       }
    userstruct, _, _,_ := uses.UserPage(eclient,session.Values["Username"].(string),session.Values["DocID"].(string))
    cs := ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)} 
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"inbox-Nil",cs)
		


}

func Settings(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
     if (test1 == nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/~", http.StatusFound)
       }
    userstruct, _,_,_ := uses.UserPage(eclient,session.Values["Username"].(string),session.Values["DocID"].(string))
    cs := ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)}  
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"settings-Nil",cs)
		


}


func AJAX(w http.ResponseWriter, r *http.Request){
    cs := ClientSide{}  
	renderTemplate(w,"ajax-nil",cs)
		


}

func WallPostCreation(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    textb := r.FormValue("block")

    textb2 := []rune(textb)
    err := uses.UserNewTextEntry(eclient,session.Values["DocID"].(string),textb2)
    if (err != nil){
    	fmt.Println(err);
    }

    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}

func WallPostComment(w http.ResponseWriter, r *http.Request){

}
/*
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    textb := r.FormValue("block")

    textb2 := []rune(textb)
    err := uses.UserNewReplyEntry(eclient,session.Values["DocID"].(string),textb2,<postid>)
    if (err != nil){
    	fmt.Println(err);
    }

    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}
*/

func ImageUpload(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    blob := r.FormValue("image-data")
    blob = blob[1:len(blob)]
 //   fmt.Println(blob)

      //fmt.Println(reflect.TypeOf(blob))
    err := uses.ChangeAccountImagesAndStatus(eclient, session.Values["DocID"].(string),blob,true,"hello","Avatar");
    if (err != nil){
    	fmt.Println(err);
    }

  //  cs := ClientSide{ImageCode:blob} 


 //   post.HoldThis(eclient,blob)
    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}
/*
func ChangeDescription(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    phone := r.FormValue["pnumber"];

 //   fmt.Println(blob)

      //fmt.Println(reflect.TypeOf(blob))
 //   err := uses.ChangeContactAndDescription(eclient, session.Values["DocID"].(string),blob,true,"hello","Avatar");
    if (err != nil){
    	fmt.Println(err);
    }

  //  cs := ClientSide{ImageCode:blob} 


 //   post.HoldThis(eclient,blob)
    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}
*/



func SettingOptions1(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
     if (test1 == nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/~", http.StatusFound)
       }
    r.ParseForm();
    newDescription := r.FormValue("");   
    err := uses.ModifyDescription(eclient, session.Values["DocID"].(string), newDescription);
    if (err != nil){
    	fmt.Println(err);
    }   

    http.Redirect(w, r, "/Settings/", http.StatusFound)	



}

func ChangeContactAndDescription(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 == nil){
    	fmt.Println(test1)
    	http.Redirect(w, r, "/~", http.StatusFound) }
	r.ParseForm()
	var pVIS bool
	var gVIS bool 
	var eVIS bool 
	phonenumber := r.FormValue("pnumber")
	phonenumbervis := r.FormValue("pnumberVis")
	if (phonenumbervis == "True"){
		pVIS = true 
	}else{
		pVIS = false 
	}
	gender := r.FormValue("gender_select")
	gendervis := r.FormValue("gender_selectVis")
	if (gendervis == "True"){
		gVIS = true 
	}else{
		gVIS = false 
	}
	email := r.FormValue("inputEmail")
	emailvis := r.FormValue("inputEmailVis")
	if (emailvis == "True"){
		eVIS = true 
	}else{
		eVIS = false 
	}
	description := r.FormValue("inputDesc")
	descriptionrune := []rune(description)

	userID := session.Values["DocID"].(string)
    err2 := uses.ChangeContactAndDescription(eclient, userID, phonenumber, pVIS, gender, gVIS, email, eVIS, descriptionrune)   
	if (err2 != nil){
		fmt.Println(err2)
	}   
	/*
    userstruct, _, _,_,_ := uses.UserPage(eclient,session.Values["DocID"].(string),session.Values["DocID"].(string))
    cs := ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)}  
	renderTemplate(w,"template2-nil",cs)
	renderTemplate(w,"profile-nil",cs)
	*/


		if (err2 == nil){
		         http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	}
		


}

func changeName(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    first := r.FormValue("fname")
    last := r.FormValue("lname")
 //   fmt.Println(blob)

      //fmt.Println(reflect.TypeOf(blob))
    err := uses.ChangeFirstAndLastName(eclient, session.Values["DocID"].(string),first,last);
    if (err != nil){
    	fmt.Println(err);
    }
    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}

func changePassword(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    oldp := r.FormValue("oldpass")
    newp := r.FormValue("confirmpass")
    oldpb := []byte(oldp)
    newpb := []byte(newp)
 //   fmt.Println(blob)

      //fmt.Println(reflect.TypeOf(blob))
    err := uses.ChangePassword(eclient, session.Values["DocID"].(string),oldpb,newpb);
    if (err != nil){
    	fmt.Println(err);
    }
    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}

func changeLocation(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    countryP := r.FormValue("country")
    countryPV := r.FormValue("countryVis")
 //   fmt.Println(countryPV)
    stateP := r.FormValue("state")
    statePV := r.FormValue("stateVis")
    cityP := r.FormValue("city")
    cityPV := r.FormValue("cityVis")
    zipP := r.FormValue("zip")
    zipPV := r.FormValue("zipVis")
    conBool := false;
    if (countryPV == "on"){
    	conBool = true; 
    }
    sBool := false;
    if (statePV == "on"){
    	sBool = true; 
    }
    cBool := false;
    if (cityPV == "on"){
   	cBool = true; 
    }
    zBool := false;
    if (zipPV == "on"){
    	zBool = true; 
    }

    err := uses.ChangeLocation(eclient, session.Values["DocID"].(string),countryP,conBool,stateP,sBool,cityP,cBool,zipP,zBool);
   if (err != nil){
   		fmt.Println(err);
   }
   http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}

func changeEDU(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
	r.ParseForm()
	typeAcc := r.FormValue("type_select")
	i,err2 := strconv.Atoi(typeAcc)
	if (err2 != nil){
		fmt.Println(err2);
	}
	highschoolName := r.FormValue("schoolname")
	highschoolGrad := r.FormValue("highSchoolGradDate")
	uniName := r.FormValue("universityName")
	var major []string
	major = append(major,r.FormValue("majors"))
	//	Year := r.FormValue("year")
	gradDate := r.FormValue("uniGradDate")

	var minor []string

	err := uses.ChangeEducation(eclient, session.Values["DocID"].(string), i, highschoolName, highschoolGrad, uniName, gradDate, major, minor)
	if (err != nil){
		fmt.Println(err);
	}
}

func deleteAccount(w http.ResponseWriter, r *http.Request){
}

func LoginError(w http.ResponseWriter, r *http.Request){	
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
     if (test1 != nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/profile/"+session.Values["DocID"].(string), http.StatusFound)
       }
	//	u.FirstName = r.FormValue("firstName")
	email := r.FormValue("email")
	fmt.Println(email)
	//	var password []byte
	password := r.FormValue("password")
	fmt.Println(password)
	//	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	passwordb := []byte(password)
	successful,sessionInfo,err2 :=  uses.Login(eclient, email, passwordb)
 	if (err2 != nil){
		fmt.Println(err2)
	
	}

	if (successful == true){
		fmt.Println("login successful")
		session.Values["DocID"] = sessionInfo.DocID
		session.Values["FirstName"] = sessionInfo.FirstName
		session.Values["LastName"] = sessionInfo.LastName
		session.Values["Email"] = sessionInfo.Email
    	expiration := time.Now().Add((30) * time.Hour)
    	cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", Expires: expiration, Path:"/"}
    	http.SetCookie(w, &cookie)
		session.Save(r,w)
    	http.Redirect(w, r, "/profile/"+session.Values["DocID"].(string), http.StatusFound)	
	}

	if (successful == false){
		fmt.Println("did not login successful")
		var errorL bool
		errorL = true
	//	cs := ClientSide{ErrorLogin: errorL}
		fmt.Println("errorL is ")
		fmt.Print(errorL)
		http.Redirect(w, r, "/loginerror-nil/", http.StatusFound)	
		
		

		
	}

	}

func LogOut(w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "session_please")
			test1, _ := session.Values["DocID"]
     if (test1 != nil){
     	session.Options.MaxAge = -1
     	session.Save(r,w)
     	http.Redirect(w, r, "/~", http.StatusFound)

       }


}


//func renderTemplate(w http.ResponseWriter, tmpl string, u *types.User) {
//	t, err := template.ParseFiles(("../../../../www/ustart.tech/"+tmpl + ".html"))
//	fmt.Println("parsing 5")
//

//	t.Execute(w,u)

func renderTemplate(w http.ResponseWriter, tmpl string, cs ClientSide) {
	//  	fmt.Println("rT called")
  	err := templates.ExecuteTemplate(w, tmpl+".html", cs)
  	if err != nil {
  		http.Error(w, err.Error(), http.StatusInternalServerError)
  	}
  }

  func renderTemplate2(w http.ResponseWriter, tmpl string, name string) {
	//  	fmt.Println("rT called")
  	err := templates.ExecuteTemplate(w, tmpl+".html", name)
  	if err != nil {
  		http.Error(w, err.Error(), http.StatusInternalServerError)
  	}
  }

 func imageEx(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fileupload := r.FormValue("inputEmail")
	fmt.Println(fileupload)


} 



func main() {
	http.HandleFunc("/test/",Test )
	fs := http.FileServer(http.Dir("../../../../www/"))
//	r.Handle("/www/", http.StripPrefix("/www/", http.FileServer(http.Dir("../../../../www/"))))
	http.Handle("/www/", http.StripPrefix("/www/", fs))
	http.HandleFunc("/Signup/", Signup)
	http.HandleFunc("/registrationcomplete/", RegistrationComplete)
	http.HandleFunc("/welcome/", Registration)
	http.HandleFunc("/profilelogin/", Login)
	http.HandleFunc("/profile/", ViewProfile)
//	r.HandleFunc("/loggedin/", Login)
	http.HandleFunc("/imagetest/",imageEx)
	http.HandleFunc("/logout/",LogOut)
	http.HandleFunc("/Inbox/",Inbox)
	http.HandleFunc("/Projects/",ProjectsPage)
	http.HandleFunc("/MyProjects/",MyProjects)
	http.HandleFunc("/Settings/",Settings)
	http.HandleFunc("/ImageUpload/",ImageUpload)
	http.HandleFunc("/changeName/",changeName)
	http.HandleFunc("/changePass/",changePassword)
	http.HandleFunc("/changeLoc/",changeLocation)
	http.HandleFunc("/changeEDU/",changeEDU)
	http.HandleFunc("/deleteAccount/",deleteAccount)
	http.HandleFunc("/UpdateDescription/",ChangeContactAndDescription)
	http.HandleFunc("/CreateProject/",CreateProject)
	http.HandleFunc("/loginerror/",LoginError)
//	http.HandleFunc("/WallTest/",WallTest)
	http.HandleFunc("/~",Home)
	http.HandleFunc("/Registration/Type/",RegisterType)
	http.HandleFunc("/callme/",call)
//	http.HandleFunc("/hellomoto/",GetComments)
	http.HandleFunc("/callme2/",call2)
	http.HandleFunc("/follow/",Follow)
	http.HandleFunc("/unfollow/",Follow)
	http.HandleFunc("/New/Post/",WallPostCreation)
	http.HandleFunc("/New/Comment/",WallPostComment)
	http.HandleFunc("/getComments/",getComments)
	http.HandleFunc("/AddComment",AddComment)

	http.HandleFunc("/ajax/",AJAX)

//	fmt.Println("testing")
	http.ListenAndServe(":5000", nil)
}