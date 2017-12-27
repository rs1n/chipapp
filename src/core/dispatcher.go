package core

import (
	"github.com/go-chi/chi"
	"github.com/sknv/chip"

	"github.com/sknv/chipapp/src/apps/api"
	"github.com/sknv/chipapp/src/apps/web"
)

type Dispatcher struct {
	ApiApp *api.Application `inject:""`
	WebApp *web.Application `inject:""`
}

func NewDispatcher(objectProvider *ObjectProvider) *Dispatcher {
	dispatcher := &Dispatcher{}
	err := Inject(dispatcher, objectProvider.Objects...)
	chip.PanicIfError(err)

	// TODO: init applications if needed.

	return dispatcher
}

// Dispatch dispatches incoming requests.
func (d *Dispatcher) Dispatch(r chi.Router) {
	d.ApiApp.Route(r)
	d.WebApp.Route(r)
}
