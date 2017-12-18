package core

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
	"github.com/sknv/mng"
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
		MgoSession *mgo.Session
	}
)

func NewServiceProvider(
	hrp HtmlRenderParams, mgoDialInfo *mgo.DialInfo,
) *ServiceProvider {
	htmlRender := &render.Html{
		IsDebug:      hrp.IsDebug,
		TemplateRoot: hrp.TemplateRoot,
		TemplateExt:  hrp.TemplateExt,
	}

	return &ServiceProvider{
		HtmlRender: htmlRender,
		Validate:   validate.NewValidate(),
		MgoSession: mng.MustDial(mgoDialInfo),
	}
}

func (sp *ServiceProvider) CleanUp() {
	log.Println("Cleaning up...")
	sp.cleanMongo()
}

func (sp *ServiceProvider) cleanMongo() {
	if sp.MgoSession != nil {
		sp.MgoSession.Close()
	}
}
