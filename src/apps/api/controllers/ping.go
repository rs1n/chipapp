package controllers

import (
	"net/http"

	"github.com/skkv/chip/render"
)

type Ping struct {
	base
}

func (c *Ping) Index(w http.ResponseWriter, _ *http.Request) {
	render.Json(w, http.StatusOK, render.M{
		"ping": "pong",
	})
}
