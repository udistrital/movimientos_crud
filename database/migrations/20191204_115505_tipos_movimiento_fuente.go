package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type TiposMovimientoFuente_20191204_115505 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TiposMovimientoFuente_20191204_115505{}
	m.Created = "20191204_115505"

	migration.Register("TiposMovimientoFuente_20191204_115505", m)
}

// Up Run the migrations
func (m *TiposMovimientoFuente_20191204_115505) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo)" +
		"VALUES(20, 'Adición Fuente', 'Adicion para una Fuente de Financiamiento', 'ad_fuente');")
	m.SQL("UPDATE movimientos.tipo_movimiento SET parametros='{\"MovOriginRubro\": \"ad_fuente_apropiacion\"}' WHERE acronimo='ad_fuente'")
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo)" +
		"VALUES(21, 'Reducción Fuente ', 'Reducción para una Fuente de Financiamiento', 'rd_fuente');")
	m.SQL("UPDATE movimientos.tipo_movimiento SET parametros='{\"MovOriginRubro\": \"rd_fuente_apropiacion\"}' WHERE acronimo='rd_fuente'")

}

// Down Reverse the migrations
func (m *TiposMovimientoFuente_20191204_115505) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos.tipo_movimiento")

}
