package controllers

import (
	"fmt"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) LoginView() {
	this.TplName = "login/login.html"
}

func (this *LoginController) Login() {
	fmt.Println("ss")
	//username := this.Input().Get("username")
	//password := this.Input().Get("password")

	this.ServeJSON()
}
