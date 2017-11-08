package global

import (
	"log"

	"github.com/rs1n/chip/render"
	"github.com/rs1n/chip/validate"
)

type (
	Global struct {
		HtmlRenderer render.Html
		Validate     *validate.Validate
	}

	HtmlRendererParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}
)

func NewGlobal(rhp HtmlRendererParams) *Global {
	htmlRenderer := render.Html{
		IsDebug:      rhp.IsDebug,
		TemplateRoot: rhp.TemplateRoot,
		TemplateExt:  rhp.TemplateExt,
	}

	return &Global{
		HtmlRenderer: htmlRenderer,
		Validate:     validate.NewValidate(),
	}
}

func (g *Global) CleanUp() {
	log.Println("Cleaning up...")
}
