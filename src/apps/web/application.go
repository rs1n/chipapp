package web

import (
	"github.com/go-chi/chi"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"

	"github.com/sknv/chipapp/src/apps"
	"github.com/sknv/chipapp/src/apps/web/controllers"
)

type Application struct {
	apps.Application

	homeController *controllers.Home
}

func NewApplication(
	htmlRender *render.Html, validate *validate.Validate,
) *Application {
	return &Application{
		homeController: controllers.NewHome(htmlRender, validate),
	}
}

func (a *Application) Route(r chi.Router) {
	r.Get("/", a.homeController.Index)
	r.Get("/hello/{name}", a.homeController.Index)
}
