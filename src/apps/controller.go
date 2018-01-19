package apps

import (
	"github.com/sknv/chip/validate"

	"github.com/sknv/chipapp/src/config"
	"github.com/sknv/chipapp/src/core/provider"
)

// Base controller.
type Controller struct {
	Config   *config.Config
	Validate *validate.Validate
}

func NewController() *Controller {
	objectProvider := provider.GetObjectProvider()
	return &Controller{
		Config:   objectProvider.Config,
		Validate: objectProvider.Validate,
	}
}
