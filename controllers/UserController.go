package controllers

import (
	"github.com/astaxie/beego"
	"server-demo/models"
)

type UserController struct {
	BaseController
}

func (this *UserController) Register() {
	var post struct {
		Username string
		Password string
		Email    string
		Role     int
		Status   int8
	}
	this.ParseForm(&post)
	user := models.NewUser()
	err, userId := user.Reg(post.Email, post.Username, post.Password, post.Role, post.Status)
	if err != nil || userId == 0 {
		beego.Info(err)
		this.ResponseJson(false, "注册失败")
	}
	this.ResponseJson(true, "注册成功")
}
