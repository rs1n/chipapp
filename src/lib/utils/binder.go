package utils

import (
	"log"
	"net/http"

	"github.com/sknv/chip/render"
)

type IStructValidator interface {
	ValidateStruct(s interface{}) error
}

func BindRequest(w http.ResponseWriter, r *http.Request, result interface{}) {
	if err := render.BindJson(r, result); err != nil {
		log.Print("error: ", err)
		RenderStatusAndAbort(w, http.StatusBadRequest)
	}
}

func BindRequestAndValidate(
	w http.ResponseWriter, r *http.Request, validate IStructValidator,
	result interface{},
) {
	BindRequest(w, r, result)

	if err := validate.ValidateStruct(result); err != nil {
		log.Print("error: ", err)
		RenderJsonAndAbort(w, http.StatusUnprocessableEntity, err)
	}
}
