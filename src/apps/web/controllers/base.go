package controllers

import (
	"net/http"

	"github.com/sknv/chipapp/src/apps"
	"github.com/sknv/chipapp/src/core/global"
)

// Base 'web' application controller.
type Base struct {
	apps.Controller
}

func (c *Base) RenderHtml(
	w http.ResponseWriter, status int, templateName string, data interface{},
) {
	g := global.GetGlobal()
	g.HTMLRenderer.Html(w, status, templateName, data)
}
