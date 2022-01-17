package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertEntradaDesarrolloIntangible_20210128_114137 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertEntradaDesarrolloIntangible_20210128_114137{}
	m.Created = "20210128_114137"

	migration.Register("InsertEntradaDesarrolloIntangible_20210128_114137", m)
}

// Run the migrations
func (m *InsertEntradaDesarrolloIntangible_20210128_114137) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../files/insert_tipomovimiento_desarrollo_intangible.up.sql")

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
func (m *InsertEntradaDesarrolloIntangible_20210128_114137) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
