package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/ravenq/gvf-server/models"
	"github.com/ravenq/gvf-server/utils"
)

// PostController operations for Post
type PostController struct {
	BaseController
}

// URLMapping ...
func (c *PostController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Count", c.Count)
}

// Post ...
// @Title Post
// @Description create Post
// @Param	body		body 	models.Post	true		"body for Post content"
// @Success 201 {int} models.Post
// @Failure 403 body is empty
// @router / [post]
func (c *PostController) Post() {
	var v models.Post
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	_, err := models.AddPost(&v)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}

// Count ...
// @Title Count
// @Description Count
// @success 201 {int} models.Post
// @Failure 403 body is empty
// @router /count [get]
func (c *PostController) Count() {
	count, err := models.PostCount()
	c.Data["json"] = utils.NewResult(count, err)
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Post by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Post
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PostController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPostById(id)
	c.Data["json"] = utils.NewResult(v, err)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Post
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Post
// @Failure 403
// @router / [get]
func (c *PostController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllPost(query, fields, sortby, order, offset, limit)
	c.Data["json"] = utils.NewResult(l, err)
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Post
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Post	true		"body for Post content"
// @Success 200 {object} models.Post
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PostController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Post{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	err := models.UpdatePostById(&v)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Post
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PostController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	err := models.DeletePost(id)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}
