package core

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"github.com/sknv/chip"
	xhttp "github.com/sknv/chip/x/net/http"
	"github.com/sknv/mng"

	"github.com/sknv/chipapp/src/config"
	"github.com/sknv/chipapp/src/core/provider"
)

const (
	publicRoot      = "./public"
	shutdownTimeout = 10 * time.Second
	templateRoot    = "./templates"
	templateExt     = ".tpl"
)

func init() {
	// Include line numbers to a log.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Run() {
	// Load the environment configuration.
	cfg := config.NewConfig()

	// Create an application's global context and schedule a cleaning.
	objectProvider := newObjectProvider(cfg)
	defer objectProvider.CleanUp()

	// Create and bootstrap a router.
	router := chi.NewRouter()
	bootstrapRouter(router, objectProvider.MgoSession)

	// Dispatch requests and serve the router on specified port.
	NewDispatcher().Dispatch(router)
	xhttp.Serve(router, cfg.Port, shutdownTimeout)
}

// newObjectProvider creates a new application's global context.
func newObjectProvider(config *config.Config) *provider.ObjectProvider {
	hrp := provider.HtmlRenderParams{
		IsDebug:      config.IsDebug,
		TemplateRoot: templateRoot,
		TemplateExt:  templateExt,
	}
	return provider.NewObjectProvider(hrp, config)
}

// bootstrapRouter plugs standard middleware, provides a Mongo session
// and serves static files.
func bootstrapRouter(router chi.Router, mgoSession *mgo.Session) {
	chip.BootstrapRouter(router)
	mng.BootstrapRouter(router, mgoSession)

	chip.ServeRoot(router, publicRoot)
}
