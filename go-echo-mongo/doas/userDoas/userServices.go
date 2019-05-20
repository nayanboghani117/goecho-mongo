package userDoas

import (
	//"fmt"
	"go-echo-mongo/go-echo-mongo/apis/handler"
	model "go-echo-mongo/go-echo-mongo/models"

	//"github.com/Sirupsen/logrus"
	//"gopkg.in/mgo.v2"

)


// CreateUser creates a new user
func CreateUser(db *handler.Handler,data  *model.User) error {
	session := db.DB.Session.Copy()
	d := session.DB(db.DB.Name)
	defer session.Close()

	response := d.C("users").Insert(data)
	//
	//if err != nil {
	//	if !mgo.IsDup(err) {
	//		e := fmt.Sprintf("Error creating user: %v", err)
	//		logrus.Panic(e)
	//		panic(e)
	//	}
	//	return err
	//}
	return response
}

//// DeleteUser deletes a user from an id
//func (db *DB) Delete(id string) error {
//	session := db.Session.Copy()
//	d := session.DB(db.Name)
//	defer session.Close()
//
//	err := d.C("users").RemoveId(bson.ObjectIdHex(id))
//	if err != nil && err != mgo.ErrNotFound {
//		e := fmt.Sprintf("Error deleting user: %s", err)
//		logrus.Panic(e)
//		panic(e)
//	} else if err == mgo.ErrNotFound {
//		return err
//	}
//
//	return nil
//}
//
////func (db *DB) Get(m bson.M, data interface{}) error {
////	session := db.Session.Copy()
////	d := session.DB(db.Name)
////	defer session.Close()
////
////	err := d.C("users").Find(m).One(data)
////
////	if err != nil && err != mgo.ErrNotFound {
////		e := fmt.Sprintf("Error finding user: %s", err)
////		logrus.Panic(e)
////		panic(e)
////	} else if err == mgo.ErrNotFound {
////		return err
////	}
////
////	return nil
////}
////
////// FindUserByEmail finds a user based on an email address
////func (db *DB) FindUserByEmail(email string, data interface{}) error {
////	m := bson.M{"email": email}
////	return db.findUser(m, data)
////}
////
////// FindUserById finds a user from an id
////func (db *DB) GetOne(id string, data interface{}) error {
////	m := bson.M{"_id": bson.ObjectIdHex(id)}
////	return db.findUser(m, data)
////}
////
////// GetAllUsers returns all the users
////func (db *DB) GetAll(data interface{}) error {
////	session := db.Session.Copy()
////	d := session.DB(db.Name)
////	defer session.Close()
////
////	err := d.C("users").Find(nil).All(data)
////	if err != nil && err != mgo.ErrNotFound {
////		e := fmt.Sprintf("Error finding users: %s", err)
////		logrus.Panic(e)
////		panic(e)
////	} else if err == mgo.ErrNotFound {
////		return err
////	}
////
////	return nil
////}
