package apps

import (
	"github.com/go-chi/chi"
	"github.com/rs1n/chip"
)

const publicRoot = "./public"

// bootstrapRouter plugs standard middleware and serves static files.
func bootstrapRouter(r chi.Router) {
	chip.BootstrapRouter(r)
	chip.ServeRoot(r, publicRoot)
}
