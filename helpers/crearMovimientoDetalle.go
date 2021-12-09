package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/imdario/mergo"
	"github.com/udistrital/movimientos_crud/helpers/utils"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
	"github.com/udistrital/utils_oas/formatdata"
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

	var idMovProcExterno string
	var idNuevoMovProcExterno string = nuevoMovimiento
	var idAntiguoMovProcExterno string

	var err error
	var idCast int
	var movimientoObtenidobyId *models.MovimientoProcesoExterno

	var saldo float64 = cuentaMovimientoDetalle.Saldo
	var valor float64 = cuentaMovimientoDetalle.Valor

	if idNuevoMovProcExterno == "" {
		idMovProcExterno = cuentaMovimientoDetalle.Mov_Proc_Ext
	} else {
		idMovProcExterno = idNuevoMovProcExterno
	}
	var estado map[string]interface{}

	if idMovProcExterno == "" {
		err := "No se ha recibido un ID de Movimiento Proceso Externo"
		panic(errorctrl.Error("crearMovimientoDetalle - idMovProcExterno == \"\"", err, "400"))
	}

	if idCast, err = strconv.Atoi(idMovProcExterno); err != nil {
		panic(errorctrl.Error("CrearMovimientoDetalle - strconv.Atoi(idMovProcExterno)", err, "400"))
	}

	if movimientoObtenidobyId, err = models.GetMovimientoProcesoExternoById(idCast); err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(movimientoObtenidobyId.Detalle), &estado); err != nil {
		panic(errorctrl.Error("CrearMovimientoDetalle - json.Unmarshal([]byte(result.Detalle), &estado)", err, "404"))
	}

	// logs.Debug("ESTADO", estado["Estado"].(string))

	var preliminarObtenido []interface{}

	filtroJsonB, _ := utils.Serializar(map[string]interface{}{
		"Estado":              "Preliminar",
		"PlanAdquisicionesId": estado["PlanAdquisicionesId"].(string),
	})

	var query map[string]string = map[string]string{
		"Detalle__json_contains": filtroJsonB,
	}

	var sortby []string = []string{
		"FechaCreacion",
	}

	var order []string = []string{
		"desc",
	}

	var fields []string = nil

	var offset int64 = int64(0)

	var limit int64 = int64(2)

	if preliminarObtenido, err = models.GetAllMovimientoProcesoExterno(query, fields, sortby, order, offset, limit); err != nil {
		panic(errorctrl.Error("CrearMovimientoDetalle - json.Unmarshal([]byte(result.Detalle), &detalleNuevoMov)", err, "500"))
	}

	if len(preliminarObtenido) <= 0 {
		err := "No se han encontrado preliminares para el plan de adquisiciones"
		logs.Error(err)
		panic(errorctrl.Error("CrearMovimientoDetalle - json.Unmarshal([]byte(result.Detalle), &detalleNuevoMov)", err, "500"))
	} else if len(preliminarObtenido) > 1 {
		var preliminarObtenidoStructed map[string]interface{}

		formatdata.FillStruct(preliminarObtenido[1], &preliminarObtenidoStructed)

		idAntiguoMovProcExterno = fmt.Sprintf("%.0f", preliminarObtenidoStructed["Id"].(float64))
	} else {
		var preliminarObtenidoStructed map[string]interface{}

		formatdata.FillStruct(preliminarObtenido[0], &preliminarObtenidoStructed)

		idAntiguoMovProcExterno = fmt.Sprintf("%.0f", preliminarObtenidoStructed["Id"].(float64))
	}

	if estado["Estado"].(string) == "Publicado" && idNuevoMovProcExterno == "" && !publicar {
		errString := "No se pueden crear movimientos detalle sobre un movimiento proceso externo publicado, se va a crear un nuevo movimiento proceso externo"
		logs.Warn(errString)

		var detalleNuevoMov map[string]interface{}

		if err := json.Unmarshal([]byte(movimientoObtenidobyId.Detalle), &detalleNuevoMov); err != nil {
			logs.Error(err)
			panic(errorctrl.Error("CrearMovimientoDetalle - json.Unmarshal([]byte(result.Detalle), &detalleNuevoMov)", err, "500"))
		}

		detalleNuevoMov["Estado"] = "Preliminar"

		var detalleNuevoMovStr []byte
		var err2 error

		if detalleNuevoMovStr, err2 = json.Marshal(detalleNuevoMov); err2 != nil {
			logs.Error(err)
			panic(errorctrl.Error("CrearMovimientoDetalle - json.Marshal(detalleNuevoMov)", err2, "500"))
		}

		nuevoMovimiento := models.MovimientoProcesoExterno{
			TipoMovimientoId:         movimientoObtenidobyId.TipoMovimientoId,
			ProcesoExterno:           movimientoObtenidobyId.ProcesoExterno,
			MovimientoProcesoExterno: movimientoObtenidobyId.MovimientoProcesoExterno,
			Activo:                   movimientoObtenidobyId.Activo,
			Detalle:                  string(detalleNuevoMovStr),
		}

		// logs.Debug("NUEVO MOVIMIENTO: ", &nuevoMovimiento)
		if adicionMov, err := models.AddMovimientoProcesoExterno(&nuevoMovimiento); err != nil {
			panic(errorctrl.Error("CrearMovimientoDetalle - models.AddMovimientoProcesoExterno(result)", err, "500"))
		} else {
			idNuevoMovProcExterno = strconv.FormatInt(adicionMov, 10)
			idMovProcExterno = idNuevoMovProcExterno
		}

	} else if estado["Estado"].(string) != "Preliminar" && !publicar {
		err := "No se reconoce el estado del movimiento proceso externo"
		panic(errorctrl.Error("crearMovimientoDetalle - estado[\"Estado\"].(string) != \"Preliminar\"", err, "500"))
	}

	// logs.Debug("MOVIMIENTO ANTIGUO: ", idAntiguoMovProcExterno)

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

	logs.Debug("DETALLE CUENTA PRESUPUESTAL: ", detalleCuenPre)

	if detalleCuenPre == "" {
		err := "No se han ingresado datos de cuentas para crear movimientos detalle"
		panic(errorctrl.Error("crearMovimientoDetalle - detalleCuenPre == \"\"", err, "400"))
	}

	// logs.Debug("INSERTAR movimiento: CrearMovimientoDetalle", idMovProcExterno)

	if registroMovimientoDetalle, err := RegistroMovimientoDetalle(detalleCuenPre, idAntiguoMovProcExterno, idMovProcExterno, saldo, valor, publicar, estado["PlanAdquisicionesId"].(string)); err != nil {
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
	idAntiguoMovProcExterno string,
	idMovProcExterno string,
	saldo float64,
	valor float64,
	publicar bool,
	planAdquisicionesId string,
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

	var nuevoDeltaAcum float64
	var nuevoSaldo float64
	var nuevoValor float64
	var err2 map[string]interface{}
	var nuevoDetalleCuenPre map[string]interface{}

	if nuevoDeltaAcum, nuevoSaldo, nuevoValor, err2 = CalcularMontos(detalleCuenPre, idAntiguoMovProcExterno, idMovProcExterno, saldo, valor, publicar, planAdquisicionesId); err2 != nil {
		logs.Error(err2)
		return models.MovimientoDetalle{}, err2
	}

	if err := json.Unmarshal([]byte(detalleCuenPre), &nuevoDetalleCuenPre); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("RegistroMovimientoDetalle - json.Unmarshal([]byte(detalleCuenPre), &nuevoDetalleCuenPre)", err, "400")
		return models.MovimientoDetalle{}, outputError
	}

	deltaEnviado := map[string]interface{}{
		"DeltaAcum": nuevoDeltaAcum,
	}

	if err := mergo.Merge(&nuevoDetalleCuenPre, deltaEnviado); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("RegistroMovimientoDetalle - mergo.Merge(&new_detalle_CuenPre, deltaEnviado)", err, "400")
		return models.MovimientoDetalle{}, outputError
	}

	for k, detalle := range nuevoDetalleCuenPre {
		if k != "RubroId" {
			switch detalle.(type) {
			case string:
				if nuevoDetalleCuenPre[k], err = strconv.Atoi(detalle.(string)); err != nil {
					logs.Error(err)
					outputError := errorctrl.Error("RegistroMovimientoDetalle - strconv.Atoi(detalle.(string))", err, "400")
					return models.MovimientoDetalle{}, outputError
				}
			default:
				logs.Warn("La variable no es de tipo string")
			}

		}
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
	idAntiguoMovProcExterno string,
	idMovProcExterno string,
	saldo float64,
	valor float64,
	publicar bool,
	planAdquisicionesId string,
) (
	detalAcumRespuesta float64,
	saldoRespuesta float64,
	valorRespuesta float64,
	outputError map[string]interface{},
) {
	defer errorctrl.ErrorControlFunction("CalcularMontos - Unhandled Error!", "500")

	var cuentaSolicitada models.CuentasMovimientoProcesoExterno
	var err error
	var detalleUltimoMovimientoDetalle map[string]interface{}
	var result models.MovimientoDetalle
	var formatError map[string]interface{}

	// logs.Debug("ID MOVIMIENTOS OBTENIDOS: ", idMovProcExterno, idAntiguoMovProcExterno)

	var infoFiltro map[string]interface{}
	json.Unmarshal([]byte(detalleCuenPre), &infoFiltro)
	var stringFiltro = make(map[string]interface{})
	for k, prop := range infoFiltro {
		if k == "RubroId" || k == "FuenteFinanciamientoId" || k == "ActividadId" {
			switch prop.(type) {
			case float64:
				propCast := fmt.Sprintf("%.0f", prop.(float64))
				stringFiltro[k] = propCast
			default:
				// logs.Debug(reflect.TypeOf(prop))
				stringFiltro[k] = prop
			}
		}
	}

	var detalleTemp []byte

	detalleTemp, err = json.Marshal(stringFiltro)

	// logs.Debug("CONSULTAR CUENTA DETALLE: ", string(detalleTemp))

	if idMovProcExterno != "" {
		cuentaSolicitada = models.CuentasMovimientoProcesoExterno{
			Cuen_Pre:     string(detalleTemp),
			Mov_Proc_Ext: idMovProcExterno,
		}

		// logs.Debug("PRIMERA CUENTA: ", cuentaSolicitada)

		result, formatError = GetUltimo(cuentaSolicitada)
		if formatError != nil {
			// logs.Debug("Entré al error")
			logs.Warn(formatError)

			if idAntiguoMovProcExterno != "" {
				cuentaSolicitada = models.CuentasMovimientoProcesoExterno{
					Cuen_Pre:     string(detalleTemp),
					Mov_Proc_Ext: idAntiguoMovProcExterno,
				}

				// logs.Debug("SEGUNDA CUENTA: ", cuentaSolicitada)

				result, formatError = GetUltimo(cuentaSolicitada)
				if formatError != nil {
					// logs.Debug("NO PUDE OBTENER")
					logs.Warn(formatError)
					if saldo != 0 {
						return saldo, saldo, saldo, nil
					} else if valor != 0 {
						return valor, valor, valor, nil
					}
					return 0, 0, 0, outputError
				}
			}

		} else {
			idAntiguoMovProcExterno = ""
		}
	}

	// logs.Debug("CUENTA SOLICITADA: ", cuentaSolicitada)

	result, formatError = GetUltimo(cuentaSolicitada)
	if formatError != nil && !publicar {
		// logs.Debug("NO OBTUVE EL ÚLTIMO")
		logs.Warn(formatError)
		if saldo != 0 {
			return saldo, saldo, saldo, nil
		} else if valor != 0 {
			return valor, valor, valor, nil
		}
		return 0, 0, 0, outputError
	}

	// logs.Debug("RESULTADO GET ULTIMO: ", result)

	if err := json.Unmarshal([]byte(result.Detalle), &detalleUltimoMovimientoDetalle); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("CalcularMontos - json.Unmarshal([]byte(result.Detalle), &detalleUltimoMovimientoDetalle)", err, "400")
		return 0, 0, 0, outputError
	}

	switch detalleUltimoMovimientoDetalle["DeltaAcum"].(type) {
	case nil:
		if saldo != 0 {
			return saldo, saldo, saldo, nil
		} else if valor != 0 {
			return valor, valor, valor, nil
		}
	}

	if publicar {
		// logs.Debug("Entré a publicar")

		valorRespuesta = detalleUltimoMovimientoDetalle["DeltaAcum"].(float64)

		var publicadoObtenido []interface{}

		filtroJsonB, _ := utils.Serializar(map[string]interface{}{
			"Estado": "Publicado",
		})

		var query map[string]string = map[string]string{
			"Detalle__json_contains": filtroJsonB,
		}

		var sortby []string = []string{
			"FechaCreacion",
		}

		var order []string = []string{
			"desc",
		}

		var fields []string = nil

		var offset int64 = int64(0)

		var limit int64 = int64(2)

		if publicadoObtenido, err = models.GetAllMovimientoProcesoExterno(query, fields, sortby, order, offset, limit); err != nil {
			logs.Error(err)
			outputError := errorctrl.Error("CalcularMontos - models.GetAllMovimientoProcesoExterno(query, fields, sortby, order, offset, limit)", err, "400")
			return valor, valor, valor, outputError
		}

		if len(publicadoObtenido) <= 0 {
			err := "No se han encontrado movimientos publicados para el plan de adquisiciones"
			logs.Error(err)
			outputError := errorctrl.Error("CalcularMontos - len(publicadoObtenido) <= 0", err, "400")
			return valor, valor, valor, outputError
		}

		var publicadoObtenidoStructed map[string]interface{}

		formatdata.FillStruct(publicadoObtenido[0], &publicadoObtenidoStructed)

		cuentaSolicitada = models.CuentasMovimientoProcesoExterno{
			Cuen_Pre:     detalleCuenPre,
			Mov_Proc_Ext: fmt.Sprintf("%.0f", publicadoObtenidoStructed["Id"].(float64)),
		}

		saldoRespuesta = result.Saldo + valorRespuesta
		detalAcumRespuesta = 0

	}

	if idAntiguoMovProcExterno != "" {
		if saldo != 0 {
			saldoRespuesta = saldo
			valorRespuesta = saldoRespuesta - result.Saldo
			detalAcumRespuesta = valorRespuesta
		} else if valor != 0 {
			valorRespuesta = valor
			saldoRespuesta = result.Saldo + valorRespuesta
			detalAcumRespuesta = valorRespuesta
		}
	} else if idMovProcExterno != "" {
		if saldo != 0 {
			saldoRespuesta = saldo
			valorRespuesta = saldoRespuesta - result.Saldo
			detalAcumRespuesta = valorRespuesta + detalleUltimoMovimientoDetalle["DeltaAcum"].(float64)
		} else if valor != 0 {
			valorRespuesta = valor
			saldoRespuesta = result.Saldo + valorRespuesta
			detalAcumRespuesta = valorRespuesta + detalleUltimoMovimientoDetalle["DeltaAcum"].(float64)
		}
	}

	return detalAcumRespuesta, saldoRespuesta, valorRespuesta, nil

}

// CrearMovimientosDetalle crea todos los movimientos detalle de un arreglo recibido
func CrearMovimientosDetalle(
	cuentasMovimientoDetalle []models.CuentasMovimientoProcesoExterno,
	publicar bool,
) (
	cuentasMovimientoDetalleRespuesta []models.MovimientoDetalle,
	outputError map[string]interface{},
) {
	defer errorctrl.ErrorControlFunction("GetAllUltimos - Unhandled Error!", "500")

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
