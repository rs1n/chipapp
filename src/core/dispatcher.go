package core

import (
	"github.com/go-chi/chi"

	"github.com/sknv/chipapp/src/apps/api"
	"github.com/sknv/chipapp/src/apps/web"
)

type Dispatcher struct {
	ApiApp *api.Application
	WebApp *web.Application
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		ApiApp: api.NewApplication(),
		WebApp: web.NewApplication(),
	}
}

// Dispatch dispatches incoming requests.
func (d *Dispatcher) Dispatch(router chi.Router) {
	d.ApiApp.Route(router)
	d.WebApp.Route(router)
}
