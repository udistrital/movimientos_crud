package movimientodetallemanager

import (
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/movimientos_crud/models"
)

// RegistrarMultipleManager realiza multiples registros en una transacción sobre las tablas:
// movimiento_proceso_externo y movimiento_detalle
// realizando primero un registro en movimiento_proceso_externo y luego en movimiento_detalle
func RegistrarMultipleManager(movimientosDetalle []*models.MovimientoDetalle) (idRegistrados []int64) {
	o := orm.NewOrm()

	if err := o.Begin(); err == nil {
		for _, movimientoDetalle := range movimientosDetalle {
			id, err := o.Insert(movimientoDetalle.MovimientoProcesoExternoId)
			if err != nil {
				o.Rollback()
				log.Panicln(err.Error())
			}

			movimientoDetalle.MovimientoProcesoExternoId.Id = int(id)

			id, err = o.Insert(movimientoDetalle)
			if err != nil {
				o.Rollback()
				log.Panicln(err.Error())
			}

			idRegistrados = append(idRegistrados, id)
		}
	} else {
		log.Panicln(err.Error())
	}
	o.Commit()
	return
}
