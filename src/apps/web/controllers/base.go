package controllers

import (
	"net/http"

	"github.com/skkv/chipapp/src/apps"
	"github.com/skkv/chipapp/src/core/global"
)

// Base 'web' application controller.
type base struct {
	apps.Controller
}

func (c *base) RenderHtml(
	w http.ResponseWriter, status int, templateName string, data interface{},
) {
	g := global.GetGlobal()
	g.HtmlRenderer.Html(w, status, templateName, data)
}
