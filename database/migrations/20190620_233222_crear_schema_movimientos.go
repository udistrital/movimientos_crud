package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaMovimientos_20190620_233222 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaMovimientos_20190620_233222{}
	m.Created = "20190620_233222"

	migration.Register("CrearSchemaMovimientos_20190620_233222", m)
}

// Run the migrations
func (m *CrearSchemaMovimientos_20190620_233222) Up() {
	m.SQL("CREATE SCHEMA IF NOT EXISTS movimientos;")
}

// Reverse the migrations
func (m *CrearSchemaMovimientos_20190620_233222) Down() {
	m.SQL("DROP SCHEMA IF EXISTS movimientos CASCADE;")
}
