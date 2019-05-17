package blog

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists blog data in database
type DAO struct{}

var collection = "blogs"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the blog with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.Blog, error) {
	blogs := []models.Blog{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&blogs)
	return blogs, err
}

// GetOne reads the blog with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Blog, error) {
	var blog models.Blog
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&blog)
	return &blog, err
}

// Create saves a new blog record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *DAO) Create(rs app.RequestScope, blog *models.Blog) error {
	blog.ID = bson.NewObjectId()
	blog.AddedBy = rs.UserName()
	blog.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(blog)
}

// Update saves the changes to an blog in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, blog *models.Blog) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	blog.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)

	return c.Update(bson.M{"_id": id}, &blog)
}

// Delete deletes an blog with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// // Count returns the number of the blog records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the blog records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.Blog, error) {
// 	blogs := []models.Blog{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&blogs)
// 	return blogs, err
// }
