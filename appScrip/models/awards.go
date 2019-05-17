package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Awards represents an Awards record.
type Awards struct {
	ID      bson.ObjectId `json:"id" db:"id" bson:"_id"`
	Title   string        `json:"title" db:"name" bson:"title"`
	Image   string        `json:"image" db:"image" bson:"image"`
	AddedBy string        `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt time.Time     `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
