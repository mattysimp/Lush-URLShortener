package handlers

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
)

// MongoStore holds the mongodb session
type MongoStore struct {
	session *mgo.Session
	col     mgo.Collection
}

func (ms *MongoStore) SetURL(url *URL) (err error) {
	err = ms.col.Insert(url)
	return err
}

func (ms *MongoStore) GetURL(code string) (url *URL, err error) {
	url = &URL{}
	err = ms.col.FindId(code).One(url)
	if err != nil {
		return nil, errors.New("url not found")
	}
	return url, nil
}

// NewDB creates MongoStore with a session created with MongoDB connection constants
func NewDB(config *Config) *MongoStore {
	info := &mgo.DialInfo{
		Addrs:    []string{config.Database.Host},
		Timeout:  10 * time.Second,
		Database: config.Database.DBName,
		Username: "",
		Password: "",
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return &MongoStore{session, *session.DB(config.Database.DBName).C(config.Database.Collection)}
}
