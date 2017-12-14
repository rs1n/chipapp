package apps

import (
	"github.com/sknv/chip/validate"

	"github.com/sknv/chipapp/src/lib/services"
)

// Base controller.
type Controller struct {
	services.Request
	*validate.Validate
}

func NewController(validate *validate.Validate) *Controller {
	return &Controller{
		Validate: validate,
	}
}
