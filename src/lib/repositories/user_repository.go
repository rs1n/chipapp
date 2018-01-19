package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sknv/mng/odm/repository"

	"github.com/sknv/chipapp/src/lib/models"
)

const collectionUsers = "users"

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{NewBaseRepository(collectionUsers)}
}

func (r *UserRepository) FindAll(
	session *mgo.Session, query bson.M, sort []string,
) ([]*models.User, error) {
	var result []*models.User
	err := r.BaseRepository.FindAll(session, query, sort, &result)
	return result, err
}

func (r *UserRepository) FindOneById(session *mgo.Session, id string) (*models.User, error) {
	result := &models.User{}
	err := r.BaseRepository.FindOneById(session, id, result)
	return result, err
}

func (r *UserRepository) FindOneByLogin(session *mgo.Session, login string) (*models.User, error) {
	result := &models.User{}
	err := r.FindOne(session, bson.M{"login": login}, result)
	return result, err
}

func (u *UserRepository) FindPage(
	session *mgo.Session, query bson.M, params repository.PagingParams,
) ([]*models.User, error) {
	var result []*models.User
	err := u.BaseRepository.FindPage(session, query, params, &result)
	return result, err
}
