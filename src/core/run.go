package core

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/rs1n/chip"
	xhttp "github.com/rs1n/chip/x/net/http"

	"github.com/rs1n/chipapp/src/config"
	"github.com/rs1n/chipapp/src/core/global"
)

const (
	publicRoot      = "./public"
	shutdownTimeout = 10 * time.Second
	templateRoot    = "./templates"
	templateExt     = ".tpl"
)

func Run() {
	// Create an application's global context and schedule a cleaning.
	initGlobal()
	defer global.GetGlobal().CleanUp()

	// Create and bootstrap a router.
	router := chi.NewRouter()
	bootstrapRouter(router)

	// Dispatch requests and serve the router.
	NewDispatcher().Dispatch(router)
	serveRouter(router)
}

// initGlobal creates a new application's global context.
func initGlobal() {
	cfg := config.GetConfig()
	global.InitGlobalFor(
		global.HtmlRendererParams{
			IsDebug:      cfg.IsDebug,
			TemplateRoot: templateRoot,
			TemplateExt:  templateExt,
		},
	)
}

// bootstrapRouter plugs standard middleware and serves static files.
func bootstrapRouter(r chi.Router) {
	chip.BootstrapRouter(r)
	chip.ServeRoot(r, publicRoot)
}

// serveRouter starts the server on specified port.
func serveRouter(r chi.Router) {
	cfg := config.GetConfig()
	xhttp.Serve(r, cfg.Port, shutdownTimeout)
}
