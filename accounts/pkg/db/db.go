package db

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session
var mongo_conn_str = "mongodb://localhost:27017"

// Creates a new session if mgoSession is nil i.e there is no active mongo session.
//If there is an active mongo session it will return a Clone
func GetMongoSession() (*mgo.Session, error) {
	var conn_str string = os.Getenv("MONGO_CONN_STR")
	if len(conn_str) <= 0 {
		conn_str = mongo_conn_str
	}
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(conn_str)
		if err != nil {
			return nil, err
		}
	}
	return mgoSession.Clone(), nil
}
