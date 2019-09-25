package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertarTiposMovimientoCdpRp_20190924_210421 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTiposMovimientoCdpRp_20190924_210421{}
	m.Created = "20190924_210421"

	migration.Register("InsertarTiposMovimientoCdpRp_20190924_210421", m)
}

// Run the migrations
func (m *InsertarTiposMovimientoCdpRp_20190924_210421) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo)" +
		"VALUES(6, 'CDP', 'Certificado Disponibilidad Presupuestal', 'cdp')," +
		"(7, 'RP', 'Registro Presupuestal', 'rp')," +
		"(8, 'Anulación CDP', 'Anulación Certificado Disponibilidad Presupuestal', 'anul_cdp')," +
		"(9, 'Anulación RP', 'Anulación Registro Presupuestal', 'anul_rp')")

}

// Reverse the migrations
func (m *InsertarTiposMovimientoCdpRp_20190924_210421) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
