package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sknv/mng/odm/repository"

	"github.com/sknv/chipapp/src/lib/models"
)

const collectionUsers = "users"

type User struct {
	*Base
}

func NewUser() *User {
	return &User{
		Base: NewBase(collectionUsers),
	}
}

func (u *User) FindOneById(session *mgo.Session, id string) (*models.User, error) {
	result := &models.User{}
	err := u.Base.FindOneById(session, id, result)
	return result, err
}

func (r *User) FindOneByLogin(session *mgo.Session, login string) (*models.User, error) {
	result := &models.User{}
	err := r.FindOne(session, bson.M{"login": login}, result)
	return result, err
}

func (u *User) FindPage(
	session *mgo.Session, query bson.M, params repository.PagingParams,
) ([]*models.User, error) {
	var result []*models.User
	err := u.Base.FindPage(session, query, params, &result)
	return result, err
}
