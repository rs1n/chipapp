package controllers

import (
	"github.com/sknv/chip/validate"

	"github.com/sknv/chipapp/src/apps"
)

// Base 'api' application controller.
type Base struct {
	*apps.Controller
}

func NewBase(validate *validate.Validate) *Base {
	return &Base{
		Controller: apps.NewController(validate),
	}
}
