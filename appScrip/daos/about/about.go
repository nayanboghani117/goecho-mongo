package about

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists about data in database
type DAO struct{}

var collection = "about"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the about with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.About, error) {
	about := []models.About{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&about)
	return about, err
}

// GetOne reads the about with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope) (*models.About, error) {
	var about models.About
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).One(&about)
	return &about, err
}

// Create saves a new about record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *DAO) Create(rs app.RequestScope, about *models.About) error {
	about.ID = bson.NewObjectId()
	about.AddedBy = rs.UserName()
	about.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(about)
}

// Update saves the changes to an about in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, about *models.About) error {
	if _, err := dao.GetOne(rs); err != nil {
		return err
	}
	about.ID = id
	about.AddedBy = rs.UserName()
	about.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &about)
}

// Delete deletes an about with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// // Count returns the number of the about records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the about records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.About, error) {
// 	abouts := []models.About{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&abouts)
// 	return abouts, err
// }
