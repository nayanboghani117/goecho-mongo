package category

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists blog data in database
type DAO struct{}

var collection = "category"
var collectionPortfolio = "portfolio"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the category with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.Category, error) {
	categories := []models.Category{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&categories)

	con := rs.Tx().DB(app.Config.DBNAME).C(collectionPortfolio)
	for i, val := range categories {
		categoryPortfolio := []models.CategoryPortfolio{}
		strCategoryID := bson.ObjectId(val.ID).Hex()
		con.Find(bson.M{"categoryId": strCategoryID}).All(&categoryPortfolio)
		categories[i].Portfolios = categoryPortfolio
	}

	return categories, err
}

// GetOne reads the category with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Category, error) {
	var category models.Category
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&category)
	return &category, err
}

// Create saves a new category record in the database.
func (dao *DAO) Create(rs app.RequestScope, category *models.Category) error {
	category.ID = bson.NewObjectId()
	category.AddedBy = rs.UserName()
	category.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(category)
}

// Update saves the changes to an blog in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, category *models.Category) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	category.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &category)
}

// Delete deletes an category with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// Authenticate Category
func (dao *DAO) Authenticate(rs app.RequestScope, email string, pass string) (*models.Category, error) {
	var category models.Category
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"email": email}).One(&category)
	return &category, err
}

// // Count returns the number of the blog records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the blog records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.Category, error) {
// 	blogs := []models.Category{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&blogs)
// 	return blogs, err
// }
