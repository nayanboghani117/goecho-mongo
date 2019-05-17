package service

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists service data in database
type DAO struct{}

var collection = "service"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the service with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.Service, error) {
	service := []models.Service{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&service)
	return service, err
}

// GetOne reads the service with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope) (*models.Service, error) {
	var service models.Service
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).One(&service)
	return &service, err
}

// Create saves a new service record in the database.
func (dao *DAO) Create(rs app.RequestScope, service *models.Service) error {
	service.ID = bson.NewObjectId()
	service.AddedBy = rs.UserName()
	service.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(service)
}

// Update saves the changes to an service in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, service *models.Service) error {
	if _, err := dao.GetOne(rs); err != nil {
		return err
	}
	service.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &service)
}

// Delete deletes an service with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// Authenticate Service
func (dao *DAO) Authenticate(rs app.RequestScope, email string, pass string) (*models.Service, error) {
	var service models.Service
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"email": email}).One(&service)
	return &service, err
}

// // Count returns the number of the service records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the service records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.Service, error) {
// 	services := []models.Service{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&services)
// 	return services, err
// }
