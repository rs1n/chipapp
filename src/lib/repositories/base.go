package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sknv/mng/odm/repository"
)

// Base application repository.
type Base struct {
	*repository.Base
}

func NewBase(collectionName string) *Base {
	return &Base{
		Base: &repository.Base{
			CollectionName: collectionName,
		},
	}
}

func (r *Base) FindPage(
	session *mgo.Session, query bson.M, params repository.PagingParams,
	result interface{},
) error {
	qry := r.Base.FindPage(session, query, params)
	err := qry.All(result)
	return err
}

func (r *Base) FindOneById(
	session *mgo.Session, id string, result interface{},
) error {
	qry := r.Base.Find(session, bson.M{"_id": bson.ObjectIdHex(id)})
	err := qry.One(result)
	return err
}
