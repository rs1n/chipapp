package forms

import (
	"net/http"

	"github.com/sknv/chip/render"
)

// Base application form.
type Base struct{}

func bindJson(r *http.Request, result interface{}) error {
	return render.BindJson(r, result)
}
