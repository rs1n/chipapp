package apps

import (
	"github.com/sknv/chip/validate"

	"github.com/sknv/chipapp/src/config"
)

// Base controller.
type Controller struct {
	*config.Config     `inject:""`
	*validate.Validate `inject:""`
}
