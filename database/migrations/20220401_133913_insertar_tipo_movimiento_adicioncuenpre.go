package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertarTipoMovimientoAdicioncuenpre_20220401_133913 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertarTipoMovimientoAdicioncuenpre_20220401_133913{}
	m.Created = "20220401_133913"

	migration.Register("InsertarTipoMovimientoAdicioncuenpre_20220401_133913", m)
}

// Run the migrations
func (m *InsertarTipoMovimientoAdicioncuenpre_20220401_133913) Up() {
	file, err := ioutil.ReadFile("../files/20220401_133913_insertar_tipo_movimiento_adicioncuenpre_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *InsertarTipoMovimientoAdicioncuenpre_20220401_133913) Down() {
	file, err := ioutil.ReadFile("../files/20220401_133913_insertar_tipo_movimiento_adicioncuenpre_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
