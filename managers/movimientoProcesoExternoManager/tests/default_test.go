package tests

import (
	"os"
	"testing"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/udistrital/movimientos_crud/managers/movimientoProcesoExternoManager"
	"github.com/udistrital/movimientos_crud/models"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+os.Getenv("MOVIMIENTOS_CRUD_DB_USER")+":"+os.Getenv("MOVIMIENTOS_CRUD_DB_PASS")+"@"+os.Getenv("MOVIMIENTOS_CRUD_DB_URL")+"/"+os.Getenv("MOVIMIENTOS_CRUD_DB_NAME")+"?sslmode=disable&search_path="+os.Getenv("MOVIMIENTOS_CRUD_DB_SCHEMA")+"")
}
func TestRegistrarMovimientoProcesoExterno(t *testing.T) {
	defer func() {
		// test fail
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		}
	}()

	tipoMovimiento := models.TipoMovimiento{
		Id: 1,
	}

	movimeintoProcesoExternoStrc := models.MovimientoProcesoExterno{
		TipoMovimientoId: &tipoMovimiento,
		ProcesoExterno:   1,
		Activo:           true,
	}

	detalleStrc := models.MovimientoDetalle{
		Valor:                      500,
		Descripcion:                "test",
		Activo:                     true,
		MovimientoProcesoExternoId: &movimeintoProcesoExternoStrc,
	}

	movimientoProcesoExternoManager.RegistrarMovimientoProcesoExterno(&movimeintoProcesoExternoStrc, &detalleStrc)

	t.Log("TestRegistrarMovimientoProcesoExterno success (OK)")
}
