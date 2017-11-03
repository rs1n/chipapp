package apps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs1n/chip"
)

const publicRoot = "./public"

// bootstrapRouter bootstraps chi.Router:
// use standard middleware and serve files.
func bootstrapRouter(r chi.Router) {
	useStandardMiddleware(r)
	serveRoot(r)
}

func useStandardMiddleware(r chi.Router) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func serveRoot(r chi.Router) {
	chip.ServeRoot(r, publicRoot)
}
