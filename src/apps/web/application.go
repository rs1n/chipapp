package web

import (
	"github.com/go-chi/chi"

	"github.com/skkv/chipapp/src/apps"
	"github.com/skkv/chipapp/src/apps/web/controllers"
)

type Application struct {
	apps.Application

	homeController *controllers.Home
}

func NewApplication() *Application {
	return &Application{
		homeController: &controllers.Home{},
	}
}

func (a *Application) Route(r chi.Router) {
	r.Get("/", a.homeController.Index)
	r.Get("/hello/{name}", a.homeController.Index)
}
