package helpers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
)

func GetUltimo(cuentaMovimientoDetalle models.CuentasMovimientoProcesoExterno) (cuentaMovimientoDetalleRespuesta models.MovimientoDetalle, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "GetUltimo - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	var datosCuenta models.Cuen_Pre

	json.Unmarshal([]byte(cuentaMovimientoDetalle.Cuen_Pre), &datosCuenta)

	var filtroJsonB map[string]interface{}

	if datosCuenta.ActividadId != "" {
		if datosCuenta.RubroId != "" {
			if datosCuenta.FuenteFinanciamientoId != "" {
				if actividadInt, err := strconv.Atoi(datosCuenta.ActividadId); err != nil {
					logs.Error(err)
					outputError = map[string]interface{}{
						"funcion": "GetUltimo - strconv.Atoi(datosCuenta.ActividadId)",
						"err":     err,
						"status":  "502",
					}
					return models.MovimientoDetalle{}, outputError
				} else {
					filtroJsonB = map[string]interface{}{
						"RubroId": datosCuenta.RubroId,
						// "FuenteFinanciamientoId": datosCuenta.FuenteFinanciamientoId,
						"ActividadId": actividadInt,
					}
				}
			}
		}
	} else if datosCuenta.ActividadId == "" {
		if datosCuenta.RubroId != "" {
			if datosCuenta.FuenteFinanciamientoId != "" {
				if fuenteInt, err := strconv.Atoi(datosCuenta.FuenteFinanciamientoId); err != nil {
					logs.Error(err)
					outputError = map[string]interface{}{
						"funcion": "GetUltimo - strconv.Atoi(datosCuenta.FuenteFinanciamientoId)",
						"err":     err,
						"status":  "502",
					}
					return models.MovimientoDetalle{}, outputError
				} else {
					filtroJsonB = map[string]interface{}{
						"RubroId":                datosCuenta.RubroId,
						"FuenteFinanciamientoId": fuenteInt,
					}
				}
			}
		}
	}

	data, _ := json.Marshal(filtroJsonB)

	datosMovProcExterno := cuentaMovimientoDetalle.Mov_Proc_Ext

	var query map[string]string

	query = map[string]string{
		"Detalle__json_contains":         string(data),
		"MovimientoProcesoExternoId__Id": datosMovProcExterno,
	}

	// Se sugiere ordenar por fecha de modificación
	sortby := []string{
		"FechaModificacion",
	}

	// El orden descendente velará por traer el último registro modificado
	order := []string{
		"desc",
	}

	// Para traer el último
	limit := int64(1)

	// Nota: Se envían los parámetros de fields como nil y offset por default como 0
	if result, err := models.GetAllMovimientoDetalle(query, nil, sortby, order, 0, limit); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "GetUltimo - models.GetAllMovimientoDetalle(query, nil, sortby, order, 0, limit)",
			"err":     err,
			"status":  "502",
		}
		return models.MovimientoDetalle{}, outputError
	} else {
		logs.Debug(result)
		var lastMovimientoDetalle models.MovimientoDetalle
		formatdata.FillStruct(result[0], &lastMovimientoDetalle)
		return lastMovimientoDetalle, nil
	}
}

func GetAllUltimos(cuentasMovimientoDetalle []models.CuentasMovimientoProcesoExterno) (cuentasMovimientoDetalleRespuesta []models.MovimientoDetalle, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "GetAllUltimos - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	for i := range cuentasMovimientoDetalle {
		if resultado, err := GetUltimo(cuentasMovimientoDetalle[i]); err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{
				"funcion": "GetAllUltimos - GetUltimo(cuentasMovimientoDetalle[i])",
				"err":     err,
				"status":  "502",
			}
			return []models.MovimientoDetalle{}, outputError
		} else {
			cuentasMovimientoDetalleRespuesta = append(cuentasMovimientoDetalleRespuesta, resultado)
		}
	}

	return cuentasMovimientoDetalleRespuesta, nil

}
