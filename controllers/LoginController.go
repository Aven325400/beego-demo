package controllers

import (
	"github.com/astaxie/beego"
	"server-demo/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	var post struct {
		Password string
		Username string
	}
	err := this.ParseForm(&post)
	beego.Info(err, post)

	userModel := models.NewUser()
	users, rows, er := userModel.UserList(1, 1, "", "",
		"u.`username` = ? and u.`password` = ?", post.Username, post.Password)
	if rows == 0 || err != nil {
		if err != nil {
			beego.Info(er)
		}
		this.ResponseJson(false, "登录失败，用户名或密码不正确")
	}
	user := users[0]
	beego.Info(user)
	this.ResponseJson(true, "登录成功")

}
