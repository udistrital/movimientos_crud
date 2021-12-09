package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

// PublicarMovimientosDetalle realiza una copia de los rubros preliminares y los publica nuevamente en otro movimiento externo con estado publicado
func PublicarMovimientosDetalle(idMovProcExterno int) (movimientosDetalleRespuesta []models.MovimientoDetalle, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("PublicarMovimientosDetalle - Unhandled Error!", "500")
	var detalleMovProcExt map[string]interface{}
	var movimientoProcExtObtenido *models.MovimientoProcesoExterno
	var err error
	var formatErr map[string]interface{}

	if movimientoProcExtObtenido, err = models.GetMovimientoProcesoExternoById(idMovProcExterno); err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("PublicarMovimientosDetalle - models.GetMovimientoProcesoExternoById(idMovProcExterno)", err, "400")
		return []models.MovimientoDetalle{}, outputError
	} else {
		if err := json.Unmarshal([]byte(movimientoProcExtObtenido.Detalle), &detalleMovProcExt); err != nil {
			logs.Error(err)
			outputError = errorctrl.Error("PublicarMovimientosDetalle - json.Unmarshal([]byte(movimientoProcExtObtenido.Detalle), &detalleMovProcExt)", err, "500")
			return []models.MovimientoDetalle{}, outputError
		}

		if detalleMovProcExt["Estado"].(string) == "Publicado" {
			err := "El movimiento ya está en estado publicado, verifique el identificador enviado"
			outputError := errorctrl.Error("crearMovimientoDetalle - estado[\"Estado\"].(string) == \"Publicado\"", err, "500")
			return []models.MovimientoDetalle{}, outputError
		} else if detalleMovProcExt["Estado"].(string) != "Preliminar" {
			err := "No se reconoce el estado del movimiento proceso externo"
			outputError := errorctrl.Error("crearMovimientoDetalle - estado[\"Estado\"].(string) != \"Preliminar\"", err, "500")
			return []models.MovimientoDetalle{}, outputError
		}
	}

	var listaRubros []models.CuentasMovimientoProcesoExterno

	if listaRubros, formatErr = ListaRubros(idMovProcExterno); formatErr != nil {
		logs.Error(formatErr)
		return []models.MovimientoDetalle{}, formatErr
	}

	// logs.Debug("LISTA DE RUBROS: ", listaRubros)

	var ultimosMovimientos []models.MovimientoDetalle

	if ultimosMovimientos, formatErr = GetAllUltimos(listaRubros); formatErr != nil {
		logs.Error(formatErr)
		return []models.MovimientoDetalle{}, formatErr
	}

	var nuevoMovimientoProcesoExterno *models.MovimientoProcesoExterno = movimientoProcExtObtenido

	if err := json.Unmarshal([]byte(nuevoMovimientoProcesoExterno.Detalle), &detalleMovProcExt); err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("PublicarMovimientosDetalle - json.Unmarshal([]byte(nuevoMovimientoProcesoExterno.Detalle), &detalleMovProcExt)", err, "500")
		return []models.MovimientoDetalle{}, outputError
	}

	estadoPublicacion := "Publicado"
	detalleMovProcExt["Estado"] = estadoPublicacion
	nuevoMovimientoProcesoExterno.Id = 0

	if detalleMovProcExtActualizado, err := json.Marshal(detalleMovProcExt); err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("crearMovimientoDetalle - json.Marshal(detalleMovProcExt)", err, "500")
		return []models.MovimientoDetalle{}, outputError
	} else {
		nuevoMovimientoProcesoExterno.Detalle = string(detalleMovProcExtActualizado)
	}

	var idMovimientoProcesoExternoInsertado int

	if movimientoProcesoExternoInsertado, err := models.AddMovimientoProcesoExterno(nuevoMovimientoProcesoExterno); err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("PublicarMovimientosDetalle - models.AddMovimientoProcesoExterno(nuevoMovimientoProcesoExterno)", err, "500")
		return []models.MovimientoDetalle{}, outputError
	} else {
		idMovimientoProcesoExternoInsertado = int(movimientoProcesoExternoInsertado)
	}

	var cuentasPublicar []models.CuentasMovimientoProcesoExterno

	if cuentasPublicar, formatErr = MovimientosDetalleCuentas(ultimosMovimientos, idMovimientoProcesoExternoInsertado); formatErr != nil {
		logs.Error(formatErr)
		return []models.MovimientoDetalle{}, formatErr
	}

	var registroCuentas []models.MovimientoDetalle

	if registroCuentas, formatErr = CrearMovimientosDetalle(cuentasPublicar, true); formatErr != nil {
		logs.Error(formatErr)
		return []models.MovimientoDetalle{}, formatErr
	} else {
		movimientosDetalleRespuesta = registroCuentas
	}

	return movimientosDetalleRespuesta, nil
}

// ListaRubros se encarga de traer los rubros asociados a un movimientos proceso externo, para luego hacer la consulta de sus últimos movimiento relacionados
func ListaRubros(idMovProcExterno int) (detalleCuentasRespuesta []models.CuentasMovimientoProcesoExterno, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("ListaRubros - Unhandled Error!", "500")

	var idMovProcExternoCast string = strconv.Itoa(idMovProcExterno)

	// Se filtra con base en el movimiento proceso externo recibido para traer todos sus movimientos detalle asociados
	query := map[string]string{
		"MovimientoProcesoExternoId__Id": idMovProcExternoCast,
	}

	fields := []string{
		"Detalle",
	}

	// Para traer todos
	limit := int64(-1)

	// Nota: Se envían los parametros sortby y order como nil, además de offset como 0 por defecto
	if result, err := models.GetAllMovimientoDetalle(query, fields, nil, nil, 0, limit); err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("ListaRubros -  models.GetAllMovimientoDetalle(query, fields, nil, nil, 0, limit)", err, "404")
		return []models.CuentasMovimientoProcesoExterno{}, outputError
	} else {
		// logs.Debug(fmt.Sprintf("result: %+v", result))
		if len(result) <= 0 {
			err := "No se encontró ningún registro que coincida"
			logs.Error(err)
			outputError = errorctrl.Error("ListaRubros - len(result) > 0", err, "404")
			return []models.CuentasMovimientoProcesoExterno{}, outputError

		}

		var allDetalleCuentas []models.CuentasMovimientoProcesoExterno
		allDetalleCuentas = make([]models.CuentasMovimientoProcesoExterno, len(result))
		for k, detalle := range result {
			var infoFiltro map[string]interface{}
			json.Unmarshal([]byte(detalle.(map[string]interface{})["Detalle"].(string)), &infoFiltro)
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
			if detalleTemp, err := json.Marshal(stringFiltro); err != nil {
				logs.Error(err)
				outputError = errorctrl.Error("ListaRubros -  json.Marshal(detalle)", err, "500")
				return []models.CuentasMovimientoProcesoExterno{}, outputError
			} else {
				// logs.Debug("DETALLE: ", detalle)
				allDetalleCuentas[k] = models.CuentasMovimientoProcesoExterno{
					Cuen_Pre:     string(detalleTemp),
					Mov_Proc_Ext: idMovProcExternoCast,
				}
			}
		}

		detalleCuentasRespuesta = RemoveDuplicateElement(allDetalleCuentas)

		// logs.Debug("RESPUESTA DE CUENTAS: ", detalleCuentasRespuesta)
	}

	return detalleCuentasRespuesta, nil
}

// RemoveDuplicateElement quita los elementos duplicados de un arreglo de CuentasMovimientoProcesoExterno
func RemoveDuplicateElement(addrs []models.CuentasMovimientoProcesoExterno) (aResp []models.CuentasMovimientoProcesoExterno) {
	result := make([]models.CuentasMovimientoProcesoExterno, 0, len(addrs))
	temp := map[models.CuentasMovimientoProcesoExterno]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// MovimientosDetalleCuentas se encarga de crear los movimientos que van a ser publicados
func MovimientosDetalleCuentas(
	movimientosDetalle []models.MovimientoDetalle,
	idMovProcExterno int,
) (
	cuentasMovimientoDetalleRespuesta []models.CuentasMovimientoProcesoExterno,
	outputError map[string]interface{},
) {
	defer errorctrl.ErrorControlFunction("MovimientosDetalleCuentas - Unhandled Error!", "500")

	idMovProcExternoCast := strconv.Itoa(idMovProcExterno)

	for _, movimiento := range movimientosDetalle {
		var infoDetalle map[string]interface{}
		json.Unmarshal([]byte(movimiento.Detalle), &infoDetalle)

		// logs.Debug("MOVIMIENTO DETALLE: ", movimiento.Detalle)

		switch infoDetalle["DeltaAcum"].(type) {
		case float64:
			cuentaRespuestaTemp := models.CuentasMovimientoProcesoExterno{
				Cuen_Pre:     movimiento.Detalle,
				Mov_Proc_Ext: idMovProcExternoCast,
				// Saldo:        infoDetalle["DeltaAcum"].(float64),
				Valor: infoDetalle["DeltaAcum"].(float64),
			}

			cuentasMovimientoDetalleRespuesta = append(cuentasMovimientoDetalleRespuesta, cuentaRespuestaTemp)
		default:
			// logs.Debug(reflect.TypeOf(infoDetalle["DeltaAcum"]))
			logs.Warn("No se contiene el campo DeltaAcum o no está en el formato adecuado")
		}
	}

	return cuentasMovimientoDetalleRespuesta, nil
}
