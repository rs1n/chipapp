package environment

import (
	"log"

	"github.com/rs1n/chip/render"
)

type (
	Environment struct {
		render.Html
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
		Html: htmlRenderer,
	}
}

func (env *Environment) CleanUp() {
	log.Println("Cleaning up...")
}
