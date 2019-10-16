package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarNombrefkFkDetalleMovimientoMovimiento_20190706_211102 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarNombrefkFkDetalleMovimientoMovimiento_20190706_211102{}
	m.Created = "20190706_211102"

	migration.Register("ModificarNombrefkFkDetalleMovimientoMovimiento_20190706_211102", m)
}

// Run the migrations
func (m *ModificarNombrefkFkDetalleMovimientoMovimiento_20190706_211102) Up() {
	m.SQL("ALTER TABLE movimientos.movimiento_detalle RENAME CONSTRAINT fk_movimiento_detalle_movimiento to fk_movimiento_detalle_movimiento_proceso_externo")
}

// Reverse the migrations
func (m *ModificarNombrefkFkDetalleMovimientoMovimiento_20190706_211102) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
