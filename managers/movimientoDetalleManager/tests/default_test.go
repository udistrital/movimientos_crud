package tests

import (
	"os"
	"testing"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	movimientodetallemanager "github.com/udistrital/movimientos_crud/managers/movimientoDetalleManager"
	"github.com/udistrital/movimientos_crud/models"
)

var insertMultipleResul []int64

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+os.Getenv("MOVIMIENTOS_CRUD_DB_USER")+":"+os.Getenv("MOVIMIENTOS_CRUD_DB_PASS")+"@"+os.Getenv("MOVIMIENTOS_CRUD_DB_URL")+"/"+os.Getenv("MOVIMIENTOS_CRUD_DB_NAME")+"?sslmode=disable&search_path="+os.Getenv("MOVIMIENTOS_CRUD_DB_SCHEMA")+"")
}
func TestRegistrarMultipleManager(t *testing.T) {
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

	var input []*models.MovimientoDetalle

	input = append(input, &detalleStrc)

	insertMultipleResul = movimientodetallemanager.RegistrarMultipleManager(input)
	if len(insertMultipleResul) == 0 {
		panic("error at registration func")
	}
	t.Log("Movimiento RegistrarMultipleManager success (OK)")
}

func TestEliminarMultiple(t *testing.T) {
	defer func() {
		// test fail
		if r := recover(); r != nil {
			t.Error("error: ", r)
			t.Fail()
		}
	}()
	var input []int
	input = append(input, int(insertMultipleResul[0]))
	if err := movimientodetallemanager.EliminarMultipleManager(input); err != nil {
		panic(err.Error())
	}

	t.Log("Delete func success (OK)")
}
