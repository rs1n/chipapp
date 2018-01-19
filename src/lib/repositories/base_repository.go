package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sknv/mng/odm/repository"
)

const maxLimit = 25

// Base repository.
type BaseRepository struct {
	*repository.BaseRepository
}

func NewBaseRepository(collectionName string) *BaseRepository {
	return &BaseRepository{
		&repository.BaseRepository{
			CollectionName: collectionName,
			MaxLimit:       maxLimit,
		},
	}
}

func (r *BaseRepository) FindAll(
	session *mgo.Session, query bson.M, sort []string, result interface{},
) error {
	qry := r.Find(session, query)
	if len(sort) > 0 {
		qry = qry.Sort(sort...)
	}
	err := qry.All(result)
	return err
}

func (r *BaseRepository) FindOne(session *mgo.Session, query bson.M, result interface{}) error {
	qry := r.Find(session, query)
	err := qry.One(result)
	return err
}

func (r *BaseRepository) FindOneById(session *mgo.Session, id string, result interface{}) error {
	qry := r.BaseRepository.Find(session, bson.M{"_id": bson.ObjectIdHex(id)})
	err := qry.One(result)
	return err
}

func (r *BaseRepository) FindPage(
	session *mgo.Session, query bson.M, params repository.PagingParams,
	result interface{},
) error {
	qry := r.BaseRepository.FindPage(session, query, params)
	err := qry.All(result)
	return err
}
