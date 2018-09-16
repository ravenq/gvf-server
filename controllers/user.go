package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/ravenq/gvf-server/models"
	"github.com/ravenq/gvf-server/utils"
)

// UserController operations for User
type UserController struct {
	BaseController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
}

// Login ...
// @Title Login
// @Description login
// @param body {username: "", password: ""}
// @success 201 {int} models.User
// @Failure 403 body is empty
// @router /login [post]
func (c *UserController) Login() {
	var p models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &p)
	p.Password = utils.MD5(p.Password)
	v, err := models.GetUserByName(p.Name)
	if err != nil {
		v, err = models.GetUserByEmail(p.Name)
	}

	if err != nil {
		c.Data["json"] = utils.FailResult(utils.ErrUserNotExist)
	} else {
		if p.Password != v.Password {
			c.Data["json"] = utils.FailResult(utils.ErrPasswordError)
		} else {
			//v.Token = utils.UUID()
			c.SetSession(utils.TOKEN, v)
			v.Token = c.CruSession.SessionID()
			c.Data["json"] = utils.NewResult(v, nil)
		}
	}
	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	v.Password = utils.MD5(v.Password)
	v.IsAdmin = false
	_, err := models.AddUser(&v)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserById(id)
	c.Data["json"] = utils.NewResult(v, err)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
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

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	c.Data["json"] = utils.NewResult(l, err)
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.User{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	err := models.UpdateUserById(&v)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	err := models.DeleteUser(id)
	c.Data["json"] = utils.NewEmptyResult(err)
	c.ServeJSON()
}
