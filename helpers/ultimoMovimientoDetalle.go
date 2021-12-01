package helpers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
	"github.com/udistrital/utils_oas/formatdata"
)

func GetUltimo(cuentaMovimientoDetalle models.CuentasMovimientoProcesoExterno) (cuentaMovimientoDetalleRespuesta models.MovimientoDetalle, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("GetUltimo - Unhandled Error!", "500")

	var datosCuenta models.Cuen_Pre

	if err := json.Unmarshal([]byte(cuentaMovimientoDetalle.Cuen_Pre), &datosCuenta); err != nil {
		outputError = errorctrl.Error("GetUltimo", err, "400")
		return models.MovimientoDetalle{}, outputError
	}

	var filtroJsonB map[string]interface{}
	if datosCuenta.ActividadId != "" && datosCuenta.RubroId != "" && datosCuenta.FuenteFinanciamientoId != "" {
		var actividadInt int
		var fuenteInt int
		var err error
		if actividadInt, err = strconv.Atoi(datosCuenta.ActividadId); err != nil {
			logs.Error(err)
			outputError = errorctrl.Error("GetUltimo - strconv.Atoi(datosCuenta.ActividadId)", err, "400")
			return models.MovimientoDetalle{}, outputError
		}

		if fuenteInt, err = strconv.Atoi(datosCuenta.FuenteFinanciamientoId); err != nil {
			logs.Error(err)
			outputError = errorctrl.Error("GetUltimo - strconv.Atoi(datosCuenta.FuenteFinanciamientoId)", err, "400")
			return models.MovimientoDetalle{}, outputError
		}

		filtroJsonB = map[string]interface{}{
			"RubroId":                datosCuenta.RubroId,
			"FuenteFinanciamientoId": fuenteInt,
			"ActividadId":            actividadInt,
		}
	} else if datosCuenta.ActividadId == "" && datosCuenta.RubroId != "" && datosCuenta.FuenteFinanciamientoId != "" {
		if fuenteInt, err := strconv.Atoi(datosCuenta.FuenteFinanciamientoId); err != nil {
			logs.Error(err)
			outputError = errorctrl.Error("GetUltimo - strconv.Atoi(datosCuenta.FuenteFinanciamientoId)", err, "400")
			return models.MovimientoDetalle{}, outputError
		} else {
			filtroJsonB = map[string]interface{}{
				"RubroId":                datosCuenta.RubroId,
				"FuenteFinanciamientoId": fuenteInt,
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
		outputError = errorctrl.Error("GetUltimo - models.GetAllMovimientoDetalle(query, nil, sortby, order, 0, limit)", err, "404")
		return models.MovimientoDetalle{}, outputError
	} else {
		// logs.Debug(fmt.Sprintf("result: %+v", result))
		if len(result) > 0 {
			var lastMovimientoDetalle models.MovimientoDetalle
			formatdata.FillStruct(result[0], &lastMovimientoDetalle)
			return lastMovimientoDetalle, nil
		} else {
			err := "No se encontró ningún registro que coincida"
			logs.Error(err)
			outputError = errorctrl.Error("GetUltimo - len(result) > 0", err, "404")
			return models.MovimientoDetalle{}, outputError
		}
	}
}

func GetAllUltimos(cuentasMovimientoDetalle []models.CuentasMovimientoProcesoExterno) (cuentasMovimientoDetalleRespuesta []models.MovimientoDetalle, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("GetAllUltimos - Unhandled Error!", "500")

	cuentasMovimientoDetalleRespuesta = make([]models.MovimientoDetalle, len(cuentasMovimientoDetalle))

	for k, cuenta := range cuentasMovimientoDetalle {
		// logs.Debug("k", k)
		var resultado models.MovimientoDetalle
		var err map[string]interface{}
		var v models.MovimientoDetalle
		if v, err = GetUltimo(cuenta); err == nil || err["status"].(string) == "404" {
			// logs.Debug(fmt.Sprintf("resultadoErr: %+v", resultado))
			resultado = v
			logs.Warn(err)
		} else {
			return nil, err
		}
		cuentasMovimientoDetalleRespuesta[k] = resultado
	}

	// logs.Debug(fmt.Sprintf("cuentasMovimientoDetalleRespuesta: %+v", cuentasMovimientoDetalleRespuesta))

	return cuentasMovimientoDetalleRespuesta, nil

}