package types

import (
	"time"
)

//invite only, publicly viewable, whitelist, blacklist, settings, members, privileges, title, description, creation, event date, option for
//creator to cancel, visibility boolean, widgets, get/post, attending/invited/not going, members list and guest list

type EventPrivileges struct {
	RoleName     string `json:"RoleName"`
	RoleID       int    `json:"RoleID"`
	MemberManage bool   `json:"MemberEdit"`
	WidgetManage bool   `json:"WidgetManage"`
	PostManage   bool   `json:"PostManage"`
	Icon         bool   `json:"Icon"`
	Banner       bool   `json:"Banner"`
	Links        bool   `json:"Links"`
	Tags         bool   `json:"Tags"`
}

type EventGuests struct {
	GuestID string `json:"GuestID"`
	Status  int    `json:"Status"` //Marks whether they are invited/going/not going, 0 for invited, 1 for going, 2 for not
	Visible bool   `json:"Visible'`
}

type EventMembers struct {
	MemberID string    `json:"MemberID"`
	Role     int       `json:"Role"`
	JoinDate time.Time `json:"JoinDate"`
	Title    string    `json:"Title"`
	Visible  bool      `json:"Visible"`
}
type Events struct {
	Name          string         `json:"Name"`
	URLName       string         `json:"URLName"`
	Description   []rune         `json:"Description"`
	Members       []EventMembers `json:"Members"`
	Guests        []EventGuests  `json:"Guests"`
	Location      LocStruct      `json:"Location"`
	EventDate     time.Time      `json:"EventDate"`
	CreationDate  time.Time      `json:"CreationDate"`
	Widgets       []Widget       `json:"Widgets"`
	Whitelist     string         `json:"Whitelist"`
	Blacklist     string         `json:"Blacklist"`
	Status        bool           `json:"Status"` //Whether this event is still ongoing or cancelled
	Public        bool           `json:"Public"` //Whether this event is publicly viewable or invite-only viewable
	Avatar        string         `json:"Avatar"`
	CroppedAvatar string         `json:"CropAvatar"`
	Banner        string         `json:"Banner"`
	MemberReqSent []string       `json:"MemberReqSent"`
	//MemberReqReceived []string       `json:"MemberReqReceived"` We probably don't want this
}
