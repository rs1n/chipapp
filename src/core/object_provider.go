package core

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
	"github.com/sknv/mng"

	"github.com/sknv/chipapp/src/config"
	"github.com/sknv/chipapp/src/lib/repositories"
)

type (
	HtmlRenderParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}

	ObjectProvider struct {
		MgoSession *mgo.Session

		// Dependencies to be injected.
		Objects []interface{}
	}
)

func NewObjectProvider(hrp HtmlRenderParams, config *config.Config) *ObjectProvider {
	htmlRender := &render.Html{
		IsDebug:      hrp.IsDebug,
		TemplateRoot: hrp.TemplateRoot,
		TemplateExt:  hrp.TemplateExt,
	}
	mgoSession := mng.MustDial(config.Mongo)

	// TODO: Provide complete dependencies to be injected here.
	objects := []interface{}{
		config,
		htmlRender,
		validate.NewValidate(nil),
		repositories.NewUser(),
	}

	return &ObjectProvider{
		Objects:    objects,
		MgoSession: mgoSession,
	}
}

func (op *ObjectProvider) CleanUp() {
	log.Println("Cleaning up...")
	op.cleanMongo()
}

func (op *ObjectProvider) cleanMongo() {
	if op.MgoSession != nil {
		op.MgoSession.Close()
	}
}
