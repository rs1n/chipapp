package controllers

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/sknv/chip"
	"github.com/sknv/chip/render"
	"github.com/sknv/mng"
	mware "github.com/sknv/mng/middleware"

	"github.com/sknv/chipapp/src/apps/api/forms"
	"github.com/sknv/chipapp/src/lib/models"
	"github.com/sknv/chipapp/src/lib/repositories"
	"github.com/sknv/chipapp/src/lib/utils"
)

type UserController struct {
	*BaseController

	UserRepo *repositories.UserRepository
}

func NewUserController() *UserController {
	return &UserController{
		BaseController: NewBaseController(),
		UserRepo:       repositories.NewUserRepository(),
	}
}

func (c *UserController) All(w http.ResponseWriter, r *http.Request) {
	// Parse the request for fetching params.
	fetchParams := utils.GetFetchingParams(w, r)

	// Fetch and sort documents.
	mgoSession := mware.GetMgoSession(r)
	docs, err := c.UserRepo.FindAll(mgoSession, fetchParams.Query, fetchParams.PagingParams.Sort)
	chip.PanicIfError(err)

	render.Json(w, http.StatusOK, docs)
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	// Parse the request for fetching params.
	fetchParams := utils.GetFetchingParams(w, r)

	// Fetch, sort and paginate documents.
	mgoSession := mware.GetMgoSession(r)
	docs, err := c.UserRepo.FindPage(mgoSession, fetchParams.Query, fetchParams.PagingParams)
	chip.PanicIfError(err)

	render.Json(w, http.StatusOK, docs)
}

func (c *UserController) Show(w http.ResponseWriter, r *http.Request) {
	// Fetch the document by id.
	mgoSession := mware.GetMgoSession(r)
	doc := c.fetchDocument(w, r, mgoSession)
	render.Json(w, http.StatusOK, doc)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	// Bind the request body to a form and validate it.
	form := c.bindRequestToForm(w, r)

	// Fill the model.
	doc := &models.User{}
	form.FillModel(w, doc)

	// Insert a model to the db.
	mgoSession := mware.GetMgoSession(r)
	err := c.UserRepo.Insert(mgoSession, doc)

	// Check the violation of unique indexes and panic in case of other error.
	if mgo.IsDup(err) {
		log.Print("error: ", err)
		utils.RenderJsonAndAbort(w, http.StatusUnprocessableEntity, forms.GetUserDupErrorMessage())
	}
	chip.PanicIfError(err)

	render.Json(w, http.StatusCreated, doc)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	// Fetch the document by id.
	mgoSession := mware.GetMgoSession(r)
	doc := c.fetchDocument(w, r, mgoSession)

	// Bind the request body to a form and validate it.
	form := c.bindRequestToForm(w, r)

	// Fill the model and update the db.
	form.FillModel(w, doc)
	err := c.UserRepo.UpdateDoc(mgoSession, doc)

	// Check the document existence, violation of unique indexes
	// and panic in case of other error.
	if mng.IsErrNotFound(err) {
		log.Print("error: ", err)
		utils.RenderStatusAndAbort(w, http.StatusNotFound)
	} else if mgo.IsDup(err) {
		log.Print("error: ", err)
		utils.RenderJsonAndAbort(w, http.StatusUnprocessableEntity, forms.GetUserDupErrorMessage())
	}
	chip.PanicIfError(err)

	render.Json(w, http.StatusOK, doc)
}

func (c *UserController) Destroy(w http.ResponseWriter, r *http.Request) {
	// Fetch the document by id and remote it.
	mgoSession := mware.GetMgoSession(r)
	doc := c.fetchDocument(w, r, mgoSession)
	err := c.UserRepo.RemoveDoc(mgoSession, doc)

	// Check the document existence and panic in case of other error.
	if mng.IsErrNotFound(err) {
		log.Print("error: ", err)
		utils.RenderStatusAndAbort(w, http.StatusNotFound)
	}
	chip.PanicIfError(err)

	render.Status(w, http.StatusNoContent)
}

// fetchDocument fetches the document by id, panics in case of error.
func (c *UserController) fetchDocument(
	w http.ResponseWriter, r *http.Request, session *mgo.Session,
) *models.User {
	// Get url parameter 'id' from request.
	id := chi.URLParam(r, "id")

	doc, err := c.UserRepo.FindOneById(session, id)
	// Check the document existence and panic in case of other error.
	if mng.IsErrNotFound(err) {
		log.Print("error: ", err)
		utils.RenderStatusAndAbort(w, http.StatusNotFound)
	}
	chip.PanicIfError(err)

	return doc
}

// bindRequestToForm binds the request body to a form and validates it.
// Panics in case of error.
func (c *UserController) bindRequestToForm(
	w http.ResponseWriter, r *http.Request,
) *forms.UserForm {
	form := &forms.UserForm{}
	utils.BindRequestAndValidate(w, r, c.Validate, form)
	return form
}
