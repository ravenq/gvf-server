package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/ravenq/gvgo/models"
	"github.com/ravenq/gvgo/utils"
	//"github.com/astaxie/beego"
)

// MyloveController operations for Mylove
type MyloveController struct {
	BaseController
}

// URLMapping ...
func (c *MyloveController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Mylove
// @Param	body		body 	models.Mylove	true		"body for Mylove content"
// @Success 201 {int} models.Mylove
// @Failure 403 body is empty
// @router / [post]
func (c *MyloveController) Post() {
	var v models.Mylove
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	_, err := models.AddMylove(&v)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Mylove by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Mylove
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MyloveController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetMyloveById(id)
	c.Data["json"] = utils.NewResult(v, err)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Mylove
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Mylove
// @Failure 403
// @router / [get]
func (c *MyloveController) GetAll() {
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

	l, err := models.GetAllMylove(query, fields, sortby, order, offset, limit)
	c.Data["json"] = utils.NewResult(l, err)
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Mylove
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Mylove	true		"body for Mylove content"
// @Success 200 {object} models.Mylove
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MyloveController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Mylove{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	err := models.UpdateMyloveById(&v)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Mylove
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MyloveController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	err := models.DeleteMylove(id)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}
