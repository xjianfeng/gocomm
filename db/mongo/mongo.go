package mongo

import (
	"gopkg.in/mgo.v2"
)

var MONGO_SESSION *mgo.Session

type MongoDb struct {
	session    *mgo.Session
	DataBase   *mgo.Database
	Collection *mgo.Collection
}

func GetMongo(db, collect string) *MongoDb {
	session := MONGO_SESSION.Clone()
	mondb := session.DB(db)
	collection := mondb.C(collect)

	mongo := MongoDb{
		session,
		mondb,
		collection,
	}
	return &mongo
}

func (m *MongoDb) UseDBCollect(dbname string, collect string) *mgo.Collection {
	m.DataBase = m.session.DB(dbname)
	m.Collection = m.DataBase.C(collect)
	return m.Collection
}

func (m *MongoDb) UseCollect(collect string) *mgo.Collection {
	m.Collection = m.DataBase.C(collect)
	return m.Collection
}

func (m *MongoDb) Close() {
	m.session.Close()
}

func SetUp(mongoUri string) {
	var err error
	MONGO_SESSION, err = mgo.Dial(mongoUri)
	MONGO_SESSION.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err.Error())
	}
}
