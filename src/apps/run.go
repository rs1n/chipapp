package apps

import (
	"time"

	"github.com/go-chi/chi"
	xhttp "github.com/rs1n/chip/x/net/http"
)

const shutdownTimeout = 10 * time.Second

func Run() {
	// Create and bootstrap a router.
	router := chi.NewRouter()
	bootstrapRouter(router)

	// Dispatch requests and serve the router.
	NewDispatcher().Dispatch(router)
	serveRouter(router)
}

// serveRouter starts the server on the specified port.
func serveRouter(r chi.Router) {
	xhttp.Serve(r, 3000, shutdownTimeout)
}
