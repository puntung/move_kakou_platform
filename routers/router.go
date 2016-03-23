package routers

import (
	"move_kakou_platform/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UploadController{}, "GET:IndexView")
	beego.Router("/login", &controllers.LoginController{}, "GET:LoginView;POST:Login")
	beego.Router("/upload", &controllers.UploadController{}, "GET:UploadView;POST:Upload")
	beego.Router("/admin", &controllers.UploadController{}, "GET:IndexView")
}
