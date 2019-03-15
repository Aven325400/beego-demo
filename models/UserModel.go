package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type User struct {
	Id       int    `orm:"column(Id)"`
	Email    string `orm:"size(50);unique;column(Email);default()"` //邮箱
	Password string `orm:"size(32);column(Password)"`               //密码
	Username string `orm:"size(16);column(Username);unique"`        //用户名
	Role     int    `orm:"column(Role)"`                            //角色
	Status   int8   `orm: column(Status)`                           //状态
}

func NewUser() *User {
	return &User{}
}

func GetUsertable() string {
	return getTable("user")
}
func Interface2Int(a interface{}) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%v", a))
	return i
}
func (u *User) UserList(p, listRows int, orderby, fields, cond string, args ...interface{}) (params []orm.Params, totalRows int, err error) {
	o := orm.NewOrm()
	if len(orderby) == 0 || orderby == "" {
		orderby = "u.Id desc"
	}
	if len(fields) == 0 {
		fields = "*"
	}
	if len(cond) > 0 {
		cond = "where " + cond
	}

	sqlCount := fmt.Sprintf("select count(u.id) cnt from %v u %v limit 1", GetUsertable(), cond)
	var one []orm.Params
	if rows, err := o.Raw(sqlCount, args).Values(&one); err == nil && rows > 0 {
		totalRows = Interface2Int(one[0]["cnt"])
	}

	sql := fmt.Sprintf("select %v from %v u %v order by %v limit %v offset %v",
		fields, GetUsertable(), cond, orderby, listRows, (p-1)*listRows)
	beego.Info(sql)
	_, err = o.Raw(sql, args...).Values(&params)
	return params, totalRows, err
}

func (u *User) FindUser(con *orm.Condition) (user User) {

	orm.NewOrm().QueryTable(GetUsertable()).SetCond(con).One(&user)

	return user
}

func (u *User) Reg(email, username, password string, role int, status int8) (error, int) {
	var (
		user User
		o    = orm.NewOrm()
	)
	if o.QueryTable(GetUsertable()).Filter("Username", username).One(&user); user.Id > 0 {
		return errors.New("用户名已经被注册"), 0
	}
	user = User{Email: email, Username: username, Password: password, Status: status, Role: role}
	_, err := o.Insert(&user)
	return err, user.Id

}
