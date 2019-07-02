package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTipoMovimientoAdicion_20190626_094217 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTipoMovimientoAdicion_20190626_094217{}
	m.Created = "20190626_094217"

	migration.Register("CrearTipoMovimientoAdicion_20190626_094217", m)
}

// Run the migrations
func (m *CrearTipoMovimientoAdicion_20190626_094217) Up() {
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo)" +
		"VALUES(1, 'Adicion', 'Adicion a una fuente de financiamiento', 'adicion');")

}

// Reverse the migrations
func (m *CrearTipoMovimientoAdicion_20190626_094217) Down() {
	m.SQL("DELETE FROM movimientos.tipo_movimiento")
}
