package core

import (
	"github.com/go-chi/chi"

	"github.com/skkv/chipapp/src/apps/api"
	"github.com/skkv/chipapp/src/apps/web"
)

type Dispatcher struct {
	apiApplication *api.Application
	webApplication *web.Application
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		apiApplication: api.NewApplication(),
		webApplication: web.NewApplication(),
	}
}

// Dispatch dispatches incoming requests.
func (d *Dispatcher) Dispatch(r chi.Router) {
	d.apiApplication.Route(r)
	d.webApplication.Route(r)
}
