package apps

import "github.com/sknv/chipapp/src/lib/services"

// Base controller.
type Controller struct {
	services.Request
	services.Validate
}
