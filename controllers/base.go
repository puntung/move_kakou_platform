package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (base *BaseController) Prepare() {
	controllername, actionname := base.GetControllerAndAction()
	base.controllerName = strings.ToLower(controllername[0 : len(controllername)-10])
	base.actionName = strings.ToLower(actionname)
	base.Data["appname"] = beego.AppConfig.String("appname")
}

func (base *BaseController) ConfigPage(name string) {
	if len(name) == 0 {
		base.Layout = "index.html"
		base.TplName = base.controllerName + "/" + base.actionName + ".html"
	} else {
		base.Layout = "index.html"
		base.TplName = name
	}
}
