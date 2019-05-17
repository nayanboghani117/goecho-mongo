package videos

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists video data in database
type DAO struct{}

var collection = "videos"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the video with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.Videos, error) {
	videos := []models.Videos{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&videos)
	return videos, err
}

// GetOne reads the video with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Videos, error) {
	var video models.Videos
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&video)
	return &video, err
}

// Create saves a new video record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *DAO) Create(rs app.RequestScope, video *models.Videos) error {
	video.ID = bson.NewObjectId()
	video.AddedBy = rs.UserName()
	video.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(video)
}

// Update saves the changes to an video in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, video *models.Videos) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	video.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &video)
}

// Delete deletes an video with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// // Count returns the number of the video records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the video records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.Videos, error) {
// 	videos := []models.Videos{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&videos)
// 	return videos, err
// }
