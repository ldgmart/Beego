package controllers

import (
	"github.com/astaxie/beego"
)


type StatusController struct {
	beego.Controller
}


func (this *StatusController) GetStatus() {
	se_manage := this.GetSession("session")
	beego.Info("sessions", se_manage)
	if se_manage == nil {
		this.Ctx.Output.SetStatus(200)
		this.Ctx.WriteString("NO")
		beego.Debug("Session not found : Status")
	return
	  }else{
		this.Ctx.Output.SetStatus(200)
		this.Ctx.WriteString("YES")
		beego.Debug("Session exists : Status")
	  }
}
