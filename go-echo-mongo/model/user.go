package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"first_name" form:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" form:"last_name" bson:"last_name"`
	Email     string `json:"email" form:"email" bson:"email"`
	PassWord  string `json:"password" form:"password" bson:"password"`
	Token     string `json:"token,omitempty" bson:"-"`
}
type Users struct {
	Users []User `json:"users" bson:"users"`
}

