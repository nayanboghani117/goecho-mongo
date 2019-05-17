package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ServiceService represents an Service in Service record.
type ServiceService struct {
	Title       string `json:"title" db:"title" bson:"title"`
	Description string `json:"description" db:"description" bson:"description"`
}

// ServiceProcess represents an Process in Service record.
type ServiceProcess struct {
	Title       string `json:"title" db:"title" bson:"title"`
	Description string `json:"description" db:"description" bson:"description"`
	Image       string `json:"image" db:"image" bson:"image"`
}

// ServiceVideo represents an Videos in Service record.
type ServiceVideo struct {
	Title string `json:"title" db:"title" bson:"title"`
	Link  string `json:"link" db:"link" bson:"link"`
}

// Service represents an Service record.
type Service struct {
	ID       bson.ObjectId    `json:"id" db:"id" bson:"_id"`
	Services []ServiceService `json:"services" db:"services" bson:"services"`
	Process  []ServiceProcess `json:"process" db:"process" bson:"process"`
	Video    ServiceVideo     `json:"video" db:"video" bson:"video"`
	AddedBy  string           `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt  time.Time        `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
