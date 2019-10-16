package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertarTipoMovimientoTraslado_20190704_152011 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTipoMovimientoTraslado_20190704_152011{}
	m.Created = "20190704_152011"

	migration.Register("InsertarTipoMovimientoTraslado_20190704_152011", m)
}

// Run the migrations
func (m *InsertarTipoMovimientoTraslado_20190704_152011) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo)" +
		"VALUES(3, 'Traslado', 'Traslado entre 2 fuentes de financiamiento', 'traslado');")

}

// Reverse the migrations
func (m *InsertarTipoMovimientoTraslado_20190704_152011) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos.tipo_movimiento")

}
