package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// PortfolioAppLinks represents an AppLinks in Portfolio record.
type PortfolioAppLinks struct {
	Android string `json:"android" db:"android" bson:"android"`
	Ios     string `json:"ios" db:"ios" bson:"ios"`
}

// PortfolioTechnologies represents an Technologies in Portfolio record.
type PortfolioTechnologies struct {
	Title string `json:"title" db:"title" bson:"title"`
	Image string `json:"image" db:"image" bson:"image"`
}

// PortfolioVideos represents an Videos in Portfolio record.
type PortfolioVideos struct {
	Title       string   `json:"title" db:"title" bson:"title"`
	Description string   `json:"description" db:"description" bson:"description"`
	Link        []string `json:"link" db:"link" bson:"link"`
}

// PortfolioApps represents an Apps in Portfolio record.
type PortfolioApps struct {
	Title        string                  `json:"title" db:"title" bson:"title"`
	Description  string                  `json:"description" db:"description" bson:"description"`
	Images       []string                `json:"images" db:"images" bson:"images"`
	AppLinks     PortfolioAppLinks       `json:"appLinks" db:"appLinks" bson:"appLinks"`
	Technologies []PortfolioTechnologies `json:"technologies" db:"technologies" bson:"technologies"`
	VideoTitle   string                  `json:"videoTitle" db:"videoTitle" bson:"videoTitle"`
	Videos       []PortfolioVideos       `json:"videos" db:"videos" bson:"videos"`
}

// Portfolio represents an Portfolio record.
type Portfolio struct {
	ID           bson.ObjectId   `json:"id" db:"id" bson:"_id"`
	CategoryID   string          `json:"categoryId" db:"categoryId" bson:"categoryId"`
	CategoryName string          `json:"categoryName" db:"categoryName" bson:"categoryName"`
	Logo         string          `json:"logo" db:"logo" bson:"logo"`
	Banner       string          `json:"banner" db:"banner" bson:"banner"`
	Title        string          `json:"title" db:"title" bson:"title"`
	ShortDesc    string          `json:"shortDesc" db:"shortDesc" bson:"shortDesc"`
	Description  string          `json:"description" db:"description" bson:"description"`
	Apps         []PortfolioApps `json:"apps" db:"apps" bson:"apps"`
	InspiredBy   []string        `json:"inspiredBy" db:"inspiredBy" bson:"inspiredBy"`
	Clients      []string        `json:"clients" db:"clients" bson:"clients"`
	AddedBy      string          `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt      time.Time       `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
