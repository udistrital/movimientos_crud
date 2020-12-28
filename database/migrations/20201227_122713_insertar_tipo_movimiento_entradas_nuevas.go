package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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
	file, err := ioutil.ReadFile("../files/inserts_entradas_tipo_movimiento.up.sql")

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
func (m *InsertarTipoMovimientoEntradasNuevas_20201227_122713) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
}
