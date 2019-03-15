package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

func init() {
	beego.Info(beego.AppConfig.String("db::password"))
	RegisterDB()
	orm.RunSyncdb("default", false, true)
}
func RegisterDB() {
	models := []interface{}{
		NewUser(),
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModelWithPrefix(beego.AppConfig.DefaultString("db::prefix", "gt_"), models...)
	dbUser := beego.AppConfig.String("db::user")
	dbPassword := beego.AppConfig.String("db::password")
	dbDatabase := beego.AppConfig.String("db::database")
	dbCharset := beego.AppConfig.String("db::charset")
	dbHost := beego.AppConfig.String("db::host")
	dbPost := beego.AppConfig.String("db::port")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%v", dbUser, dbPassword, dbHost, dbPost, dbDatabase, dbCharset, "Asia%2FShanghai")
	//beego.Info(dbDatabase, dbPost, dbUser, dbCharset)
	//fmt.Print(dbLink)
	maxIdle := beego.AppConfig.DefaultInt("db::maxIdle", 50)
	maxConn := beego.AppConfig.DefaultInt("db::maxConn", 300)
	if err := orm.RegisterDataBase("default", "mysql", dbLink, maxIdle, maxConn); err != nil {
		panic(err)
	}

}

func getTable(table string) string {
	prefix := beego.AppConfig.DefaultString("db::prefix", "hc_")
	if !strings.HasPrefix(table, prefix) {
		table = prefix + table
	}
	return table
}
