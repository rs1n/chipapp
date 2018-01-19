package controllers

import (
	"net/http"

	"github.com/sknv/chip/render"
)

type PingController struct {
	*BaseController
}

func NewPingController() *PingController {
	return &PingController{NewBaseController()}
}

func (c *PingController) Index(w http.ResponseWriter, _ *http.Request) {
	render.Json(w, http.StatusOK, render.M{
		"ping": "pong",
	})
}
