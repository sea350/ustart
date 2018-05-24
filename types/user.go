package types

import (
	"time"
)

//Link ... Link related data
type Link struct {
	URL  string `json:"URL"`
	Name string `json:"Name"`
}

//ProjectInfo ... project info, duh
type ProjectInfo struct {
	ProjectID string `json:"ProjectID"`
	Visible   bool   `json:"Visible"`
}

//Request ... uhm
type Request struct {
	SenderID  string    `json:"UserID"`
	Timestamp time.Time `json:"Timestamp"`
}

//LoginWarning ... Security countermeasure for checking amount of login attempts and locking out IP address for repeated failures
type LoginWarning struct {
	LastAttempt     time.Time `json:"LastAttempt"`      //Time since the Last Failed Login Attempt
	NumberAttempts  int       `json:"NumberOfAttempts"` //Number of Failed Login Attempts
	LockoutUntil    time.Time `json:"LockoutUntil"`     //Lockout Until User can attempt again
	IPAddress       string    `json:"IPAddress"`        //IP address of Failed Login Attempt Offender
	LockoutCounter  int64     `json:"LockoutCounter"`   //Amount of Lockouts the IP address
	LoginWarningsIP []string  `json:"LoginWarnings"`    //Records the IP addresses attempting to login to specific account
}

//User ... all user related data
type User struct {
	Password        []byte        `json:"Password"` // Maybe we shouldn't keep it in plain text later?
	FirstName       string        `json:"FirstName"`
	LastName        string        `json:"LastName"`
	Email           string        `json:"Email"`
	EmailVis        bool          `json:"EmailVis"`
	Gender          string        `json:"Gender"`
	GenderVis       bool          `json:"GenderVis"`
	Phone           string        `json:"Phone"`
	PhoneVis        bool          `json:"PhoneVis"`
	Description     []rune        `json:"Description"`
	Blurb           []rune        `json:"Blurb"`
	Username        string        `json:"Username"`
	Location        LocStruct     `json:"Location"`
	HighSchool      string        `json:"HighSchool"`
	HSGradDate      string        `json:"HSGradDate"`
	CollegeGradDate string        `json:"CollegeGradDate"`
	University      string        `json:"UndergradSchool"`
	Majors          []string      `json:"Majors"`
	Minors          []string      `json:"Minors"`
	Class           int           `json:"Class"` //freshman:0,sophomore:1...
	Dob             time.Time     `json:"Dob"`
	AccCreation     time.Time     `json:"AccCreation"`
	Visible         bool          `json:"Visible"`
	Status          bool          `json:"Status"`
	ExpirationDate  time.Time     `json:"ExpirationDate"`
	Avatar          string        `json:"Avatar"`
	CroppedAvatar   string        `json:"CropAvatar"`
	Banner          string        `json:"Banner"`
	Organization    string        `json:"Organization"`
	Category        string        `json:"Category"`
	Paid            bool          `json:"Paid"`
	AccType         int           `json:"AccType"` //highschool, college etc
	QuickLinks      []Link        `json:"QuickLinks"`
	Tags            []string      `json:"Tags"`
	SearchHist      []string      `json:"SearchHist"`
	LikedEntryIDs   []string      `json:"LikedEntryIds"`
	Projects        []ProjectInfo `json:"Projects"`
	BlockedUsers    []string      `json:"BlockedUsers"`
	BlockedBy       []string      `json:"BlockedBy"`
	ConversationIDs []string      `json:"ConversationIDs"`
	EntryIDs        []string      `json:"EntryIDs"`
	Following       []string      `json:"Following"`
	Followers       []string      `json:"Followers"`
	Colleagues      []string      `json:"Colleagues"`
	SentCollReq     []string      `json:"SentCollReq"`
	ReceivedCollReq []string      `json:"ReceivedCollReq"`
	SentProjReq     []string      `json:"SentProjReq"`
	ReceivedProjReq []string      `json:"ReceivedProjReq"`
	FirstLogin      bool          `json:"FirstLogin"`
	UserWidgets     []string      `json:"UserWidgets"`
}
