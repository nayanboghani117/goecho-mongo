package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Testimonial represents an Testimonial record.
type Testimonial struct {
	ID          bson.ObjectId `json:"id" db:"id" bson:"_id"`
	Title       string        `json:"title" db:"name" bson:"title"`
	Link        string        `json:"link" db:"link" bson:"link"`
	Description string        `json:"description" db:"description" bson:"description"`
	AddedBy     string        `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt     time.Time     `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
