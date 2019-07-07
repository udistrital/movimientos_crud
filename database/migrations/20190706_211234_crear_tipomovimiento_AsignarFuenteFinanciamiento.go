package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTipomovimientoAsignarFuenteFinanciamiento_20190706_211234 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTipomovimientoAsignarFuenteFinanciamiento_20190706_211234{}
	m.Created = "20190706_211234"

	migration.Register("CrearTipomovimientoAsignarFuenteFinanciamiento_20190706_211234", m)
}

// Run the migrations
func (m *CrearTipomovimientoAsignarFuenteFinanciamiento_20190706_211234) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("insert into movimientos.tipo_movimiento (id, nombre, descripcion, acronimo)" +
		"values (3, 'AsignarFuenteFinanciamiento', 'Asigna valores de la fuente a sus dependencias de acuerdo a la apropiación del rubro', 'crear_ff');")
}

// Reverse the migrations
func (m *CrearTipomovimientoAsignarFuenteFinanciamiento_20190706_211234) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
