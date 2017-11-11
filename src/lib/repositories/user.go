package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/sknv/chipapp/src/lib/models"
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

func (u *User) FindPage(
	session *mgo.Session, query bson.M, limit, skip int,
) ([]*models.User, error) {
	result := []*models.User{}
	err := u.Base.FindPage(session, query, limit, skip, &result)
	return result, err
}

func (u *User) FindOneByHexId(
	session *mgo.Session, id string,
) (*models.User, error) {
	result := &models.User{}
	err := u.Base.FindOneByHexId(session, id, result)
	return result, err
}
