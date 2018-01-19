package web

import (
	"github.com/go-chi/chi"

	"github.com/sknv/chipapp/src/apps/web/controllers"
)

type Application struct {
	HomeCtrl *controllers.HomeController
}

func NewApplication() *Application {
	return &Application{
		HomeCtrl: controllers.NewHomeController(),
	}
}

func (a *Application) Route(router chi.Router) {
	router.Get("/", a.HomeCtrl.Index)
	router.Get("/hello/{name}", a.HomeCtrl.Index)
}
