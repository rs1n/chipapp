package controllers

import (
	"net/http"

	"github.com/rs1n/chip/render"

	"github.com/rs1n/chipapp/src/lib/repositories"
)

type User struct {
	*api

	userRepository *repositories.User
}

func NewUser() *User {
	return &User{
		api:            &api{},
		userRepository: repositories.NewUser(),
	}
}

func (c *User) Index(w http.ResponseWriter, r *http.Request) {
	users, err := c.userRepository.FindPage()
	if err != nil {
		panic(err)
	}
	render.Json(w, http.StatusOK, users)
}

func (c *User) Show(w http.ResponseWriter, r *http.Request) {
	user, err := c.userRepository.FindOneByHexId("")
	if err != nil {
		panic(err)
	}
	render.Json(w, http.StatusOK, user)
}

func (c *User) Create(w http.ResponseWriter, r *http.Request) {
	if err := c.userRepository.Insert(nil); err != nil {
		panic(err)
	}
	render.Json(w, http.StatusCreated, render.M{})
}

func (c *User) Update(w http.ResponseWriter, r *http.Request) {
	if err := c.userRepository.Update(nil); err != nil {
		panic(err)
	}
	render.Json(w, http.StatusOK, render.M{})
}

func (c *User) Destroy(w http.ResponseWriter, r *http.Request) {
	if err := c.userRepository.Remove(nil); err != nil {
		panic(err)
	}
	render.Json(w, http.StatusOK, render.M{})
}
