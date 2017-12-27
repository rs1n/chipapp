package utils

import (
	"net/http"

	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
)

func BindRequest(w http.ResponseWriter, r *http.Request, result interface{}) {
	if err := render.BindJson(r, result); err != nil {
		RenderStatusAndAbort(w, http.StatusBadRequest)
	}
}

func BindRequestAndValidate(
	w http.ResponseWriter, r *http.Request, validate *validate.Validate,
	result interface{},
) {
	BindRequest(w, r, result)

	if err := validate.ValidateStruct(result); err != nil {
		RenderJsonAndAbort(w, http.StatusUnprocessableEntity, err)
	}
}
