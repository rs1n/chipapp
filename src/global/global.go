package global

import (
	"log"

	"github.com/rs1n/chip/render"
)

type (
	Global struct {
		HtmlRenderer render.Html
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
	}
}

func (g *Global) CleanUp() {
	log.Println("Cleaning up...")
}
