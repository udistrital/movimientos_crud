package helpers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/imdario/mergo"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

func CrearMovimientoDetalle(cuentaMovimientoDetalle models.CuentasMovimientoProcesoExterno, publicar bool, nuevoMovimiento string) (movimientoDetalleRegistrado *models.MovimientoDetalle, movimientoCambiado string, outputError map[string]interface{}) {
	o := orm.NewOrm()

	if err := o.Begin(); err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	var idMovProcExterno string
	var idNuevoMovProcExterno string = nuevoMovimiento
	if idNuevoMovProcExterno == "" {
		idMovProcExterno = cuentaMovimientoDetalle.Mov_Proc_Ext
	} else {
		idMovProcExterno = idNuevoMovProcExterno
	}
	var estado map[string]interface{}

	if idMovProcExterno == "" {
		err := "No se ha recibido un ID de Movimiento Proceso Externo"
		panic(errorctrl.Error("crearMovimientoDetalle - idMovProcExterno == \"\"", err, "400"))
	} else {
		if idCast, err := strconv.Atoi(idMovProcExterno); err != nil {
			panic(err)
		} else {
			if result, err := models.GetMovimientoProcesoExternoById(idCast); err != nil {
				logs.Error(err)
				panic(err)
			} else {
				if err := json.Unmarshal([]byte(result.Detalle), &estado); err != nil {
					logs.Error(err)
					panic(err)
				}
				if estado["Estado"].(string) == "Publicado" && idNuevoMovProcExterno == "" {
					err := "No se pueden crear movimientos detalle sobre un movimiento proceso externo publicado, se va a crear un nuevo movimiento proceso externo"
					logs.Warn(err)
					// logs.Debug("RESULT: ", result.Activo)
					var detalleNuevoMov map[string]interface{}

					if err := json.Unmarshal([]byte(result.Detalle), &detalleNuevoMov); err != nil {
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
						TipoMovimientoId:         result.TipoMovimientoId,
						ProcesoExterno:           result.ProcesoExterno,
						MovimientoProcesoExterno: result.MovimientoProcesoExterno,
						Activo:                   result.Activo,
						Detalle:                  string(detalleNuevoMovStr),
					}

					// logs.Debug("NUEVO MOVIMIENTO: ", &nuevoMovimiento)
					if _, err := models.AddMovimientoProcesoExterno(&nuevoMovimiento); err != nil {
						panic(errorctrl.Error("CrearMovimientoDetalle - models.AddMovimientoProcesoExterno(result)", err, "500"))
					}
				} else if estado["Estado"].(string) != "Preliminar" {
					err := "No se reconoce el estado del movimiento proceso externo"
					panic(errorctrl.Error("crearMovimientoDetalle - estado[\"Estado\"].(string) != \"Preliminar\"", err, "500"))
				}
			}
		}
	}

	saldo := cuentaMovimientoDetalle.Saldo
	valor := cuentaMovimientoDetalle.Valor

	if !publicar {
		if saldo == 0 && valor == 0 {
			err := "Tanto el saldo como el valor tienen un valor de 0, no se puede añadir el movimiento detalle"
			panic(errorctrl.Error("crearMovimientoDetalle - saldo == 0 && valor == 0", err, "400"))
		} else if saldo != 0 && valor != 0 {
			err := "Tanto el saldo como el valor tienen un valor diferente de 0, no se puede añadir el movimiento detalle"
			panic(errorctrl.Error("crearMovimientoDetalle - saldo != 0 && valor != 0", err, "400"))
		}
	} else {
		if saldo != valor {
			err := "El saldo y el valor de publicación debe ser el mismo"
			panic(errorctrl.Error("crearMovimientoDetalle - saldo != valor", err, "400"))
		}
	}

	detalleCuenPre := cuentaMovimientoDetalle.Cuen_Pre

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

func RegistroMovimientoDetalle(detalleCuenPre string, idMovProcExterno string, saldo float64, valor float64, publicar bool) (registroMovimientoDetalleRespuesta models.MovimientoDetalle, outputError map[string]interface{}) {
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

	if nuevoDeltaAcum, nuevoSaldo, nuevoValor, err2 = CalcularMontos(detalleCuenPre, saldo, valor, publicar); err2 != nil {
		logs.Error(err2)
		outputError := errorctrl.Error("RegistroMovimientoDetalle - CalcularDeltaAcum(detalleCuenPre)", err2, "500")
		return models.MovimientoDetalle{}, outputError
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

func CalcularMontos(detalleCuenPre string, saldo float64, valor float64, publicar bool) (detalAcumRespuesta float64, saldoRespuesta float64, valorRespuesta float64, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("CalcularMontos - Unhandled Error!", "500")

	cuentaSolicitada := models.CuentasMovimientoProcesoExterno{
		Cuen_Pre: detalleCuenPre,
	}

	var detalleUltimoMovimientoDetalle map[string]interface{}
	var result models.MovimientoDetalle
	var err map[string]interface{}

	if result, err = GetUltimo(cuentaSolicitada); err != nil {
		logs.Warn(err)
		if saldo != 0 {
			return saldo, saldo, saldo, nil
		} else if valor != 0 {
			return valor, valor, valor, nil
		} else {
			outputError := errorctrl.Error("RegistroMovimientoDetalle - CalcularMontos(detalleCuenPre)", err, "500")
			return 0, 0, 0, outputError
		}
	} else if err == nil && publicar {
		if saldo != 0 {
			return 0, saldo, saldo, nil
		} else if valor != 0 {
			return 0, valor, valor, nil
		} else {
			outputError := errorctrl.Error("RegistroMovimientoDetalle - CalcularMontos(detalleCuenPre)", err, "500")
			return 0, 0, 0, outputError
		}
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

	if saldo != 0 {
		saldoRespuesta = saldo
		valorRespuesta = saldo - detalleUltimoMovimientoDetalle["DeltaAcum"].(float64)
		detalAcumRespuesta = valorRespuesta + result.Saldo
	} else if valor != 0 {
		valorRespuesta = valor
		saldoRespuesta = detalleUltimoMovimientoDetalle["DeltaAcum"].(float64) + valor
		detalAcumRespuesta = valor + result.Saldo
	}

	return detalAcumRespuesta, saldoRespuesta, valorRespuesta, nil

}

func CrearMovimientosDetalle(cuentasMovimientoDetalle []models.CuentasMovimientoProcesoExterno, publicar bool) (cuentasMovimientoDetalleRespuesta []models.MovimientoDetalle, outputError map[string]interface{}) {
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
