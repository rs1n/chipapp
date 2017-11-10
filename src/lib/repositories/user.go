package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/skkv/chipapp/src/lib/models"
)

const userCollectionName = "users"

type User struct {
	*Base
}

func NewUser() *User {
	return &User{
		Base: NewBase(userCollectionName),
	}
}

func (r *User) FindPage(
	session *mgo.Session, query bson.M, skip, limit int,
) ([]*models.User, error) {
	result := []*models.User{}
	err := r.Base.FindPage(session, query, skip, limit, &result)
	return result, err
}

func (r *User) FindOneByHexId(
	session *mgo.Session, id string,
) (*models.User, error) {
	result := &models.User{}
	err := r.Base.FindOneByHexId(session, id, result)
	return result, err
}
