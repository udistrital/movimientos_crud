module github.com/udistrital/movimientos_crud

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.0
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/prometheus/common v0.21.0
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/udistrital/auditoria v0.0.0-20200115201815-9680ae9c2515
	github.com/udistrital/utils_oas v0.0.0-20211125230753-1091d2af48e2
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/astaxie/beego v1.12.3 => github.com/udistrital/beego v1.12.4-0.20211126032252-ee78ca48b207
