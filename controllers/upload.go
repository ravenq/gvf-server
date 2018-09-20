package controllers

import (
	"github.com/astaxie/beego"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/ravenq/gvf-server/utils"
)

// UploadController update for qiniu
type UploadController struct {
	beego.Controller
}

// URLMapping ...
func (c *UploadController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get ...
// @Title Get token
// @Description get upload token
// @Success 200 {string} token
// @Failure 403 get token fail
// @router / [get]
func (c *UploadController) Get() {
	ak := beego.AppConfig.String("QINIU_AK")
	sk := beego.AppConfig.String("QINIU_SK")
	bt := beego.AppConfig.String("QINIU_BUCKET")
	mac := qbox.NewMac(ak, sk)
	putPolicy := storage.PutPolicy{
		Scope: bt,
	}
	upToken := putPolicy.UploadToken(mac)
	c.Data["json"] = utils.SuccessResult(upToken)
	c.ServeJSON()
}
