package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/movimientos_crud/helpers"
	movimientoDetalleManager "github.com/udistrital/movimientos_crud/managers/movimientoDetalleManager"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
	"github.com/udistrital/utils_oas/responseformat"
)

// MovimientoDetalleController operations for MovimientoDetalle
type MovimientoDetalleController struct {
	beego.Controller
}

// URLMapping ...
func (c *MovimientoDetalleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("RegistrarMultiple", c.RegistrarMultiple)
	c.Mapping("PostUltimoMovDetalle", c.PostUltimoMovDetalle)
}

// RegistrarMultiple ...
// @Title RegistrarMultiple
// @Description Registra multiples movimientos proceso externo y movimientos detalle
// @Param	body		body 	[]models.MovimientoDetalle	true		"body for MovimientoDetalle content"
// @Success 201 {int} responseformat
// @Failure 403 body is empty
// @router /registrar_multiple [post]
func (c *MovimientoDetalleController) RegistrarMultiple() {
	var v []*models.MovimientoDetalle

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		log.Panicln(err.Error())
	}

	ids := movimientoDetalleManager.RegistrarMultipleManager(v)

	response := make(map[string]interface{})
	response["Ids"] = ids
	responseformat.SetResponseFormat(&c.Controller, response, "", 201)
}

// DeleteMultiple ...
// @Title DeleteMultiple
// @Description delete the MovimientoDetalle with transaction
// @Param	body		body 	[]int	true		"Array of (int) IDs that you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Body is empty
// @router /eliminar_multiple [post]
func (c *MovimientoDetalleController) DeleteMultiple() {

	var (
		err                  error
		movimientoDetalleIDS []int
	)

	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &movimientoDetalleIDS); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}

	if err = movimientoDetalleManager.EliminarMultipleManager(movimientoDetalleIDS); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}

	c.Data["json"] = "OK"

}

// Post ...
// @Title Post
// @Description create MovimientoDetalle
// @Param	body		body 	models.MovimientoDetalle	true		"body for MovimientoDetalle content"
// @Success 201 {int} models.MovimientoDetalle
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoDetalleController) Post() {
	var v models.MovimientoDetalle
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddMovimientoDetalle(&v); err == nil {
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
// @Description get MovimientoDetalle by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MovimientoDetalle
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MovimientoDetalleController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetMovimientoDetalleById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
}

// GetAll ...
// @Title Get All
// @Description get MovimientoDetalle
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MovimientoDetalle
// @Failure 403
// @router / [get]
func (c *MovimientoDetalleController) GetAll() {
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

	l, err := models.GetAllMovimientoDetalle(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
}

// Put ...
// @Title Put
// @Description update the MovimientoDetalle
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.MovimientoDetalle	true		"body for MovimientoDetalle content"
// @Success 200 {object} models.MovimientoDetalle
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MovimientoDetalleController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.MovimientoDetalle{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateMovimientoDetalleById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err
	}
}

// Delete ...
// @Title Delete
// @Description delete the MovimientoDetalle
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MovimientoDetalleController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMovimientoDetalle(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
}

// PostUltimoMovDetalle ...
// @Title PostUltimoMovDetalle
// @Description post UltimoMovDetalle se encarga de devolver el último movimiento detalle asociado a una denominada cuenta presupuestal
// @Param     body      body   []models.CuentasMovimientoProcesoExterno  true   "Valor de la cuenta presupuestal o las cuentas presupuestales de las que quiere recuperar el último movimiento"
// @Success   200   {object}   []models.MovimientoDetalle
// @Failure   403   body is empty
// @router /postUltimoMovDetalle [post]
func (c *MovimientoDetalleController) PostUltimoMovDetalle() {
	defer errorctrl.ErrorControlController(c.Controller, "MovimientoDetalleController")
	var arrayCuentas []models.CuentasMovimientoProcesoExterno

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &arrayCuentas); err != nil {
		panic(err)
	}

	if result, err := helpers.GetAllUltimos(arrayCuentas); err != nil {
		// logs.Debug("error")
		panic(err)
	} else {
		// logs.Debug("Información: ", arrayCuentas, result)
		c.Data["json"] = result
		c.Data["status"] = 200
	}
	c.ServeJSON()

}

// CrearMovimientosDetalle ...
// @Title CrearMovimientosDetalle
// @Description post CrearMovimientosDetalle se encarga de devolver crear los movimientos detalle correspondientes a las cuentas recibidas
// @Param     body      body   []models.CuentasMovimientoProcesoExterno  true   "Cuentas presupuestales con su respectivo movimiento proceso externo y el valor/saldo afectado"
// @Success   201   {object}   []models.MovimientoDetalle
// @Failure   403   body is empty
// @router /crearMovimientosDetalle [post]
func (c *MovimientoDetalleController) CrearMovimientosDetalle() {
	defer errorctrl.ErrorControlController(c.Controller, "CrearMovimientosDetalle")

	var arrayCuentas []models.CuentasMovimientoProcesoExterno

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &arrayCuentas); err != nil {
		panic(err)
	}

	if result, err := helpers.CrearMovimientosDetalle(arrayCuentas); err != nil {
		// logs.Debug("error")
		panic(err)
	} else {
		// logs.Debug("Información: ", arrayCuentas, result)
		c.Data["json"] = result
		c.Data["status"] = 201
	}
	c.ServeJSON()
}
