package main

import (
	"fmt"
	"github.com/ravenq/gvf-server/models"
	_ "github.com/ravenq/gvf-server/routers"
	"github.com/ravenq/gvf-server/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbHost := beego.AppConfig.String("DB_MYSQL_HOST")
	dbPwd := beego.AppConfig.String("DB_MYSQL_PASSWORD")
	cnnStr := fmt.Sprintf("root:%s@tcp(%s:3306)/myblog?charset=utf8", dbHost, dbPwd)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", cnnStr)
	orm.RunSyncdb("default", false, true)

	// init admin user
	u := models.User{
		Name:     beego.AppConfig.String("ADMIN_NAME"),
		Nick:     beego.AppConfig.String("ADMIN_NICK"),
		Password: utils.MD5(beego.AppConfig.String("ADMIN_PASSWORD")),
		Email:    beego.AppConfig.String("ADMIN_EMAIL"),
		IsAdmin:  true,
	}
	o := orm.NewOrm()
	o.Insert(&u)
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionEnableSidInHTTPHeader = true
	beego.BConfig.WebConfig.Session.SessionNameInHTTPHeader = "Session"
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
