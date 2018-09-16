package controllers

import (
	"github.com/ravenq/gvf-server/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ravenq/gvf-server/utils"
)

// BaseController auth base controller.
type BaseController struct {
	beego.Controller
	authMethods []string
}

// Init init
func (c *BaseController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	c.AuthInit()
}

// AuthInit set methed which need validate
func (c *BaseController) AuthInit() {
	c.authMethods = append(c.authMethods, "Post")
	c.authMethods = append(c.authMethods, "Put")
	c.authMethods = append(c.authMethods, "Delete")
}

// Prepare validate
func (c *BaseController) Prepare() {
	_, actionName := c.GetControllerAndAction()
	token := c.GetSession(utils.TOKEN)
	for _, i := range c.authMethods {
		if i == actionName {
			if token == nil {	
				c.Data["json"] = utils.FailResult(utils.ErrNeedLogin)
				c.ServeJSON()
			}

			if u, ok := token.(models.User); ok && !u.IsAdmin {
				c.Data["json"] = utils.FailResult(utils.ErrNeedAdmin)
				c.ServeJSON()
			}
			break	
		}
	}
}
