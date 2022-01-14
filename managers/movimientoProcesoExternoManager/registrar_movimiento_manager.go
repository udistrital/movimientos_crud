package movimientoProcesoExternoManager

import (
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/movimientos_crud/models"
)

func RegistrarMovimientoProcesoExterno(movimientoProcesoExterno *models.MovimientoProcesoExterno, movimientoDetalle *models.MovimientoDetalle) {
	o := orm.NewOrm()
	err := o.Begin()

	id, err := o.Insert(movimientoProcesoExterno)
	if err != nil {
		o.Rollback()
		panic(err)
	}
	movimientoProcesoExterno.Id = int(id)
	movimientoDetalle.MovimientoProcesoExternoId = movimientoProcesoExterno
	_, err = o.Insert(movimientoDetalle)
	if err != nil {
		o.Rollback()
		panic(err)
	}
	o.Commit()
}
