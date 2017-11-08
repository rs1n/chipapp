package controllers

import (
	"net/http"

	"github.com/rs1n/chipapp/src/core/global"
)

// Base 'web' application controller.
type base struct{}

func (c *base) Html(
	w http.ResponseWriter, status int, templateName string, data interface{},
) {
	g := global.GetGlobal()
	g.HtmlRenderer.Html(w, status, templateName, data)
}
