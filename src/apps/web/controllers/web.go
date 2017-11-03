package controllers

import (
	"net/http"

	"github.com/rs1n/chipapp/src/environment"
)

// Base 'web' application controller.
type web struct{}

func (c *web) Html(
	w http.ResponseWriter, status int, templateName string, data interface{},
) {
	env := environment.GetEnvironment()
	env.HtmlRenderer.Html(w, status, templateName, data)
}
