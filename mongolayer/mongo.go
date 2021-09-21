package mongolayer

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"myevents/persistence"
)

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	mongoSession, err := mgo.Dial(connection)

	if err != nil {
		return nil, err
	}

	return &MongoDBLayer{
		session: mongoSession,
	}, err
}

func (mgoLayer *MongoDBLayer) Addevent(e persistence.Event) ([]byte, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error){
	s := mgoLayer.getFreshSession()
	defer s.Close()

	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayerr *MongoDBLayer) getFreshSession() *mgo.Session{
	return mgoLayerr.session.Copy()
}