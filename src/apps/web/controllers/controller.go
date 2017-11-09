package controllers

import (
	"net/http"

	"github.com/rs1n/chipapp/src/apps"
	"github.com/rs1n/chipapp/src/core/global"
)

// Base 'web' application controller.
type controller struct {
	apps.Controller
}

func (c *controller) RenderHtml(
	w http.ResponseWriter, status int, templateName string, data interface{},
) {
	g := global.GetGlobal()
	g.HtmlRenderer.Html(w, status, templateName, data)
}
