package core

import (
	"github.com/go-chi/chi"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"

	"github.com/sknv/chipapp/src/apps/api"
	"github.com/sknv/chipapp/src/apps/web"
)

type Dispatcher struct {
	apiApplication *api.Application
	webApplication *web.Application
}

func NewDispatcher(
	htmlRender *render.Html, validate *validate.Validate,
) *Dispatcher {
	return &Dispatcher{
		apiApplication: api.NewApplication(validate),
		webApplication: web.NewApplication(htmlRender, validate),
	}
}

// Dispatch dispatches incoming requests.
func (d *Dispatcher) Dispatch(r chi.Router) {
	d.apiApplication.Route(r)
	d.webApplication.Route(r)
}
