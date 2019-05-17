package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Identity interface
type Identity interface {
	GetID() string
	GetName() string
}

// User represents an user record.
type User struct {
	ID       	bson.ObjectId `json:"id" db:"id" bson:"_id"`
	FirstName   string        `json:"firstname" db:"firstname" bson:"firstname"`
	LastName  	string 		  `json:"lastname" db:"lastname" bson:"lastname"`
	Email    	string        `json:"email" db:"email" bson:"email"`
	PassWord 	string        `json:"password" db:"password" bson:"password"`
	AddedBy  	string        `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt  	time.Time     `json:"addedAt" db:"addedAt" bson:"addedAt"`
}

//GetID for getting id
func (u User) GetID() string {
	return string(u.ID)
}

//GetName for getting name
func (u User) GetName() string {
	return u.FirstName
}

type Users struct {
	Users []User `json:"users" bson:"users"`
}