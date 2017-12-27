package api

import (
	"github.com/go-chi/chi"

	"github.com/sknv/chipapp/src/apps/api/controllers"
)

const scope = "/api"

type Application struct {
	PingCtrl    *controllers.Ping    `inject:""`
	SessionCtrl *controllers.Session `inject:""`
	UserCtrl    *controllers.User    `inject:""`
}

func (a *Application) Route(r chi.Router) {
	r.Route(scope, func(r chi.Router) {
		a.routePing(r)
		a.routeSession(r)
		a.routeUser(r)
	})
}

func (a *Application) routePing(r chi.Router) {
	r.Get("/ping", a.PingCtrl.Index)
}

func (a *Application) routeSession(r chi.Router) {
	r.Route("/login", func(router chi.Router) {
		router.Post("/", a.SessionCtrl.Login)
	})
}

func (a *Application) routeUser(r chi.Router) {
	r.Route("/users", func(router chi.Router) {
		router.Get("/", a.UserCtrl.Index)
		router.Get("/{id}", a.UserCtrl.Show)
		router.Post("/", a.UserCtrl.Create)
		router.Put("/{id}", a.UserCtrl.Update)
		router.Delete("/{id}", a.UserCtrl.Destroy)
	})
}
