package app

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

// RequestScope contains the application-specific information that are carried around in a request.
type RequestScope interface {
	Logger
	// UserID returns the ID of the user for the current request
	UserID() string
	// SetUserID sets the ID of the currently authenticated user
	SetUserID(id string)
	// UserName returns the ID of the user for the current request
	UserName() string
	// SetUserName sets the ID of the currently authenticated user
	SetUserName(name string)
	// RequestID returns the ID of the current request
	// Tx returns the currently active database transaction that can be used for DB query purpose
	Tx() *mgo.Session
	// SetTx sets the database transaction
	SetTx(tx *mgo.Session)
	// Now returns the timestamp representing the time when the request is being processed
	Now() time.Time
	// RequestID returns the ID of the current request
	RequestID() string
}

type requestScope struct {
	Logger                 // the logger tagged with the current request information
	userID    string       // an ID identifying the current user
	userName  string       // an name identifying the current user
	now       time.Time    // the time when the request is being processed
	tx        *mgo.Session // the currently active transaction
	requestID string       // an ID identifying one or multiple correlated HTTP requests
}

func (rs *requestScope) UserID() string {
	return rs.userID
}

func (rs *requestScope) SetUserID(id string) {
	rs.Logger.SetField("UserID", id)
	rs.userID = id
}

func (rs *requestScope) UserName() string {
	return rs.userName
}

func (rs *requestScope) SetUserName(name string) {
	rs.Logger.SetField("UserName", name)
	rs.userName = name
}

func (rs *requestScope) Tx() *mgo.Session {
	return rs.tx
}

func (rs *requestScope) SetTx(tx *mgo.Session) {
	rs.tx = tx
}

func (rs *requestScope) Now() time.Time {
	return rs.now
}

func (rs *requestScope) RequestID() string {
	return rs.requestID
}

// newRequestScope creates a new RequestScope with the current request information.
func newRequestScope(now time.Time, logger *logrus.Logger, request *http.Request) RequestScope {
	l := NewLogger(logger, logrus.Fields{})
	requestID := request.Header.Get("X-Request-Id")
	if requestID != "" {
		l.SetField("RequestID", requestID)
	}
	return &requestScope{
		Logger:    l,
		now:       now,
		requestID: requestID,
	}
}
