package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Blog represents an blog record.
type Blog struct {
	ID          bson.ObjectId `json:"id" db:"id" bson:"_id"`
	Title       string        `json:"title" db:"title" bson:"title"`
	Description string        `json:"description" db:"description" bson:"description"`
	Images      []string      `json:"images" db:"images" bson:"images"`
	Link        string        `json:"link" db:"link" bson:"link"`
	Likes       int           `json:"likes" db:"likes" bson:"likes"`
	Views       int           `json:"views" db:"views" bson:"views"`
	AddedBy     string        `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt     time.Time     `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
