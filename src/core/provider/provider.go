package provider

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
	"github.com/sknv/mng"

	"github.com/sknv/chipapp/src/config"
)

var objectProvider *ObjectProvider

type (
	HtmlRenderParams struct {
		IsDebug      bool
		TemplateRoot string
		TemplateExt  string
	}

	ObjectProvider struct {
		Config     *config.Config
		HtmlRender *render.Html
		MgoSession *mgo.Session
		Validate   *validate.Validate
	}
)

func GetObjectProvider() *ObjectProvider {
	if objectProvider == nil {
		panic("objectProvider is not initialized")
	}
	return objectProvider
}

func NewObjectProvider(hrp HtmlRenderParams, config *config.Config) *ObjectProvider {
	if objectProvider != nil {
		panic("objectProvider is already initialized")
	}

	htmlRender := &render.Html{
		IsDebug:      hrp.IsDebug,
		TemplateRoot: hrp.TemplateRoot,
		TemplateExt:  hrp.TemplateExt,
	}

	objectProvider = &ObjectProvider{
		Config:     config,
		HtmlRender: htmlRender,
		MgoSession: mng.MustDial(config.Mongo),
		Validate:   validate.NewValidate(nil), // Use a validator with the default translator.
	}
	return objectProvider
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
