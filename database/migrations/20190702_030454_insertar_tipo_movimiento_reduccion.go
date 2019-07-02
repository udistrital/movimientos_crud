package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertarTipoMovimientoReduccion_20190702_030454 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTipoMovimientoReduccion_20190702_030454{}
	m.Created = "20190702_030454"

	migration.Register("InsertarTipoMovimientoReduccion_20190702_030454", m)
}

// Run the migrations
func (m *InsertarTipoMovimientoReduccion_20190702_030454) Up() {
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo)" +
		"VALUES(2, 'Reducción', 'Reducción a una fuente de financiamiento', 'reduccion');")
}

// Reverse the migrations
func (m *InsertarTipoMovimientoReduccion_20190702_030454) Down() {
	m.SQL("DELETE FROM movimientos.tipo_movimiento")
}
