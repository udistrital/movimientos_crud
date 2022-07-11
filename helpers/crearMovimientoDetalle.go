package helpers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

// CrearMovimientoDetalle se encarga de insertar un movimiento detalle en la base de datos
func CrearMovimientoDetalle(
	cuentaMovimientoDetalle models.CuentasMovimientoProcesoExterno,
	publicar bool,
	nuevoMovimiento string,
) (
	movimientoDetalleRegistrado *models.MovimientoDetalle,
	movimientoCambiado string,
	outputError map[string]interface{},
) {
	// Inicio conexion DB
	o := orm.NewOrm()

	if err := o.Begin(); err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
			panic(r)
		}
		o.Commit()
	}()
	// Variables
	var idMovProcExterno string
	var idNuevoMovProcExterno string = nuevoMovimiento

	var saldo float64 = cuentaMovimientoDetalle.Saldo
	var valor float64 = cuentaMovimientoDetalle.Valor
	// Si hay varias cuentas para registrar
	if idNuevoMovProcExterno == "" {
		idMovProcExterno = cuentaMovimientoDetalle.Mov_Proc_Ext
	} else {
		idMovProcExterno = idNuevoMovProcExterno
	}

	if idMovProcExterno == "" {
		err := "No se ha recibido un ID de Movimiento Proceso Externo"
		panic(errorctrl.Error("crearMovimientoDetalle - idMovProcExterno == \"\"", err, "400"))
	}

	if !publicar {
		if saldo == 0 && valor == 0 {
			err := "Tanto el saldo como el valor tienen un valor de 0, no se puede añadir el movimiento detalle"
			panic(errorctrl.Error("crearMovimientoDetalle - saldo == 0 && valor == 0", err, "400"))
		} else if saldo != 0 && valor != 0 {
			err := "Tanto el saldo como el valor tienen un valor diferente de 0, no se puede añadir el movimiento detalle"
			panic(errorctrl.Error("crearMovimientoDetalle - saldo != 0 && valor != 0", err, "400"))
		}
	}

	detalleCuenPre := cuentaMovimientoDetalle.Cuen_Pre

	// logs.Debug("DETALLE CUENTA PRESUPUESTAL: ", detalleCuenPre)

	if detalleCuenPre == "" {
		err := "No se han ingresado datos de cuentas para crear movimientos detalle"
		panic(errorctrl.Error("crearMovimientoDetalle - detalleCuenPre == \"\"", err, "400"))
	}

	// logs.Debug("INSERTAR movimiento: CrearMovimientoDetalle", idMovProcExterno)

	if registroMovimientoDetalle, err := RegistroMovimientoDetalle(detalleCuenPre, idMovProcExterno, saldo, valor, publicar); err != nil {
		logs.Error(err)
		panic(err)
	} else {
		if result, err := models.AddMovimientoDetalle(&registroMovimientoDetalle); err != nil {
			logs.Error(err)
			panic(err)
		} else {
			resultCast := int(result)
			if movimientoDetalleRegistrado, err = models.GetMovimientoDetalleById(resultCast); err != nil {
				logs.Error(err)
				panic(err)
			}
		}
	}

	return movimientoDetalleRegistrado, idNuevoMovProcExterno, nil
}

// RegistroMovimientoDetalle obtiene la estructura del movimiento detalle a ser insertado
func RegistroMovimientoDetalle(
	detalleCuenPre string,
	idMovProcExterno string,
	saldo float64,
	valor float64,
	publicar bool,
) (
	registroMovimientoDetalleRespuesta models.MovimientoDetalle,
	outputError map[string]interface{},
) {
	defer errorctrl.ErrorControlFunction("RegistroMovimientoDetalle - Unhandled Error!", "500")

	var idMovProcExternoCast int
	var registroMovProcExterno models.MovimientoProcesoExterno
	var err error

	if idMovProcExternoCast, err = strconv.Atoi(idMovProcExterno); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("RegistroMovimientoDetalle - strconv.Atoi(idMovProcExterno)", err, "400")
		return models.MovimientoDetalle{}, outputError
	}

	// logs.Debug("INSERTAR movimiento: ", idMovProcExternoCast)

	registroMovProcExterno = models.MovimientoProcesoExterno{
		Id: idMovProcExternoCast,
	}

	var nuevoSaldo float64
	var nuevoValor float64
	var err2 map[string]interface{}
	var nuevoDetalleCuenPre map[string]interface{}

	if nuevoSaldo, nuevoValor, err2 = CalcularMontos(detalleCuenPre, idMovProcExterno, saldo, valor, publicar); err2 != nil {
		logs.Error(err2)
		return models.MovimientoDetalle{}, err2
	}

	if err := json.Unmarshal([]byte(detalleCuenPre), &nuevoDetalleCuenPre); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("RegistroMovimientoDetalle - json.Unmarshal([]byte(detalleCuenPre), &nuevoDetalleCuenPre)", err, "400")
		return models.MovimientoDetalle{}, outputError
	}

	var nuevoDetalleCuenPreCast []byte

	if nuevoDetalleCuenPreCast, err = json.Marshal(nuevoDetalleCuenPre); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("RegistroMovimientoDetalle - json.Marshal(nuevoDetalleCuenPre)", err, "400")
		return models.MovimientoDetalle{}, outputError
	}

	registroMovimientoDetalleRespuesta = models.MovimientoDetalle{
		Activo:                     true,
		Descripcion:                "Creación de movimiento detalle",
		Detalle:                    string(nuevoDetalleCuenPreCast),
		MovimientoProcesoExternoId: &registroMovProcExterno,
		Saldo:                      nuevoSaldo,
		Valor:                      nuevoValor,
	}

	return registroMovimientoDetalleRespuesta, nil
}

// CalcularMontos devuelve los montos a insetar en valor, saldo y delta acumulado del movimiento respectivo
func CalcularMontos(
	detalleCuenPre string,
	idMovProcExterno string,
	saldo float64,
	valor float64,
	publicar bool,
) (
	saldoRespuesta float64,
	valorRespuesta float64,
	outputError map[string]interface{},
) {
	defer errorctrl.ErrorControlFunction("CalcularMontos - Unhandled Error!", "500")

	var cuentaSolicitada models.CuentasMovimientoProcesoExterno
	var err error
	var result models.MovimientoDetalle
	var formatError map[string]interface{}

	logs.Debug("ID MOVIMIENTOS OBTENIDOS: ", idMovProcExterno)

	var infoFiltro map[string]interface{}
	json.Unmarshal([]byte(detalleCuenPre), &infoFiltro)
	var stringFiltro = make(map[string]interface{})
	for k, prop := range infoFiltro {
		if k == "RubroId" || k == "FuenteFinanciamientoId" || k == "ActividadId" || k == "PlanAdquisicionesId" {
			stringFiltro[k] = prop
		}
	}

	var detalleTemp []byte

	if detalleTemp, err = json.Marshal(stringFiltro); err != nil {
		logs.Error(err)
	}

	//logs.Debug("CONSULTAR CUENTA DETALLE: ", string(detalleTemp))

	if idMovProcExterno != "" {
		cuentaSolicitada = models.CuentasMovimientoProcesoExterno{
			Cuen_Pre:     string(detalleTemp),
			Mov_Proc_Ext: idMovProcExterno,
		}

		//logs.Debug("CUENTA: ", cuentaSolicitada)

		result, formatError = GetUltimo(cuentaSolicitada)
		if formatError != nil {
			// logs.Debug("Entré al error")
			logs.Warn(formatError)
		}
	}

	if valor != 0 {
		valorRespuesta = valor
		saldoRespuesta = result.Saldo + valorRespuesta
	} else if saldo != 0 {
		saldoRespuesta = saldo
		valorRespuesta = saldoRespuesta - result.Saldo
	}
	return saldoRespuesta, valorRespuesta, nil

}

// CrearMovimientosDetalle crea todos los movimientos detalle de un arreglo recibido
func CrearMovimientosDetalle(
	cuentasMovimientoDetalle []models.CuentasMovimientoProcesoExterno,
	publicar bool,
) (
	cuentasMovimientoDetalleRespuesta []models.MovimientoDetalle,
	outputError map[string]interface{},
) {
	const funcion = "CrearMovimientosDetalle - "
	defer errorctrl.ErrorControlFunction(funcion+"Unhandled Error!", "500")

	cuentasMovimientoDetalleRespuesta = make([]models.MovimientoDetalle, len(cuentasMovimientoDetalle))
	var movimientoCambiado string = ""
	for k, cuenta := range cuentasMovimientoDetalle {
		// logs.Debug("k", k)
		var resultado models.MovimientoDetalle
		var err map[string]interface{}
		var v *models.MovimientoDetalle
		if v, movimientoCambiado, err = CrearMovimientoDetalle(cuenta, publicar, movimientoCambiado); err == nil || err["status"].(string) == "404" {
			// logs.Debug(fmt.Sprintf("resultadoErr: %+v", resultado))
			resultado = *v
			logs.Warn(err)
		} else {
			return nil, err
		}
		cuentasMovimientoDetalleRespuesta[k] = resultado
	}

	// logs.Debug(fmt.Sprintf("cuentasMovimientoDetalleRespuesta: %+v", cuentasMovimientoDetalleRespuesta))

	return cuentasMovimientoDetalleRespuesta, nil

}
