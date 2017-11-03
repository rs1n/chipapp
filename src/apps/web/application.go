package web

import (
	"github.com/go-chi/chi"

	"github.com/rs1n/chipapp/src/apps/web/controllers"
)

type Application struct {
	homeController *controllers.Home
}

func NewApplication() *Application {
	return &Application{
		homeController: controllers.NewHome(),
	}
}

func (a *Application) Route(r chi.Router) {
	r.Get("/", a.homeController.Index)
	r.Get("/hello/{name}", a.homeController.Index)
}
