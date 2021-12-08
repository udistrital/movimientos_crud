package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	movimientoProcesoExternoManager "github.com/udistrital/movimientos_crud/managers/movimientoProcesoExternoManager"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/responseformat"
)

// MovimientoProcesoExternoController operations for MovimientoProcesoExterno
type MovimientoProcesoExternoController struct {
	beego.Controller
}

// URLMapping ...
func (c *MovimientoProcesoExternoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Regitrar", c.RegistrarMovimiento)
}

// RegistrarMovimiento ...
// @Title RegistrarMovimiento
// @Description Registra un movimiento completamente, tanto el de proceso externo como el detalle
// @Param	body		body 	map[string]interface{}	true		"map[string]interface{}"
// @Success 201 {int} models.MovimientoProcesoExterno
// @Failure 403 body is empty
// @router /registrar_movimiento [post]
func (c *MovimientoProcesoExternoController) RegistrarMovimiento() {
	registrarMovimientoCommon(c)
}

// RegistrarMovimientoOld ...
// @Title RegistrarMovimientoOld (deprecated/old/wrong path)
// @Description Registra un movimiento completamente, tanto el de proceso externo como el detalle (deprecated/old/wrong path!)
// @Param	body		body 	map[string]interface{}	true		"map[string]interface{}"
// @Success 201 {int} models.MovimientoProcesoExterno
// @Failure 403 body is empty
// @router registrar_movimiento [post]
func (c *MovimientoProcesoExternoController) RegistrarMovimientoOld() {
	registrarMovimientoCommon(c)
}

func registrarMovimientoCommon(c *MovimientoProcesoExternoController) {
	var movimiento map[string]interface{}

	layoutDate := "2006-01-02"
	dataResponse := make(map[string]interface{})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &movimiento); err == nil {
		tipoMovimiento := models.TipoMovimiento{Id: int(movimiento["TipoMovimiento"].(float64))}
		movimientoProcesoExterno := models.MovimientoProcesoExterno{
			TipoMovimientoId: &tipoMovimiento,
			ProcesoExterno:   int64(movimiento["ProcesoExterno"].(float64)),
		}

		detalleMovimiento := movimiento["MovimientoDetalle"].(map[string]interface{})
		t, err := time.Parse(layoutDate, detalleMovimiento["FechaRegistro"].(string))
		if err != nil {
			panic(err)
		}
		movimientoDetalle := models.MovimientoDetalle{
			Valor:         detalleMovimiento["Valor"].(float64),
			FechaCreacion: t,
			Descripcion:   detalleMovimiento["Descripcion"].(string),
		}

		movimientoProcesoExternoManager.RegistrarMovimientoProcesoExterno(&movimientoProcesoExterno, &movimientoDetalle)
		dataResponse["status"] = "registrado"
		responseformat.SetResponseFormat(&c.Controller, dataResponse, "", 200)
	} else {
		panic(err)
	}
}

// Post ...
// @Title Post
// @Description create MovimientoProcesoExterno
// @Param	body		body 	models.MovimientoProcesoExterno	true		"body for MovimientoProcesoExterno content"
// @Success 201 {int} models.MovimientoProcesoExterno
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoProcesoExternoController) Post() {
	var v models.MovimientoProcesoExterno
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddMovimientoProcesoExterno(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err
	}
}

// GetOne ...
// @Title Get One
// @Description get MovimientoProcesoExterno by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MovimientoProcesoExterno
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MovimientoProcesoExternoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMovimientoProcesoExternoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
}

// GetAll ...
// @Title Get All
// @Description get MovimientoProcesoExterno
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MovimientoProcesoExterno
// @Failure 403
// @router / [get]
func (c *MovimientoProcesoExternoController) GetAll() {
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

	// logs.Debug(query)

	l, err := models.GetAllMovimientoProcesoExterno(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
}

// movimientoFiltroJsonB ...
// @Title movimientoFiltroJsonB
// @Description get movimientoFiltroJsonB permite obtener el movimiento proceso externo sin hacerle transformaciones al query del jsonb
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MovimientoProcesoExterno
// @Failure 403
// @router /movimientoFiltroJsonB [get]
func (c *MovimientoProcesoExternoController) MovimientoFiltroJsonB() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string
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

	query = map[string]string{
		"Detalle__json_contains": c.GetString("query"),
	}

	// logs.Debug(query)

	l, err := models.GetAllMovimientoProcesoExterno(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
}

// Put ...
// @Title Put
// @Description update the MovimientoProcesoExterno
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.MovimientoProcesoExterno	true		"body for MovimientoProcesoExterno content"
// @Success 200 {object} models.MovimientoProcesoExterno
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MovimientoProcesoExternoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.MovimientoProcesoExterno{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateMovimientoProcesoExternoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err
		}
	} else {
		c.Data["json"] = err.Error()
	}
}

// Delete ...
// @Title Delete
// @Description delete the MovimientoProcesoExterno
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MovimientoProcesoExternoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMovimientoProcesoExterno(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
}
