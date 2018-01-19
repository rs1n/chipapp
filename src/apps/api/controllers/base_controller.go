package controllers

import "github.com/sknv/chipapp/src/apps"

type BaseController struct {
	*apps.Controller
}

func NewBaseController() *BaseController {
	return &BaseController{apps.NewController()}
}
