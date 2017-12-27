package controllers

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/sknv/chip"
	"github.com/sknv/chip/render"
	"github.com/sknv/mng"
	mnware "github.com/sknv/mng/middleware"

	"github.com/sknv/chipapp/src/apps/api/forms"
	"github.com/sknv/chipapp/src/lib/models"
	"github.com/sknv/chipapp/src/lib/repositories"
	"github.com/sknv/chipapp/src/lib/utils"
)

type User struct {
	*Base `inject:""`

	UserRepo *repositories.User `inject:""`
}

func (c *User) Index(w http.ResponseWriter, r *http.Request) {
	// Parse request for fetching params.
	fetchPrm := utils.GetFetchingParams(w, r)

	// Fetch and paginate users.
	mgoSession := mnware.GetMgoSession(r)
	users, err := c.UserRepo.FindPage(mgoSession, fetchPrm.Query, fetchPrm.PagingParams)
	chip.PanicIfError(err)

	render.Json(w, http.StatusOK, users)
}

func (c *User) Show(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id.
	mgoSession := mnware.GetMgoSession(r)
	user := c.fetchUser(w, r, mgoSession)
	render.Json(w, http.StatusOK, user)
}

func (c *User) Create(w http.ResponseWriter, r *http.Request) {
	// Bind the request body to a user form and validate it.
	userForm := c.bindRequestToUserForm(w, r)

	// Fill the model.
	user := &models.User{}
	userForm.FillModel(w, user)

	// Insert model to the db.
	mgoSession := mnware.GetMgoSession(r)
	err := c.UserRepo.Insert(mgoSession, user)

	// Check the violation of unique indexes and panic in case of other error.
	if mgo.IsDup(err) {
		utils.RenderPlainAndAbort(w, http.StatusUnprocessableEntity, err.Error())
	}
	chip.PanicIfError(err)

	render.Json(w, http.StatusCreated, user)
}

func (c *User) Update(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id.
	mgoSession := mnware.GetMgoSession(r)
	user := c.fetchUser(w, r, mgoSession)

	// Bind the request body to a user form and validate it.
	userForm := c.bindRequestToUserForm(w, r)

	// Fill the model and update the db.
	userForm.FillModel(w, user)
	err := c.UserRepo.UpdateDoc(mgoSession, user)

	// Check the document existence, violation of unique indexes
	// and panic in case of other error.
	if mng.IsErrNotFound(err) {
		utils.RenderStatusAndAbort(w, http.StatusNotFound)
	} else if mgo.IsDup(err) {
		utils.RenderPlainAndAbort(w, http.StatusUnprocessableEntity, err.Error())
	}
	chip.PanicIfError(err)

	render.Json(w, http.StatusOK, user)
}

func (c *User) Destroy(w http.ResponseWriter, r *http.Request) {
	// Fetch the user by id and remote it.
	mgoSession := mnware.GetMgoSession(r)
	user := c.fetchUser(w, r, mgoSession)
	err := c.UserRepo.RemoveDoc(mgoSession, user)

	// Check the document existence and panic in case of other error.
	if mng.IsErrNotFound(err) {
		utils.RenderStatusAndAbort(w, http.StatusNotFound)
	}
	chip.PanicIfError(err)

	render.Status(w, http.StatusNoContent)
}

// fetchUser fetches the user by id, panics in case of error.
func (c *User) fetchUser(
	w http.ResponseWriter, r *http.Request, session *mgo.Session,
) *models.User {
	// Get url parameter 'id' from request.
	id := chi.URLParam(r, "id")

	// Check the document existence and panic in case of other error.
	user, err := c.UserRepo.FindOneById(session, id)
	if mng.IsErrNotFound(err) {
		utils.RenderStatusAndAbort(w, http.StatusNotFound)
	}
	chip.PanicIfError(err)

	return user
}

// bindRequestToUserForm binds the request body to a user form and validates it.
// Panics in case of error.
func (c *User) bindRequestToUserForm(w http.ResponseWriter, r *http.Request) *forms.User {
	userForm := &forms.User{}
	utils.BindRequestAndValidate(w, r, c.Validate, userForm)
	return userForm
}
