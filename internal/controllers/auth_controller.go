package controllers

import "net/http"

type authController struct {
}

func NewAuthController() *authController {
	return &authController{}
}

func (c *authController) Register(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}

func (c *authController) Login(_ http.ResponseWriter, _ *http.Request) {
	// TODO
}
