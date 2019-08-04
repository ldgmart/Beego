package controllers

import (
	"github.com/astaxie/beego"
)


type MainController struct {
	beego.Controller
}


func (this *MainController) Get() {
	this.TplName = "index.html"
	this.Render()
	se_manage := this.GetSession("session")
	if se_manage == nil {
		beego.Debug("Session not found : Default")
	return
	  }else{
		  beego.Debug("Session exists : Default")
	  }
}
