package controllers

import (
	"github.com/sknv/chip/render"

	"github.com/sknv/chipapp/src/apps"
	"github.com/sknv/chipapp/src/core/provider"
)

type BaseController struct {
	*apps.Controller

	HtmlRender *render.Html
}

func NewBaseController() *BaseController {
	objectProvider := provider.GetObjectProvider()
	return &BaseController{
		Controller: apps.NewController(),
		HtmlRender: objectProvider.HtmlRender,
	}
}
