package controllers

import (
	"net/http"

	"github.com/sknv/chip/render"
)

type Ping struct {
	*Base `inject:""`
}

func (c *Ping) Index(w http.ResponseWriter, _ *http.Request) {
	render.Json(w, http.StatusOK, render.M{
		"ping": "pong",
	})
}
