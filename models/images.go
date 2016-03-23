package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Temp struct {
	Id         int       `orm:"column(id);"`
	Path       string    `orm:"column(path);size(255)"`
	Banned     int       `orm:"column(banned);int`
	CreateTime time.Time `orm:"column(create_time);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(
		new(Temp),
		new(Users))
}
