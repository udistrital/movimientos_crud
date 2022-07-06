package helpers

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_crud/helpers/utils"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

func UltimoMovimientoDetallePublicado(estado string, cuenta string) (cuentaRespuesta []interface{}, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("ultimoMovimientoDetallePublicado - Unhandled Error!", "500")

	var infoFiltro map[string]interface{}
	var err error

	json.Unmarshal([]byte(cuenta), &infoFiltro)
	// logs.Debug("INFOFILTRO: ", infoFiltro)
	var stringFiltro = make(map[string]interface{})
	for k, prop := range infoFiltro {
		if k == "ActividadId" {
			stringFiltro[k] = prop
		}
	}

	// logs.Debug("STRINGFILTRO: ", stringFiltro)

	var detalleTemp []byte

	if detalleTemp, err = json.Marshal(stringFiltro); err != nil {
		logs.Error(err)
	}

	filtroJsonBEstado, _ := utils.Serializar(map[string]interface{}{
		"Estado": estado,
	})

	var query map[string]string = map[string]string{
		"MovimientoProcesoExternoId__Detalle__json_contains": filtroJsonBEstado,
		"Detalle__json_contains":                             string(detalleTemp),
	}

	var sortby []string = []string{
		"FechaCreacion",
	}

	var order []string = []string{
		"desc",
	}

	var limit int64 = int64(1)

	if ultimoMovimiento, err := models.GetAllMovimientoDetalle(query, nil, sortby, order, 0, limit); err != nil {
		logs.Error(err)
		outputError = errorctrl.Error("ultimoMovimientoDetallePublicado - models.GetAllMovimientoDetalle(query, nil, sortby, order, 0, limit)", err, "500")
		return []interface{}{}, outputError
	} else {
		cuentaRespuesta = ultimoMovimiento
	}

	return cuentaRespuesta, nil

}
