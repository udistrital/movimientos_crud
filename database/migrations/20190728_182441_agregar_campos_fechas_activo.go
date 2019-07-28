package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCamposFechasActivo_20190728_182441 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCamposFechasActivo_20190728_182441{}
	m.Created = "20190728_182441"

	migration.Register("AgregarCamposFechasActivo_20190728_182441", m)
}

// Run the migrations
func (m *AgregarCamposFechasActivo_20190728_182441) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo ADD COLUMN activo boolean;")
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo ADD COLUMN fecha_creacion timestamp;")
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo ADD COLUMN fecha_modificacion timestamp;")
	m.SQL("ALTER TABLE movimientos.tipo_movimiento ADD COLUMN activo boolean;")
	m.SQL("ALTER TABLE movimientos.tipo_movimiento ADD COLUMN fecha_creacion timestamp;")
	m.SQL("ALTER TABLE movimientos.tipo_movimiento ADD COLUMN fecha_modificacion timestamp;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle ADD COLUMN activo boolean;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle ADD COLUMN fecha_creacion timestamp;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle ADD COLUMN fecha_modificacion timestamp;")
}

// Reverse the migrations
func (m *AgregarCamposFechasActivo_20190728_182441) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo DROP COLUMN IF EXISTS activo CASCADE;")
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo DROP COLUMN IF EXISTS fecha_creacion CASCADE;")
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo DROP COLUMN IF EXISTS fecha_modificacion CASCADE;")
	m.SQL("ALTER TABLE movimientos.tipo_movimiento DROP COLUMN IF EXISTS activo CASCADE;")
	m.SQL("ALTER TABLE movimientos.tipo_movimiento DROP COLUMN IF EXISTS fecha_creacion CASCADE;")
	m.SQL("ALTER TABLE movimientos.tipo_movimiento DROP COLUMN IF EXISTS fecha_modificacion CASCADE;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle DROP COLUMN IF EXISTS activo CASCADE;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle DROP COLUMN IF EXISTS fecha_creacion CASCADE;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle DROP COLUMN IF EXISTS fecha_modificacion CASCADE;")
}
