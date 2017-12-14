package controllers

import (
	"github.com/sknv/chip/render"
	"github.com/sknv/chip/validate"

	"github.com/sknv/chipapp/src/apps"
)

// Base 'web' application controller.
type Base struct {
	*apps.Controller

	HtmlRender *render.Html
}

func NewBase(htmlRender *render.Html, validate *validate.Validate) *Base {
	return &Base{
		Controller: apps.NewController(validate),
		HtmlRender: htmlRender,
	}
}
