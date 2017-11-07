package repositories

import (
	"github.com/rs1n/chipapp/src/lib/models"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (r *User) FindPage() ([]*models.User, error) {
	result := []*models.User{
		{
			Id:   "abc123",
			Name: "name",
			Profile: models.Profile{
				Email: "email@example.com",
			},
			Images: []*models.Image{
				{
					Src:   "/home/demodev/foo-thumb.png",
					Style: "thumb",
				},
			},
		},
	}

	return result, nil
}

func (r *User) FindOneByHexId(id string) (*models.User, error) {
	result := &models.User{
		Id:   "abc123",
		Name: "name",
		Profile: models.Profile{
			Email: "email@example.com",
		},
		Images: []*models.Image{
			{
				Src:   "/home/demodev/foo-thumb.png",
				Style: "thumb",
			},
		},
	}

	return result, nil
}

func (r *User) Insert(user *models.User) error {
	return nil
}

func (r *User) Update(user *models.User) error {
	return nil
}

func (r *User) Remove(user *models.User) error {
	return nil
}
