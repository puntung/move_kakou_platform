package controllers

import (
	"move_kakou_platform/libs"
	"move_kakou_platform/models"
	"os"
	"time"

	"github.com/astaxie/beego/orm"
)

type UploadController struct {
	BaseController
}

func (this *UploadController) IndexView() {
	this.ConfigPage("kakou/kakou_info.html")
}

func (this *UploadController) UploadView() {
	this.ConfigPage("upload/upload.html")
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "upload/upload_script.html"
}

func (this *UploadController) Upload() {
	f, h, _ := this.GetFile("image")
	path := this.Ctx.Request.URL.EscapedPath()
	date_str := time.Now().Format("20060102")
	path = "./static/img" + path + "/" + date_str
	f.Close()
	if ok, _ := libs.PathExists(path); !ok {
		os.Mkdir(path, os.ModePerm)
	}
	err := this.SaveToFile("image", path+"/"+h.Filename)
	if err != nil {
		panic(err)
		return
	}
	o := orm.NewOrm()
	image := new(models.Temp)
	image.Path = path + "/" + h.Filename
	_, err = o.Insert(image)
	if err != nil {
		panic(err)
		return
	}
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "sucess"}
	this.ServeJSON()
}
