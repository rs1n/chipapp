package utils

import (
	"net/http"

	"github.com/sknv/chip"
	"github.com/sknv/chip/render"
)

func RenderJsonAndAbort(w http.ResponseWriter, status int, data interface{}) {
	render.Json(w, status, data)
	chip.AbortHandler()
}

func RenderStatusAndAbort(w http.ResponseWriter, status int) {
	render.Status(w, status)
	chip.AbortHandler()
}

func RenderPlainAndAbort(w http.ResponseWriter, status int, data string) {
	render.Plain(w, status, data)
	chip.AbortHandler()
}
