package controllers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sknv/chip/render"
)

type HomeController struct {
	*BaseController
}

func NewHomeController() *HomeController {
	return &HomeController{NewBaseController()}
}

func (c *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		name = "world"
	}
	c.HtmlRender.Html(w, http.StatusOK, "web/home/index", render.M{
		"pageTitle": "Welcome",
		"name":      name,
	})
}
