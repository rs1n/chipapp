package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"
)

type Home struct {
	*Base
}

func NewHome(htmlRender *render.Html, validate *validate.Validate) *Home {
	return &Home{
		Base: NewBase(htmlRender, validate),
	}
}

func (c *Home) Index(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		name = "world"
	}
	c.HtmlRender.Html(w, http.StatusOK, "web/home/index", render.M{
		"pageTitle": "Welcome",
		"name":      name,
	})
}
