package core

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/sknv/chip"
	xhttp "github.com/sknv/chip/x/net/http"

	"github.com/sknv/chipapp/src/config"
)

const (
	publicRoot      = "./public"
	shutdownTimeout = 10 * time.Second
	templateRoot    = "./templates"
	templateExt     = ".tpl"
)

func Run() {
	// Load the environment configuration.
	cfg := config.NewConfig()

	// Create an application's global context and schedule a cleaning.
	serviceProvider := initServiceProvider(cfg)
	defer serviceProvider.CleanUp()

	// Create and bootstrap a router.
	router := chi.NewRouter()
	bootstrapRouter(router)

	// Dispatch requests and serve the router on specified port.
	NewDispatcher(
		serviceProvider.HtmlRender,
		serviceProvider.PgSession,
		serviceProvider.Validate,
	).Dispatch(router)
	xhttp.Serve(router, cfg.Port, shutdownTimeout)
}

// initServiceProvider creates a new application's global context.
func initServiceProvider(cfg *config.Config) *ServiceProvider {
	hrp := HtmlRenderParams{
		IsDebug:      cfg.IsDebug,
		TemplateRoot: templateRoot,
		TemplateExt:  templateExt,
	}
	return NewServiceProvider(hrp, cfg.Postgres)
}

// bootstrapRouter plugs standard middleware, provides a Mongo session
// and serves static files.
func bootstrapRouter(r chi.Router) {
	chip.BootstrapRouter(r)
	chip.ServeRoot(r, publicRoot)
}
