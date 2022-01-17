package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ActualizacionCamposSaldoParametrosDetalle_20211117_233000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActualizacionCamposSaldoParametrosDetalle_20211117_233000{}
	m.Created = "20211117_233000"

	migration.Register("ActualizacionCamposSaldoParametrosDetalle_20211117_233000", m)
}

// Run the migrations
func (m *ActualizacionCamposSaldoParametrosDetalle_20211117_233000) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../files/20211117_233000_actualizacion_campos_saldo_parametros_detalle_up.sql")

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
func (m *ActualizacionCamposSaldoParametrosDetalle_20211117_233000) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../files/20211117_233000_actualizacion_campos_saldo_parametros_detalle_down.sql")

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
