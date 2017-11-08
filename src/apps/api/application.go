package api

import (
	"github.com/go-chi/chi"

	"github.com/rs1n/chipapp/src/apps/api/controllers"
)

const scope = "/api"

type Application struct {
	pingController *controllers.Ping
	userController *controllers.User
}

func NewApplication() *Application {
	return &Application{
		pingController: &controllers.Ping{},
		userController: controllers.NewUser(),
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
