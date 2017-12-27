package web

import (
	"github.com/go-chi/chi"

	"github.com/sknv/chipapp/src/apps/web/controllers"
)

type Application struct {
	HomeCtrl *controllers.Home `inject:""`
}

func (a *Application) Route(r chi.Router) {
	r.Get("/", a.HomeCtrl.Index)
	r.Get("/hello/{name}", a.HomeCtrl.Index)
}
