package controllers

import (
	"github.com/sknv/chip/render"

	"github.com/sknv/chipapp/src/apps"
)

type Base struct {
	*apps.Controller `inject:""`

	HtmlRender *render.Html `inject:""`
}
