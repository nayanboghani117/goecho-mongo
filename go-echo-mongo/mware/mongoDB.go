package mware

import (
	"errors"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jpillora/backoff"
	"gopkg.in/mgo.v2"
)

// DB represents the structure of our database
type DB struct {
	Name    string
	Session *mgo.Session
	Uri     string
}

// NewDB creates a DB instance
func NewDB(uri string) DB {
	return DB{
		Name: "Demoapp",
		Uri:  uri,
	}
}

// Dial connects to the server
func (db *DB) Dial() (err error) {
	b := &backoff.Backoff{
		Jitter: true,
	}

	for {
		db.Session, err = mgo.Dial(db.Uri)

		if err != nil {
			d := b.Duration()
			logrus.Errorf("%s, reconnecting in %s", err, d)
			time.Sleep(d)
			continue
		}

		b.Reset()

		logrus.Info("Successfully connected to MongoDB")

		db.Session.SetSocketTimeout(time.Second * 3)
		db.Session.SetSyncTimeout(time.Second * 3)
		return nil
	}
}

// Init initializes the database and creates indexes
func (db *DB) Init() (err error) {
	db.Dial()
	session := db.Session.Copy()
	d := session.DB(db.Name)
	defer session.Close()

	c := d.C("user")
	if c == nil {
		e := fmt.Sprint("Error creating collection")
		logrus.Error(e)
		return errors.New(e)
	}

	usersIndex := mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true,
	}
	c.EnsureIndex(usersIndex)

	return nil
}
