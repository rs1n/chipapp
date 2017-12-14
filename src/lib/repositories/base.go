package repositories

import (
	"github.com/sknv/pgup/orm/repository"
	"upper.io/db.v3"
)

// Base application repository.
type Base struct {
	*repository.Base
}

func NewBase(session db.Database, collectionName string) *Base {
	return &Base{
		Base: &repository.Base{
			Session:        session,
			CollectionName: collectionName,
		},
	}
}

func (r *Base) FindPage(
	dest interface{}, params repository.PagingParams, query ...interface{},
) error {
	res := r.Base.FindPage(params, query...)
	err := res.All(dest)
	return err
}

func (r *Base) FindOneById(dest interface{}, id int64) error {
	res := r.Base.Find("id", id)
	err := res.One(dest)
	return err
}

func (r *Base) Insert(record interface{}) (int64, error) {
	id, err := r.Base.Insert(record)
	if err != nil {
		return 0, err
	}
	return id.(int64), nil
}
