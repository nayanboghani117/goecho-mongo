package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// VideosVideos represents an Videos in Videos record.
type VideosVideos struct {
	Title       string `json:"title" db:"title" bson:"title"`
	Description string `json:"description" db:"description" bson:"description"`
	Link        string `json:"link" db:"link" bson:"link"`
}

// Videos represents an Videos record.
type Videos struct {
	ID      bson.ObjectId  `json:"id" db:"id" bson:"_id"`
	App     string         `json:"app" db:"name" bson:"app"`
	Videos  []VideosVideos `json:"videos" db:"image" bson:"videos"`
	AddedBy string         `json:"addedBy" db:"addedBy" bson:"addedBy"`
	AddedAt time.Time      `json:"addedAt" db:"addedAt" bson:"addedAt"`
}
