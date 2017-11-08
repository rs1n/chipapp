package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs1n/chip"
	"github.com/rs1n/chip/render"

	"github.com/rs1n/chipapp/src/apps/api/forms"
	"github.com/rs1n/chipapp/src/core/validate"
	"github.com/rs1n/chipapp/src/lib/models"
	"github.com/rs1n/chipapp/src/lib/repositories"
)

type User struct {
	base

	userRepository *repositories.User
}

func NewUser() *User {
	return &User{
		userRepository: &repositories.User{},
	}
}

func (c *User) Index(w http.ResponseWriter, r *http.Request) {
	users, err := c.userRepository.FindPage()
	chip.PanicIfError(err)
	render.Json(w, http.StatusOK, users)
}

func (c *User) Show(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id.
	user := c.fetchUser(w, r)
	render.Json(w, http.StatusOK, user)
}

func (c *User) Create(w http.ResponseWriter, r *http.Request) {
	// Bind the request body to a user form and validate it.
	userForm := c.bindRequestToUserForm(w, r)

	// Fill the model.
	user := &models.User{}
	userForm.FillModel(user)

	if err := c.userRepository.Insert(user); err != nil {
		panic(err)
	}
	render.Json(w, http.StatusCreated, user)
}

func (c *User) Update(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id.
	user := c.fetchUser(w, r)

	// Bind the request body to a user form and validate it.
	userForm := c.bindRequestToUserForm(w, r)

	// Fill the model.
	userForm.FillModel(user)

	if err := c.userRepository.Update(user); err != nil {
		panic(err)
	}
	render.Json(w, http.StatusOK, user)
}

func (c *User) Destroy(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id.
	user := c.fetchUser(w, r)

	if err := c.userRepository.Remove(user); err != nil {
		panic(err)
	}
	render.Status(w, http.StatusNoContent)
}

// fetchUser fetches the user by id, panics in case of error.
func (c *User) fetchUser(w http.ResponseWriter, r *http.Request) *models.User {
	// Get url parameter.
	id := chi.URLParam(r, "id")

	user, err := c.userRepository.FindOneByHexId(id)
	chip.PanicIfError(err)
	return user
}

// bindRequestToUserForm binds the request body to a user form and validates it.
// Panics in case of error.
func (c *User) bindRequestToUserForm(
	w http.ResponseWriter, r *http.Request,
) *forms.User {
	userForm := forms.NewUser(r)

	if err := validate.Struct(userForm); err != nil {
		render.Json(w, http.StatusUnprocessableEntity, err)
		chip.AbortHandler()
	}
	return userForm
}
