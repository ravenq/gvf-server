package controllers

import (
	"github.com/astaxie/beego"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/ravenq/gvgo/conf"
	"github.com/ravenq/gvgo/utils"
)

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
	mac := qbox.NewMac(conf.QINIU_AK, conf.QINIU_SK)
	putPolicy := storage.PutPolicy{
		Scope: conf.QINIU_BUCKET,
	}
	upToken := putPolicy.UploadToken(mac)
	c.Data["json"] = utils.SuccessResult(upToken)
	c.ServeJSON()
}
