package main

import (
	_ "move_kakou_platform/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	maxIdle := 300
	maxConn := 300
	if dbport == "" {
		dbport = "3306"
	}
	DBAdrr := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	err := orm.RegisterDataBase(
		"default",
		"mysql",
		DBAdrr,
		maxIdle,
		maxConn)
	if err != nil {
		panic(err)
	}
}

func main() {
	beego.Run()
}
