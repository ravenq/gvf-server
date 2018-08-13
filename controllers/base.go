package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"github.com/ravenq/gvgo/utils"
)

// BaseController auth base controller.
type BaseController struct {
	beego.Controller
	authMethods []string
}

func (c *BaseController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
	c.AuthInit()
}

func (c *BaseController) AuthInit() {
	c.authMethods = append(c.authMethods, "GetAll")
	c.authMethods = append(c.authMethods, "Post")
	c.authMethods = append(c.authMethods, "Put")
	c.authMethods = append(c.authMethods, "Delete")
}

// Prepare validate
func (c *BaseController) Prepare() {
	_, actionName := c.GetControllerAndAction()
	sess, _ := utils.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	token := sess.Get(utils.TOKEN)
	for _, i := range c.authMethods {
		if i == actionName {
			if token == nil {
				c.Data["json"] = utils.FailResult(utils.ErrNeedLogin)
				c.ServeJSON()
			}
			break
		}
	}
}
