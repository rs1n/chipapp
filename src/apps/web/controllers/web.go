package controllers

import (
	"net/http"

	"github.com/rs1n/chipapp/src/global"
)

// Base 'web' application controller.
type web struct{}

func (c *web) Html(
	w http.ResponseWriter, status int, templateName string, data interface{},
) {
	g := global.GetGlobal()
	g.HtmlRenderer.Html(w, status, templateName, data)
}
