package core

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/sknv/chip"
	xhttp "github.com/sknv/chip/x/net/http"
	"github.com/sknv/mng"

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
	objectProvider := initObjectProvider(cfg)
	defer objectProvider.CleanUp()

	// Create and bootstrap a router.
	router := chi.NewRouter()
	bootstrapRouter(router, objectProvider.MgoSession)

	// Dispatch requests and serve the router on specified port.
	NewDispatcher(objectProvider).Dispatch(router)
	xhttp.Serve(router, cfg.Port, shutdownTimeout)
}

// initObjectProvider creates a new application's global context.
func initObjectProvider(config *config.Config) *ObjectProvider {
	hrp := HtmlRenderParams{
		IsDebug:      config.IsDebug,
		TemplateRoot: templateRoot,
		TemplateExt:  templateExt,
	}
	return NewObjectProvider(hrp, config)
}

// bootstrapRouter plugs standard middleware, provides a Mongo session
// and serves static files.
func bootstrapRouter(r chi.Router, mgoSession *mgo.Session) {
	chip.BootstrapRouter(r)
	mng.BootstrapRouter(r, mgoSession)

	chip.ServeRoot(r, publicRoot)
}
