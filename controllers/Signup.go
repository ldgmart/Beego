package controllers

import (
	"github.com/astaxie/beego"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	c.TplName = "Signup.tpl"
	c.Render()
}