package main

import (
	"github.com/astaxie/beego/client/orm/migration"
)

// DO NOT MODIFY
type InsertarTipoMovimientoEntradasNuevas_20201227_122713 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTipoMovimientoEntradasNuevas_20201227_122713{}
	m.Created = "20201227_122713"

	migration.Register("InsertarTipoMovimientoEntradasNuevas_20201227_122713", m)
}

// Run the migrations
func (m *InsertarTipoMovimientoEntradasNuevas_20201227_122713) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(nombre, descripcion, acronimo, activo, parametros)" +
		"VALUES('Adquisición por caja menor', 'Entrada por adquisición por caja menor', 'e_arka', true, '')," +
		"(Adquisición por compras en el extranjero, 'Entrada por adquisición por compras en el extranjero', 'e_arka', 'true', '')," +
		"(Partes por aprovechamientos, 'Entrada por partes por aprovechamientos', 'e_arka', 'true', '')," +
		"(Adiciones y mejoras, 'Entrada por adiciones y mejoras', 'e_arka', 'true', '')," +
		"(Intangibles adquiridos, 'Entrada por Intangibles adquiridos', 'e_arka', 'true', '')," +
		"(Provisional, 'Entrada de bienes entregados de manera provisional', 'e_arka', 'true', '')")

}

// Reverse the migrations
func (m *InsertarTipoMovimientoEntradasNuevas_20201227_122713) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos.tipo_movimiento")

}
