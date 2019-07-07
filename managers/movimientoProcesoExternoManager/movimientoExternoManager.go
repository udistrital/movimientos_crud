package movimientoProcesoExternoManager

// import (
// 	"log"
// 	"time"
// 	"github.com/astaxie/beego/orm"
// 	"github.com/udistrital/movimientos_crud/models"
// )

// // RegistrarMultipleManager transacción para registrar multiples registros en movimiento_proceso_extenro
// // y movimiento_detalle
// func RegistrarMultipleManager(movimientosExternos []*models.MovimientoProcesoExterno) (idRegistrados []int64) {
// 	o := orm.NewOrm()

// 	if err := o.Begin(); err == nil {
// 		for _, movimientoExterno := range movimientosExternos {
// 			id, err := o.Insert(movimientoExterno)
// 			if err != nil {
// 				o.Rollback()
// 				log.Panicln(err.Error())
// 			}
// 			movimientoExterno.Id = id

// 			movimientoDetalle := models.MovimientoDetalle {
// 				MovimientoId : movimientoExterno,
// 				Valor: 0,
// 				FechaRegistro: time.Now(),
// 				Descripcion: ""
// 			}

// 			_, err := models.MovimientoDetalle.AddMovimientoDetalle()
// 			if err != nil {
// 				o.Rollback()
// 				log.Panicln(err.Error())
// 			}
// 		}
// 	}
}

