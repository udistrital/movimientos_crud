package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarColumnaMovimientoId_20190706_212049 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarColumnaMovimientoId_20190706_212049{}
	m.Created = "20190706_212049"

	migration.Register("ModificarColumnaMovimientoId_20190706_212049", m)
}

// Run the migrations
func (m *ModificarColumnaMovimientoId_20190706_212049) Up() {
	m.SQL("ALTER TABLE movimientos.movimiento_detalle " +
		"RENAME COLUMN movimiento_id to movimiento_proceso_externo_id;")
}

// Reverse the migrations
func (m *ModificarColumnaMovimientoId_20190706_212049) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
