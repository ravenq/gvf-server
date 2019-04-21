package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ravenq/gvf-server/models"
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
	c.authMethods = append(c.authMethods, "Post")
	c.authMethods = append(c.authMethods, "Put")
	c.authMethods = append(c.authMethods, "Delete")
}

// MappingAuth ...
func (c *BaseController) MappingAuth(m string) {
	c.authMethods = append(c.authMethods, m)
}

// UnMappingAuth ...
func (c *BaseController) UnMappingAuth(m string) {
	ret := make([]string, 0, len(c.authMethods))
	for _, val := range c.authMethods {
		if val != m {
			ret = append(ret, val)
		}
	}
	c.authMethods = ret
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
