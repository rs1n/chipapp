package api

import (
	"github.com/go-chi/chi"

	"github.com/rs1n/chipapp/src/apps/api/controllers"
)

const scope = "/api"

type Application struct {
	pingController *controllers.Ping
}

func NewApplication() *Application {
	return &Application{
		pingController: controllers.NewPing(),
	}
}

func (a *Application) Route(r chi.Router) {
	r.Route(scope, func(r chi.Router) {
		a.routePing(r)
	})
}

func (a *Application) routePing(r chi.Router) {
	r.Get("/ping", a.pingController.Index)
}
