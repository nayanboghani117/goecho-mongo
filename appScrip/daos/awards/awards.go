package awards

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists awards data in database
type DAO struct{}

var collection = "awards"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the awards with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.Awards, error) {
	awards := []models.Awards{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&awards)
	return awards, err
}

// GetOne reads the awards with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Awards, error) {
	var awards models.Awards
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&awards)
	return &awards, err
}

// Create saves a new awards record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *DAO) Create(rs app.RequestScope, awards *models.Awards) error {
	awards.ID = bson.NewObjectId()
	awards.AddedBy = rs.UserName()
	awards.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(awards)
}

// Update saves the changes to an awards in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, awards *models.Awards) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	awards.ID = id
	awards.AddedBy = rs.UserName()
	awards.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &awards)
}

// Delete deletes an awards with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// // Count returns the number of the awards records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the awards records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.Awards, error) {
// 	awardss := []models.Awards{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&awardss)
// 	return awardss, err
// }
