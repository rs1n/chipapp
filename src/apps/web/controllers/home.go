package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs1n/chip/render"
)

type Home struct {
	*web
}

func NewHome() *Home {
	return &Home{
		web: &web{},
	}
}

func (c *Home) Index(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		name = "world"
	}
	render.Plain(w, http.StatusOK, "Hello, "+name)
}
