package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablasMovimientos_20190621_032023 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablasMovimientos_20190621_032023{}
	m.Created = "20190621_032023"

	migration.Register("CrearTablasMovimientos_20190621_032023", m)
}

// Run the migrations
func (m *CrearTablasMovimientos_20190621_032023) Up() {
	file, err := ioutil.ReadFile("../files/crear_tablas_movimientos.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablasMovimientos_20190621_032023) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
