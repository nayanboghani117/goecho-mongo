package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ContactUs represents an ContactUs record.
type ContactUs struct {
	ID      bson.ObjectId `json:"id" db:"id" bson:"_id"`
	Name    string        `json:"name" db:"name" bson:"name"`
	Email   string        `json:"email" db:"email" bson:"email"`
	Phone   string        `json:"phone" db:"phone" bson:"phone"`
	Subject string        `json:"subject" db:"subject" bson:"subject"`
	Body    string        `json:"body" db:"body" bson:"body"`
	// Read    string        `json:"read" db:"read" bson:"read"`
	AddedAt time.Time `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
