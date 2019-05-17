package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// AboutPeople represents an Peoples in About record.
type AboutPeople struct {
	Name        string `json:"name" db:"name" bson:"name"`
	Designation string `json:"designation" db:"designation" bson:"designation"`
	ProfilePic  string `json:"profilePic" db:"profilePic" bson:"profilePic"`
}

// AboutTeam represents an Teams in About record.
type AboutTeam struct {
	Team    string `json:"team" db:"team" bson:"team"`
	Members int    `json:"members" db:"members" bson:"members"`
}

// AboutVideo represents an Videos in About record.
type AboutVideo struct {
	Title       string `json:"title" db:"title" bson:"title"`
	Description string `json:"description" db:"description" bson:"description"`
	Link        string `json:"link" db:"link" bson:"link"`
}

// AboutFaq represents an Faqs in About record.
type AboutFaq struct {
	Title       string `json:"title" db:"title" bson:"title"`
	Description string `json:"description" db:"description" bson:"description"`
}

// About represents an About record.
type About struct {
	ID          bson.ObjectId `json:"id" db:"id" bson:"_id"`
	Peoples     []AboutPeople `json:"peoples" db:"peoples" bson:"peoples"`
	Description string        `json:"description" db:"description" bson:"description"`
	Teams       []AboutTeam   `json:"teams" db:"teams" bson:"teams"`
	Video       AboutVideo    `json:"video" db:"video" bson:"video"`
	Faqs        []AboutFaq    `json:"faqs" db:"faqs" bson:"faqs"`
	AddedBy     string        `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt     time.Time     `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
