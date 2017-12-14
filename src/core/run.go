package core

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/sknv/chip"
	xhttp "github.com/sknv/chip/x/net/http"

	"github.com/sknv/chipapp/src/config"
	"github.com/sknv/chipapp/src/core/global"
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

	// Dispatch requests and serve the router on specified port.
	NewDispatcher().Dispatch(router)
	xhttp.Serve(router, config.GetConfig().Port, shutdownTimeout)
}

// initGlobal creates a new application's global context.
func initGlobal() {
	cfg := config.GetConfig()
	hrp := global.HtmlRenderParams{
		IsDebug:      cfg.IsDebug,
		TemplateRoot: templateRoot,
		TemplateExt:  templateExt,
	}
	global.InitGlobalFor(hrp, cfg.Postgres)
}

// bootstrapRouter plugs standard middleware, provides a Mongo session
// and serves static files.
func bootstrapRouter(r chi.Router) {
	chip.BootstrapRouter(r)
	chip.ServeRoot(r, publicRoot)
}
