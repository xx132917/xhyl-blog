package controllers

import (
	"errors"
	"strconv"
	"strings"
	"xhylblog/models"
)

//  ArticleTypeController operations for ArticleType
type ArticleTypeController struct {
	BaseController
}

// URLMapping ...
func (c *ArticleTypeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ArticleType
// @Param	body		body 	models.ArticleType	true		"body for ArticleType content"
// @Success 201 {int} models.ArticleType
// @Failure 403 body is empty
// @router / [post]
func (c *ArticleTypeController) Post() {
	var v models.ArticleType
	c.JsonParam(&v)
	if _, err := models.AddArticleType(&v); err == nil {
		c.Ok(v)
	} else {
		c.Error(err.Error())
	}
}

// GetOne ...
// @Title Get One
// @Description get ArticleType by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ArticleType
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ArticleTypeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetArticleTypeById(id)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Ok(v)
	}
}

// GetAll ...
// @Title Get All
// @Description get ArticleType
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ArticleType
// @Failure 403
// @router / [get]
func (c *ArticleTypeController) GetAll() {
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
				c.Error(errors.New("Error: invalid query key/value pair"))
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllArticleType(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Error(err.Error())
	} else {
		c.Ok(l)
	}
}

// Put ...
// @Title Put
// @Description update the ArticleType
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ArticleType	true		"body for ArticleType content"
// @Success 200 {object} models.ArticleType
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ArticleTypeController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.ArticleType{Id: id}
	c.JsonParam(&v)
	if err := models.UpdateArticleTypeById(&v); err == nil {
		c.Ok("OK")
	} else {
		c.Error(err.Error())
	}
}

// Delete ...
// @Title Delete
// @Description delete the ArticleType
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ArticleTypeController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteArticleType(id); err == nil {
		c.Ok("OK")
	} else {
		c.Error(err.Error())
	}
}
