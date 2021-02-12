package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertArkaBajasTraslados_20210212_095043 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertArkaBajasTraslados_20210212_095043{}
	m.Created = "20210212_095043"

	migration.Register("InsertArkaBajasTraslados_20210212_095043", m)
}

// Run the migrations
func (m *InsertArkaBajasTraslados_20210212_095043) Up() {

	file, err := ioutil.ReadFile("../files/20210212_095043_insertArkaBajasTraslados_up.sql")

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
func (m *InsertArkaBajasTraslados_20210212_095043) Down() {

	file, err := ioutil.ReadFile("../files/20210212_095043_insertArkaBajasTraslados_down.sql")

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
