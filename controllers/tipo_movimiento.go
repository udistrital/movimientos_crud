package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/udistrital/movimientos_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// TipoMovimientoController operations for TipoMovimiento
type TipoMovimientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoMovimientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoMovimiento
// @Param	body		body 	models.TipoMovimiento	true		"body for TipoMovimiento content"
// @Success 201 {int} models.TipoMovimiento
// @Failure 403 body is empty
// @router / [post]
func (c *TipoMovimientoController) Post() {
	var v models.TipoMovimiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTipoMovimiento(&v); err == nil {
			// logs.Info(fmt.Sprintf("Éxito: %v+", v))
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(fmt.Sprintf("Error1: %v+, error: %v+", v, err))
			c.Data["json"] = err.Error()
		}
	} else {
		logs.Error(fmt.Sprintf("Error2: %v+", v))
		c.Data["json"] = err.Error()
	}
}

// GetOne ...
// @Title Get One
// @Description get TipoMovimiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoMovimiento
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoMovimientoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoMovimientoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
}

// GetAll ...
// @Title Get All
// @Description get TipoMovimiento
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoMovimiento
// @Failure 403
// @router / [get]
func (c *TipoMovimientoController) GetAll() {
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

	l, err := models.GetAllTipoMovimiento(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
}

// Put ...
// @Title Put
// @Description update the TipoMovimiento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoMovimiento	true		"body for TipoMovimiento content"
// @Success 200 {object} models.TipoMovimiento
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoMovimientoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.TipoMovimiento{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTipoMovimientoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
}

// Delete ...
// @Title Delete
// @Description delete the TipoMovimiento
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoMovimientoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoMovimiento(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
}
