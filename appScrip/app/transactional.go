package app

import (
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/fault"
	"gopkg.in/mgo.v2"
)

// Transactional returns a handler that encloses the nested handlers with a DB transaction.
// If a nested handler returns an error or a panic happens, it will rollback the transaction.
// Otherwise it will commit the transaction after the nested handlers finish execution.
// By calling app.Context.SetRollback(true), you may also explicitly request to rollback the transaction.
func Transactional(db *mgo.Session) routing.Handler {
	return func(c *routing.Context) error {

		session := db.Copy()
		col := session.DB("baarkooDB").C("books")

		index := mgo.Index{
			Key:        []string{"isbn"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err := col.EnsureIndex(index)

		rs := GetRequestScope(c)
		rs.SetTx(db)

		err = fault.PanicHandler(rs.Errorf)(c)

		var e error
		if e != nil {
			if err == nil {
				// the error will be logged by an error handler
				return e
			}
			// log the tx error only
			rs.Error(e)
		}

		return err
	}
}
