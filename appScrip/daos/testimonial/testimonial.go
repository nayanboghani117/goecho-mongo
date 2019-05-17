package testimonial

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists testimonial data in database
type DAO struct{}

var collection = "testimonial"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the testimonial with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.Testimonial, error) {
	testimonial := []models.Testimonial{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&testimonial)
	return testimonial, err
}

// GetOne reads the testimonial with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Testimonial, error) {
	var testimonial models.Testimonial
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&testimonial)
	return &testimonial, err
}

// Create saves a new testimonial record in the database.
func (dao *DAO) Create(rs app.RequestScope, testimonial *models.Testimonial) error {
	testimonial.ID = bson.NewObjectId()
	testimonial.AddedBy = rs.UserName()
	testimonial.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(testimonial)
}

// Update saves the changes to an testimonial in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, testimonial *models.Testimonial) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	testimonial.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &testimonial)
}

// Delete deletes an testimonial with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// Authenticate Testimonial
func (dao *DAO) Authenticate(rs app.RequestScope, email string, pass string) (*models.Testimonial, error) {
	var testimonial models.Testimonial
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"email": email}).One(&testimonial)
	return &testimonial, err
}

// // Count returns the number of the testimonial records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the testimonial records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.Testimonial, error) {
// 	testimonials := []models.Testimonial{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&testimonials)
// 	return testimonials, err
// }
