package handlers

import (
	"time"

	"gopkg.in/mgo.v2"
)

// Constants used in MongoDB
const (
	hosts           = "localhost:27020"
	database        = "urlshortener"
	username        = ""
	password        = ""
	collection      = "urls"
	pageSizeDefault = 5
)

// MongoStore holds the mongodb session
type MongoStore struct {
	session *mgo.Session
}

func (ms *MongoStore) SetURL(url *URL) (err error) {
	col := ms.session.DB(database).C(collection)
	err = col.Insert(url)
	return err
}

// NewDB creates MongoStore with a session created with MongoDB connection constants
func NewDB() *MongoStore {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  10 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return &MongoStore{session}
}
