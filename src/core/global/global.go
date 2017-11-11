package global

import (
	"log"

	"github.com/sknv/chip/mng"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
	"gopkg.in/mgo.v2"
)

type (
	Global struct {
		HtmlRenderer *render.Html
		Validate     *validate.Validate
		MgoSession   *mgo.Session
	}

	HtmlRendererParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}
)

func NewGlobal(rhp HtmlRendererParams, mgoDialInfo *mgo.DialInfo) *Global {
	htmlRenderer := &render.Html{
		IsDebug:      rhp.IsDebug,
		TemplateRoot: rhp.TemplateRoot,
		TemplateExt:  rhp.TemplateExt,
	}

	return &Global{
		HtmlRenderer: htmlRenderer,
		Validate:     validate.NewValidate(),
		MgoSession:   mng.MustDial(mgoDialInfo),
	}
}

func (g *Global) CleanUp() {
	log.Println("Cleaning up...")
	g.cleanMongo()
}

func (g *Global) cleanMongo() {
	if g.MgoSession != nil {
		g.MgoSession.Close()
	}
}
