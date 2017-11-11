package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sknv/chip/render"
)

type Home struct {
	base
}

func (c *Home) Index(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		name = "world"
	}
	c.RenderHtml(w, http.StatusOK, "web/home/index", render.M{
		"pageTitle": "Welcome",
		"name":      name,
	})
}
