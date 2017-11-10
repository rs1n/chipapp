package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/skkv/chip"
	"github.com/skkv/chip/mng"
	"github.com/skkv/chip/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/skkv/chipapp/src/apps/api/forms"
	"github.com/skkv/chipapp/src/lib/models"
	"github.com/skkv/chipapp/src/lib/repositories"
)

type User struct {
	base

	userRepository *repositories.User
}

func NewUser() *User {
	return &User{
		userRepository: repositories.NewUser(),
	}
}

func (c *User) Index(w http.ResponseWriter, r *http.Request) {
	mgoSession := mng.GetMgoSessionForRequest(r)

	// Fetch and paginate users.
	users, err := c.userRepository.FindPage(mgoSession, bson.M{}, 0, 0)
	chip.PanicIfError(err)
	render.Json(w, http.StatusOK, users)
}

func (c *User) Show(w http.ResponseWriter, r *http.Request) {
	mgoSession := mng.GetMgoSessionForRequest(r)

	// Fetch the user by id.
	user := c.fetchUser(w, r, mgoSession)
	render.Json(w, http.StatusOK, user)
}

func (c *User) Create(w http.ResponseWriter, r *http.Request) {
	// Bind the request body to a user form and validate it.
	userForm := c.bindRequestToUserForm(w, r)

	// Fill the model.
	user := &models.User{}
	userForm.FillModel(user)

	// Insert model to the db.
	mgoSession := mng.GetMgoSessionForRequest(r)
	err := c.userRepository.Insert(mgoSession, user)

	// Check the violation of unique indexes and panic in case of other error.
	if mgo.IsDup(err) {
		render.Json(w, http.StatusUnprocessableEntity, err)
		chip.AbortHandler()
	}
	chip.PanicIfError(err)

	render.Json(w, http.StatusCreated, user)
}

func (c *User) Update(w http.ResponseWriter, r *http.Request) {
	mgoSession := mng.GetMgoSessionForRequest(r)

	// Fetch the user by id.
	user := c.fetchUser(w, r, mgoSession)

	// Bind the request body to a user form and validate it.
	userForm := c.bindRequestToUserForm(w, r)

	// Fill the model and update the db.
	userForm.FillModel(user)
	err := c.userRepository.Update(mgoSession, user)

	// Check the document existence, violation of unique indexes
	// and panic in case of other error.
	if mng.IsErrNotFound(err) {
		render.Status(w, http.StatusNotFound)
		chip.AbortHandler()
	} else if mgo.IsDup(err) {
		render.Json(w, http.StatusUnprocessableEntity, err)
		chip.AbortHandler()
	}
	chip.PanicIfError(err)

	render.Json(w, http.StatusOK, user)
}

func (c *User) Destroy(w http.ResponseWriter, r *http.Request) {
	mgoSession := mng.GetMgoSessionForRequest(r)

	// Fetch the user by id and remote it.
	user := c.fetchUser(w, r, mgoSession)
	err := c.userRepository.Remove(mgoSession, user)

	// Check the document existence and panic in case of other error.
	if mng.IsErrNotFound(err) {
		render.Status(w, http.StatusNotFound)
		chip.AbortHandler()
	}
	chip.PanicIfError(err)

	render.Status(w, http.StatusNoContent)
}

// fetchUser fetches the user by id, panics in case of error.
func (c *User) fetchUser(
	w http.ResponseWriter, r *http.Request, session *mgo.Session,
) *models.User {
	// Get url parameter from request.
	id := chi.URLParam(r, "id")

	// Check the document existence and panic in case of other error.
	user, err := c.userRepository.FindOneByHexId(session, id)
	if mng.IsErrNotFound(err) {
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
