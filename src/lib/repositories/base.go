package repositories

import (
	"github.com/sknv/chip/mng/odm/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/sknv/chipapp/src/config"
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
