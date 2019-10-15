package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarParametrosJsonTipoMovimiento_20191007_074007 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarParametrosJsonTipoMovimiento_20191007_074007{}
	m.Created = "20191007_074007"

	migration.Register("AgregarParametrosJsonTipoMovimiento_20191007_074007", m)
}

// Run the migrations
func (m *AgregarParametrosJsonTipoMovimiento_20191007_074007) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE movimientos.tipo_movimiento ADD COLUMN parametros json")
	m.SQL("UPDATE movimientos.tipo_movimiento SET parametros='{\"CuentaContraCredito\": true, \"TipoMovimientoCuentaCredito\": \"traslado_destino\"}' WHERE acronimo='traslado'")

}

// Reverse the migrations
func (m *AgregarParametrosJsonTipoMovimiento_20191007_074007) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
