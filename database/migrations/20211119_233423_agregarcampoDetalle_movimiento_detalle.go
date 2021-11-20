package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarcampoDetalleMovimientoDetalle_20211119_233423 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarcampoDetalleMovimientoDetalle_20211119_233423{}
	m.Created = "20211119_233423"

	migration.Register("AgregarcampoDetalleMovimientoDetalle_20211119_233423", m)
}

// Run the migrations
func (m *AgregarcampoDetalleMovimientoDetalle_20211119_233423) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../files/20211119_233423_agregarcampoDetalle_movimiento_detalle_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *AgregarcampoDetalleMovimientoDetalle_20211119_233423) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../files/20211119_233423_agregarcampoDetalle_movimiento_detalle_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
