package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CategoryPortfolio represents an Portfolio in Category record.
type CategoryPortfolio struct {
	ID          bson.ObjectId `json:"id" db:"id" bson:"_id"`
	Title       string        `json:"title" db:"title" bson:"title"`
	ShortDesc   string        `json:"shortDesc" db:"shortDesc" bson:"shortDesc"`
	Description string        `json:"description" db:"description" bson:"description"`
	Logo        string        `json:"logo" db:"logo" bson:"logo"`
	Banner      string        `json:"banner" db:"banner" bson:"banner"`
}

// Category represents an Category record.
type Category struct {
	ID         bson.ObjectId       `json:"id" db:"id" bson:"_id"`
	Title      string              `json:"title" db:"name" bson:"title"`
	Image      string              `json:"image" db:"image" bson:"image"`
	Keywords   []string            `json:"keywords" db:"keywords" bson:"keywords"`
	Portfolios []CategoryPortfolio `json:"portfolios" db:"portfolios" bson:"portfolios"`
	AddedBy    string              `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt    time.Time           `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
