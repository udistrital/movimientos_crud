package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertartTipoMovimientoSuspension_20190714_163618 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertartTipoMovimientoSuspension_20190714_163618{}
	m.Created = "20190714_163618"

	migration.Register("InsertartTipoMovimientoSuspension_20190714_163618", m)
}

// Run the migrations
func (m *InsertartTipoMovimientoSuspension_20190714_163618) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo)" +
		"VALUES(5, 'Suspension', 'Suspension a un valor de fuente de financiamiento', 'suspension');")
}

// Reverse the migrations
func (m *InsertartTipoMovimientoSuspension_20190714_163618) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
