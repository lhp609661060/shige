package models

import (
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

type MgoInterface interface {
	Create()
	BulkCreate()
	Update()
	Query()
	Find()
	Delete()
}

var mgoSession *mgo.Session
var mgoErr error

func MgoInit() *mgo.Session {
	mgoSession, mgoErr = mgo.Dial("localhost")
	if mgoErr != nil {
		panic(mgoErr)
	}


	mgoSession.SetMode(mgo.Monotonic, true)

	return mgoSession
}

func MgoDb(dbName string) *mgo.Database {
	return mgoSession.DB(dbName)
}

