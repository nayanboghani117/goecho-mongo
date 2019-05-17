package contactus

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists contactus data in database
type DAO struct{}

var collection = "contactus"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the contactus with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.ContactUs, error) {
	contactus := []models.ContactUs{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&contactus)
	return contactus, err
}

// GetOne reads the contactus with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.ContactUs, error) {
	var contactus models.ContactUs
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&contactus)
	return &contactus, err
}

// Create saves a new contactus record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *DAO) Create(rs app.RequestScope, contactus *models.ContactUs) error {
	contactus.ID = bson.NewObjectId()
	contactus.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(contactus)
}

// Update saves the changes to an contactus in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, contactus *models.ContactUs) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	contactus.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)

	return c.Update(bson.M{"_id": id}, &contactus)
}

// Delete deletes an contactus with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// // Count returns the number of the contactus records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the contactus records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.ContactUs, error) {
// 	contactus := []models.ContactUs{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&contactus)
// 	return contactus, err
// }
