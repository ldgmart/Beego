package controllers

import (
	"github.com/astaxie/beego"
     "sugoee/models"
)

type LoginController struct {
	beego.Controller
}
func render_Login(this *LoginController, tpl string, msg string){
	this.TplName = tpl
	this.Data["Content"] = msg
	this.Render()
}
func (this *LoginController) Login() {

	//session := this.StartSession()
	//UserID := session.Get("UserID")
	/*
	if UserID != nil {
		// User is logged in already, display another page
		return
	}*/
	id := this.Ctx.Input.Query("id")
	if id == "" {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("Please fill out id")
		return
	}

	password := this.Ctx.Input.Query("password")
	if password == "" {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("Please fill out password")
		return
	}
	m, err := models.Login(id, password)

	if err != nil {
		beego.Info(m, err)
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("No match result")
		//render_Login(this, "alert.tpl", "No match result")
		return
	}else{

	this.Ctx.Output.SetStatus(200)
	this.Ctx.WriteString("Success")
	se_manage := make(map[string]interface{})
	se_manage["ID"] = m.ID
	beego.Info("In SigninController:Post - Creating new session", se_manage)
	this.SetSession("session", se_manage)
}
}

