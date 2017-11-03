package controllers

import (
	"net/http"

	"github.com/rs1n/chip/render"
)

type Ping struct {
	*api
}

func NewPing() *Ping {
	return &Ping{
		api: &api{},
	}
}

func (c *Ping) Index(w http.ResponseWriter, _ *http.Request) {
	render.Json(w, http.StatusOK, render.M{
		"ping": "pong",
	})
}
