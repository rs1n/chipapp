package api

import (
	"github.com/go-chi/chi"

	"github.com/sknv/chipapp/src/apps/api/controllers"
)

const scope = "/api"

type Application struct {
	PingCtrl    *controllers.PingController
	SessionCtrl *controllers.SessionController
	UserCtrl    *controllers.UserController
}

func NewApplication() *Application {
	return &Application{
		PingCtrl:    controllers.NewPingController(),
		SessionCtrl: controllers.NewSessionController(),
		UserCtrl:    controllers.NewUserController(),
	}
}

func (a *Application) Route(router chi.Router) {
	router.Route(scope, func(r chi.Router) {
		a.routePing(r)
		a.routeSession(r)
		a.routeUser(r)
	})
}

func (a *Application) routePing(router chi.Router) {
	router.Get("/ping", a.PingCtrl.Index)
}

func (a *Application) routeSession(router chi.Router) {
	router.Route("/login", func(r chi.Router) {
		r.Post("/", a.SessionCtrl.Login)
	})
}

func (a *Application) routeUser(router chi.Router) {
	router.Route("/users", func(r chi.Router) {
		r.Get("/", a.UserCtrl.Index)
		r.Get("/all", a.UserCtrl.All)
		r.Get("/{id}", a.UserCtrl.Show)
		r.Post("/", a.UserCtrl.Create)
		r.Put("/{id}", a.UserCtrl.Update)
		r.Delete("/{id}", a.UserCtrl.Destroy)
	})
}
