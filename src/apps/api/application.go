package api

import (
	"github.com/go-chi/chi"
	"github.com/sknv/chip/validate"
	"upper.io/db.v3"

	"github.com/sknv/chipapp/src/apps"
	"github.com/sknv/chipapp/src/apps/api/controllers"
)

const scope = "/api"

type Application struct {
	apps.Application

	pingController *controllers.Ping
	userController *controllers.User
}

func NewApplication(
	session db.Database, validate *validate.Validate,
) *Application {
	return &Application{
		pingController: controllers.NewPing(validate),
		userController: controllers.NewUser(session, validate),
	}
}

func (a *Application) Route(r chi.Router) {
	r.Route(scope, func(r chi.Router) {
		a.routePing(r)
		a.routeUser(r)
	})
}

func (a *Application) routePing(r chi.Router) {
	r.Get("/ping", a.pingController.Index)
}

func (a *Application) routeUser(r chi.Router) {
	r.Route("/users", func(router chi.Router) {
		router.Get("/", a.userController.Index)
		router.Get("/{id}", a.userController.Show)
		router.Post("/", a.userController.Create)
		router.Put("/{id}", a.userController.Update)
		router.Delete("/{id}", a.userController.Destroy)
	})
}
