package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello");
}

func (c *MainController) Sql2() {
	c.Ctx.WriteString("sql2");
}
