package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) ResponseJson(isSuccess bool, msg string, data ...interface{}) {
	code := 100
	if isSuccess {
		code = 200
	}
	ret := map[string]interface{}{"code": code, "message": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}
