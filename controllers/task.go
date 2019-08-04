package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
     "sugoee/models"
)

type TaskController struct {
	beego.Controller
}

func render(this *TaskController, tpl string, msg string){
	this.TplName = tpl
	this.Data["Content"] = msg
	this.Render()
}

func (this *TaskController) NewSignup() {

	//this.TplName = "index.html"
	//this.Render()
	
id := this.Ctx.Input.Query("id")
	if id == "" {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("Please fill out this field")
		return
	}

	password := this.Ctx.Input.Query("password")
	if password == "" {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("Please fill out this field")
		return
	}
	email := this.Ctx.Input.Query("email")
	if email == "" {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("Please fill out this field")
		return
	}
	FridaLog:=logs.GetLogger()
	FridaLog.Println("\nid : " + id, "\npassword " + password + "\nemail " + email)
	
	Flag, err := models.IsID(id)
	if Flag == true {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("Duplicate ID")
		return
	}

	m, err := models.Signup(id, password, email)
	if err != nil {
		this.Ctx.Output.SetStatus(401)
		this.Ctx.WriteString("Error occurred")
		return
	}else{
		this.Ctx.Output.SetStatus(200)
		this.Ctx.WriteString("Success")
	}
	beego.Info(m)
}
