package repositories

import (
	"github.com/sknv/pgup/orm/repository"

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
	params repository.PagingParams, query ...interface{},
) ([]*models.User, error) {
	var result []*models.User
	err := u.Base.FindPage(&result, params, query...)
	return result, err
}

func (u *User) FindOneById(id int64) (*models.User, error) {
	result := &models.User{}
	err := u.Base.FindOneById(result, id)
	return result, err
}
