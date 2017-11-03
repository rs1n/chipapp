package apps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// bootstrapRouter bootstraps chi.Router:
// use standard middleware.
func bootstrapRouter(r chi.Router) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}
