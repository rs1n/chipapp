package core

import (
	"log"

	"github.com/sknv/chip"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
	"upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

type (
	HtmlRenderParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}

	ServiceProvider struct {
		HtmlRender *render.Html
		*validate.Validate
		PgSession  db.Database
	}
)

func NewServiceProvider(
	hrp HtmlRenderParams, connectionURL *postgresql.ConnectionURL,
) *ServiceProvider {
	htmlRenderer := &render.Html{
		IsDebug:      hrp.IsDebug,
		TemplateRoot: hrp.TemplateRoot,
		TemplateExt:  hrp.TemplateExt,
	}

	pgSession, err := postgresql.Open(connectionURL)
	chip.PanicIfError(err)

	return &ServiceProvider{
		HtmlRender: htmlRenderer,
		Validate:   validate.NewValidate(),
		PgSession:  pgSession,
	}
}

func (sp *ServiceProvider) CleanUp() {
	log.Println("Cleaning up...")
	sp.cleanPostgres()
}

func (sp *ServiceProvider) cleanPostgres() {
	if sp.PgSession != nil {
		sp.PgSession.Close()
	}
}
