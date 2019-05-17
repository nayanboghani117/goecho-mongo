package user

import (
	"appScrip/app"
	"appScrip/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DAO persists blog data in database
type DAO struct{}

var collection = "users"

// NewDAO creates a new DAO
func NewDAO() *DAO {
	return &DAO{}
}

// Get reads the user with the specified ID from the database.
func (dao *DAO) Get(rs app.RequestScope) ([]models.User, error) {
	users := []models.User{}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{}).All(&users)
	return users, err
}

// GetOne reads the user with the specified ID from the database.
func (dao *DAO) GetOne(rs app.RequestScope, id bson.ObjectId) (*models.User, error) {
	var user models.User
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"_id": id}).One(&user)
	return &user, err
}

// Create saves a new user record in the database.
func (dao *DAO) Create(rs app.RequestScope, user *models.User) error {
	user.ID = bson.NewObjectId()
	user.AddedBy = rs.UserName()
	user.AddedAt = time.Now()
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Insert(user)
}

// Update saves the changes to an blog in the database.
func (dao *DAO) Update(rs app.RequestScope, id bson.ObjectId, user *models.User) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	user.ID = id
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Update(bson.M{"_id": id}, &user)
}

// Delete deletes an user with the specified ID from the database.
func (dao *DAO) Delete(rs app.RequestScope, id bson.ObjectId) error {
	if _, err := dao.GetOne(rs, id); err != nil {
		return err
	}
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	return c.Remove(bson.M{"_id": id})
}

// Authenticate User
func (dao *DAO) Authenticate(rs app.RequestScope, email string, pass string) (*models.User, error) {
	var user models.User
	c := rs.Tx().DB(app.Config.DBNAME).C(collection)
	err := c.Find(bson.M{"email": email}).One(&user)
	return &user, err
}

// // Count returns the number of the blog records in the database.
// func (dao *DAO) Count(rs app.RequestScope) (int, error) {
// 	var count int
// 	err := rs.Tx().Select("COUNT(*)").From("artist").Row(&count)
// 	return count, err
// }

// // Query retrieves the blog records with the specified offset and limit from the database.
// func (dao *DAO) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
// 	blogs := []models.User{}
// 	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&blogs)
// 	return blogs, err
// }
