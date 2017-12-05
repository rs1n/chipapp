package global

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
	"github.com/sknv/mng"
)

type (
	Global struct {
		HtmlRenderer *render.Html
		Validate     *validate.Validate
		MgoSession   *mgo.Session
	}

	HtmlRenderParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}
)

func NewGlobal(hrp HtmlRenderParams, mgoDialInfo *mgo.DialInfo) *Global {
	htmlRenderer := &render.Html{
		IsDebug:      hrp.IsDebug,
		TemplateRoot: hrp.TemplateRoot,
		TemplateExt:  hrp.TemplateExt,
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
