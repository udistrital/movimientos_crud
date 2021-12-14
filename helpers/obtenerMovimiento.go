package helpers

import (
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/movimientos_crud/helpers/utils"
	"github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/utils_oas/errorctrl"
)

func ObtenerMovimientos(estado string, atributo string, origen string) (movimientosRespuesta []interface{}, outputError map[string]interface{}) {
	defer errorctrl.ErrorControlFunction("ObtenerMovimientos - Unhandled Error!", "500")

	var movimientoObtenido []interface{}
	var err error

	var filtroJsonB map[string]interface{} = map[string]interface{}{
		"Estado": estado,
	}

	if origen != "" {
		filtroJsonB[origen] = atributo
	}

	filtroSerialized, _ := utils.Serializar(filtroJsonB)

	var query map[string]string = map[string]string{
		"Detalle__json_contains": filtroSerialized,
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

	if movimientoObtenido, err = models.GetAllMovimientoProcesoExterno(query, fields, sortby, order, offset, limit); err != nil {
		logs.Error(err)
		outputError := errorctrl.Error("ObtenerMovimientos - models.GetAllMovimientoProcesoExterno(query, fields, sortby, order, offset, limit)", err, "500")
		return []interface{}{}, outputError
	}

	return movimientoObtenido, nil

}
