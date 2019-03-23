package main

import (
	_ "export_excel/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init()  {
	host := beego.AppConfig.String("db_host")
	port := beego.AppConfig.String("db_port")
	database:=beego.AppConfig.String("db_database")
	username:=beego.AppConfig.String("db_username")
	password:=beego.AppConfig.String("db_password")
	dsn:=username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&loc=Asia%2FShanghai"

	orm.RegisterDataBase("default", "mysql", dsn, 30)

	db := orm.NewOrm()
	db.Using("default")
}

func main() {
	orm.Debug = true
	beego.Run()
}

