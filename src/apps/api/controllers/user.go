package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sknv/chip"
	"github.com/sknv/chip/render"
	"github.com/sknv/pgup"

	"github.com/sknv/chipapp/src/apps/api/forms"
	"github.com/sknv/chipapp/src/lib/models"
	"github.com/sknv/chipapp/src/lib/repositories"
)

type User struct {
	Base

	userRepository *repositories.User
}

func NewUser() *User {
	return &User{
		userRepository: repositories.NewUser(),
	}
}

func (c *User) Index(w http.ResponseWriter, r *http.Request) {
	// Parse request for fetching params.
	fp, err := c.GetFetchingParamsForRequest(r)
	if err != nil {
		render.Status(w, http.StatusBadRequest)
		chip.AbortHandler()
	}

	// Fetch and paginate users.
	users, err := c.userRepository.FindPage(fp.PagingParams, fp.Query)
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

	// Insert model to the db.
	id, err := c.userRepository.Insert(user)

	// Check the violation of unique indexes and panic in case of other error.
	if pgup.IsPgDup(err) {
		render.Plain(w, http.StatusUnprocessableEntity, err.Error())
		chip.AbortHandler()
	}
	chip.PanicIfError(err)

	user.Id = id
	render.Json(w, http.StatusCreated, user)
}

func (c *User) Update(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id.
	user := c.fetchUser(w, r)

	// Bind the request body to a user form and validate it.
	userForm := c.bindRequestToUserForm(w, r)

	// Fill the model and update the db.
	userForm.FillModel(user)
	err := c.userRepository.UpdateRecord(user)

	// Check the document existence, violation of unique indexes
	// and panic in case of other error.
	if pgup.IsErrNoMoreRows(err) {
		render.Status(w, http.StatusNotFound)
		chip.AbortHandler()
	} else if pgup.IsPgDup(err) {
		render.Plain(w, http.StatusUnprocessableEntity, err.Error())
		chip.AbortHandler()
	}
	chip.PanicIfError(err)

	render.Json(w, http.StatusOK, user)
}

func (c *User) Destroy(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id and remote it.
	user := c.fetchUser(w, r)
	err := c.userRepository.DeleteRecord(user)

	// Check the document existence and panic in case of other error.
	if pgup.IsErrNoMoreRows(err) {
		render.Status(w, http.StatusNotFound)
		chip.AbortHandler()
	}
	chip.PanicIfError(err)

	render.Status(w, http.StatusNoContent)
}

// fetchUser fetches the user by id, panics in case of error.
func (c *User) fetchUser(w http.ResponseWriter, r *http.Request) *models.User {
	// Get url parameter from request.
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Status(w, http.StatusBadRequest)
		chip.AbortHandler()
	}

	// Check the document existence and panic in case of other error.
	user, err := c.userRepository.FindOneById(id)
	if pgup.IsErrNoMoreRows(err) {
		render.Status(w, http.StatusNotFound)
		chip.AbortHandler()
	}
	chip.PanicIfError(err)

	return user
}

// bindRequestToUserForm binds the request body to a user form and validates it.
// Panics in case of error.
func (c *User) bindRequestToUserForm(
	w http.ResponseWriter, r *http.Request,
) *forms.User {
	userForm := forms.NewUser(r)
	if err := c.ValidateStruct(userForm); err != nil {
		render.Json(w, http.StatusUnprocessableEntity, err)
		chip.AbortHandler()
	}
	return userForm
}
