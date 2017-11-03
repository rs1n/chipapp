package environment

import (
	"log"

	"github.com/rs1n/chip/render"
)

type (
	Environment struct {
		HtmlRenderer render.Html
	}

	HtmlRendererParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}
)

func NewEnvironment(rhp HtmlRendererParams) *Environment {
	htmlRenderer := render.Html{
		IsDebug:      rhp.IsDebug,
		TemplateRoot: rhp.TemplateRoot,
		TemplateExt:  rhp.TemplateExt,
	}

	return &Environment{
		HtmlRenderer: htmlRenderer,
	}
}

func (env *Environment) CleanUp() {
	log.Println("Cleaning up...")
}
