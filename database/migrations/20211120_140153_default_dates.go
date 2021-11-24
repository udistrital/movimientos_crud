package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type DefaultDates_20211120_140153 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DefaultDates_20211120_140153{}
	m.Created = "20211120_140153"

	migration.Register("DefaultDates_20211120_140153", m)
}

// Run the migrations
func (m *DefaultDates_20211120_140153) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../files/20211120_140153_default_dates_up.sql")

	if err != nil {
		fmt.Println(err)
	}

	// Procedimiento adicional para ejecutar correctamente las funciones del script

	a := regexp.MustCompile(`\$\$`).FindAllIndex([]byte(string(file)), -1)
	var rangesTriggers []int
	for _, i := range a {
		rangesTriggers = append(rangesTriggers, i[0])
	}

	for i, ch := range string(file) {
		if ch == 59 {
			for j := 0; j < len(rangesTriggers)-1; j = j + 2 {
				if i > rangesTriggers[j] && i < rangesTriggers[j+1] {
					file[i] = 35
				}
			}
		}
	}
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		request = strings.ReplaceAll(request, "#", ";")
		fmt.Println(request)
		m.SQL(request)
	}

}

// Reverse the migrations
func (m *DefaultDates_20211120_140153) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../files/20211120_140153_default_dates_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}
