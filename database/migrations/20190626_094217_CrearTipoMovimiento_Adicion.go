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
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(nombre, descripcion)" +
		"VALUES('Adicion', 'Adicion a una fuente de financiamiento');")

}

// Reverse the migrations
func (m *CrearTipoMovimientoAdicion_20190626_094217) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos.tipo_movimiento")
}
