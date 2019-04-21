package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/ravenq/gvf-server/models"
	"github.com/ravenq/gvf-server/utils"
)

// CommentsController operations for Post
type CommentsController struct {
	BaseController
}

// URLMapping ...
func (c *CommentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Fetch", c.Fetch)
}


// Prepare ...
func (c *CommentsController) Prepare() {
	_, actionName := c.GetControllerAndAction()
	token := c.GetSession(utils.TOKEN)
	if "Post" == actionName {
		if token == nil {
			c.Data["json"] = utils.FailResult(utils.ErrNeedLogin)
			c.ServeJSON()
		}
		return
	}

	c.BaseController.Prepare()
}

// Post ...
// @Title Post
// @Description create Comments
// @Param	body		body 	models.Comments	true		"body for Comments content"
// @Success 201 {int} models.Comments
// @Failure 403 body is empty
// @router / [post]
func (c *CommentsController) Post() {
	var v models.Comments
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	v.RemoteAddr = utils.RemoteIP(c.Ctx.Request)
	_, err := models.AddComments(&v)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Comments by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comments
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommentsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCommentsById(id)
	c.Data["json"] = utils.NewResult(v, err)
	c.ServeJSON()
}

// Fetch ...
// @Title Fetch
// @Description get Comments by comment id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comments
// @Failure 403 : id is empty
// @router /fetch/:id [get]
func (c *CommentsController) Fetch() {
	idStr := c.Ctx.Input.Param(":id")
	
	l, err := models.GetCommentsByCommentId(idStr)

	// sort
	m := make(map[int64]*models.Comments)
	for i, v := range l {
		m[v.Id] = &l[i]
	}

	var ml []*models.Comments
	for _, v := range l {
		if v.Parent < 1 {
			ml = append(ml, m[v.Id])
		} else {
			if p, ok := m[v.Parent]; ok {
				p.Replies = append(p.Replies, v)
			}
		}
	}

	c.Data["json"] = utils.NewResult(ml, err)
	c.ServeJSON()
}
