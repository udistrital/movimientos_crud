package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type QuitarFechaRegistro_20190730_204356 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &QuitarFechaRegistro_20190730_204356{}
	m.Created = "20190730_204356"

	migration.Register("QuitarFechaRegistro_20190730_204356", m)
}

// Run the migrations
func (m *QuitarFechaRegistro_20190730_204356) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo DROP COLUMN IF EXISTS fecha_registro CASCADE;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle DROP COLUMN IF EXISTS fecha_registro CASCADE;")
}

// Reverse the migrations
func (m *QuitarFechaRegistro_20190730_204356) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
