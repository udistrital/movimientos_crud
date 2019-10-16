package movimientodetallemanager

import (
	"log"

	"github.com/astaxie/beego/logs"

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

// EliminarMultipleManager realiza multiples borrados en una transacción sobre las tablas:
// movimiento_proceso_externo y movimiento_detalle
// realizando primero un registro en movimiento_proceso_externo y luego en movimiento_detalle
func EliminarMultipleManager(IDSmovimientosDetalle []int) (err error) {
	o := orm.NewOrm()
	v := models.MovimientoDetalle{}
	movimientosToDelete := make(map[int]*models.MovimientoProcesoExterno)
	if err = o.Begin(); err == nil {
		defer func() {
			if r := recover(); r != nil {
				o.Rollback()
				logs.Error(r)
			}
		}()
		for _, id := range IDSmovimientosDetalle {
			v.Id = id
			if err = o.Read(&v); err == nil {

				movimientosToDelete[v.MovimientoProcesoExternoId.Id] = v.MovimientoProcesoExternoId

				var num int64

				if num, err = o.Delete(&models.MovimientoDetalle{Id: id}); err == nil {
					logs.Info("Number of MovimientoDetalle deleted in database:", num)
				} else {
					log.Panicln(err.Error())
				}

			} else {
				log.Panicln(err.Error())
			}
		}

		for _, movimiento := range movimientosToDelete {
			if err = o.Read(movimiento); err == nil {

				var num int64
				if num, err = o.Delete(movimiento); err == nil {
					logs.Info("Number of MovimientoProcesoExterno deleted in database:", num)
				} else {
					log.Panicln(err.Error())
				}

			} else {
				log.Panicln(err.Error())
			}
		}

	} else {
		log.Panicln(err.Error())
	}
	o.Commit()
	return err
}
