package controllers

import (
	"encoding/json"
	"fmt"
	"move_kakou_platform/libs"
	"move_kakou_platform/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
}

var sessionName = beego.AppConfig.String("SessionName")

func (this *LoginController) LoginView() {
	beego.ReadFromRequest(&this.Controller)
	this.TplName = "login/login.html"
}

func (this *LoginController) Login() {
	flash := beego.NewFlash()
	errMsg := ""
	var user models.Users
	var ob models.Users
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	//backTo := this.GetString("back_to")

	o := orm.NewOrm()
	user.Username = ob.Username
	err := o.Read(&user, "username")
	fmt.Println(libs.Md5([]byte(ob.Password + user.Scope)))
	if err != nil || user.Password != libs.Md5([]byte(ob.Password+user.Scope)) {
		fmt.Println("error", err.Error())
		errMsg = "用户名或者密码错误"
	} else if user.Banned == 1 {
		errMsg = "用户已被禁用"
	} else {
		this.SetSession(sessionName, user.Id)
		//this.Redirect("/", 302)
		//sthis.Data["json"] = {}
		this.ServeJSON()
		return
	}
	flash.Error(errMsg)
	flash.Store(&this.Controller)
	this.Redirect("/login", 302)
	//this.ServeJSON()
}

var AuthRequest = func(ctx *context.Context) {
	//fmt.Println("sessionname", sessionName)
	id, ok := ctx.Input.Session(sessionName).(int)
	if !ok && ctx.Input.URI() != "/login" && ctx.Input.URI() != "/logout" {
		ctx.Redirect(302, "/login")
		return
	}
	var user models.Users
	var err error
	qs := orm.NewOrm()
	user.Id = id
	err = qs.Read(&user, "id")
	if err != nil {
		ctx.Redirect(302, "/login")
		return
	}
	//qs.LoadRelated()
	ctx.Input.SetData("user", user)
}
