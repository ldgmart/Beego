package main

import (
	"github.com/astaxie/beego"
	_ "sugoee/routers"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func main() {
	//beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
	sessionconf := &session.ManagerConfig{
		CookieName: "beegosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()
	beego.Run()
}
