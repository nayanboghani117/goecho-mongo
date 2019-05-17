package portfolio

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists portfolio data in database
type DAO struct{}

var collection = "portfolio"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the portfolio with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.Portfolio, error) {
	portfolios := []models.Portfolio{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&portfolios)
	return portfolios, err
}

// GetByCategory reads the portfolio with the specified ID from the database.
func (dao *DAO) GetByCategory(rs app.RequestScope, id string) ([]models.Portfolio, error) {
	portfolios := []models.Portfolio{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"categoryId": id}).All(&portfolios)
	return portfolios, err
}

// GetOne reads the portfolio with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.Portfolio, error) {
	var portfolio models.Portfolio
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&portfolio)
	return &portfolio, err
}

// Create saves a new portfolio record in the database.
// The Artist.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *DAO) Create(rs app.RequestScope, portfolio *models.Portfolio) error {
	portfolio.ID = bson.NewObjectId()
	portfolio.AddedBy = rs.UserName()
	portfolio.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(portfolio)
}

// Update saves the changes to an portfolio in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, portfolio *models.Portfolio) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	portfolio.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &portfolio)
}

// Delete deletes an portfolio with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// // Count returns the number of the portfolio records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the portfolio records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.Portfolio, error) {
// 	portfolios := []models.Portfolio{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&portfolios)
// 	return portfolios, err
// }
