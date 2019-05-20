package model

import (
	"gopkg.in/mgo.v2/bson"
)

// User example
type User struct {
	ID    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"first_name" form:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" form:"last_name" bson:"last_name"`
	Email     string `json:"email" form:"email" bson:"email"`
	PassWord  string `json:"password" form:"password" bson:"password"`
	Token     string `json:"token,omitempty" bson:"-"`
}
//users example
type Users struct {
	Users []User `json:"users" bson:"users"`
}

//func NewUser(email, password string) *User {
//	p, _ := util.HashPassword(password)
//	return &User{
//		ID:       bson.NewObjectId(),
//		Email:    email,
//		PassWord: p,
//	}
//}
