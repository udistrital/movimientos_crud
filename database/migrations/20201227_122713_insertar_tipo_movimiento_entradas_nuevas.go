package main

import (
	"github.com/astaxie/beego/migration"
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
		"(id, nombre, descripcion, acronimo, activo)" +
		"VALUES(29, 'Caja menor', 'Entrada por adquisicion por caja menor', 'e_arka', true);")
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo, activo)" +
		"VALUES(24, 'Compras extranjeras', 'Entrada por adquisición por compras en el extranjero', 'e_arka', true);")
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo, activo)" +
		"VALUES(25, 'Aprovechamientos', 'Entrada por partes por aprovechamientos', 'e_arka', true);")
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo, activo)" +
		"VALUES(26, 'Adiciones y mejoras', 'Entrada por adiciones y mejoras', 'e_arka', true);")
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo, activo)" +
		"VALUES(27, 'Intangibles', 'Entrada por Intangibles adquiridos', 'e_arka', true);")
	m.SQL("INSERT INTO movimientos.tipo_movimiento" +
		"(id, nombre, descripcion, acronimo, activo)" +
		"VALUES(28, 'Provisional', 'Entrada de bienes entregados de manera provisional', 'e_arka', true);")

}

// Reverse the migrations
func (m *InsertarTipoMovimientoEntradasNuevas_20201227_122713) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos.tipo_movimiento")

}
