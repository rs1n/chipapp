package repositories

import (
	"github.com/skkv/chip/mng/odm/document"
	"github.com/skkv/chip/mng/odm/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/skkv/chipapp/src/config"
)

// Base application repository.
type Base struct {
	*repository.Base
}

func NewBase(collectionName string) *Base {
	return &Base{
		Base: &repository.Base{
			DbName:         config.GetConfig().Mongo.Database,
			CollectionName: collectionName,
		},
	}
}

func (r *Base) FindOneByHexId(
	session *mgo.Session, id string, result interface{},
) error {
	return r.FindOneById(session, bson.ObjectIdHex(id), result)
}

func (r *Base) Insert(session *mgo.Session, doc interface{}) error {
	return r.Base.Insert(session, doc)
}

func (r *Base) Update(session *mgo.Session, doc document.IIdentifier) error {
	return r.Base.UpdateDoc(session, doc)
}

func (r *Base) Remove(session *mgo.Session, doc document.IIdentifier) error {
	return r.Base.RemoveDoc(session, doc)
}
