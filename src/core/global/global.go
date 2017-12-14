package global

import (
	"log"

	"github.com/sknv/chip"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
	"upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

type (
	Global struct {
		HTMLRenderer *render.Html
		Validate     *validate.Validate
		PgSession    db.Database
	}

	HtmlRenderParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}
)

func NewGlobal(
	hrp HtmlRenderParams, connectionURL *postgresql.ConnectionURL,
) *Global {
	htmlRenderer := &render.Html{
		IsDebug:      hrp.IsDebug,
		TemplateRoot: hrp.TemplateRoot,
		TemplateExt:  hrp.TemplateExt,
	}

	pgSession, err := postgresql.Open(connectionURL)
	chip.PanicIfError(err)

	return &Global{
		HTMLRenderer: htmlRenderer,
		Validate:     validate.NewValidate(),
		PgSession:    pgSession,
	}
}

func (g *Global) CleanUp() {
	log.Println("Cleaning up...")
	g.cleanPostgres()
}

func (g *Global) cleanPostgres() {
	if g.PgSession != nil {
		g.PgSession.Close()
	}
}
