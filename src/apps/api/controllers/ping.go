package controllers

import (
	"net/http"

	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
)

type Ping struct {
	*Base
}

func NewPing(validate *validate.Validate) *Ping {
	return &Ping{
		Base: NewBase(validate),
	}
}

func (c *Ping) Index(w http.ResponseWriter, _ *http.Request) {
	render.Json(w, http.StatusOK, render.M{
		"ping": "pong",
	})
}
