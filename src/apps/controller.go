package apps

import "github.com/skkv/chipapp/src/core/global"

// Base controller.
type Controller struct{}

func (c *Controller) ValidateStruct(s interface{}) error {
	g := global.GetGlobal()
	return g.Validate.Struct(s)
}
