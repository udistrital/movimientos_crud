module github.com/udistrital/movimientos_crud

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/lib/pq v1.10.0
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/udistrital/auditoria v0.0.0-20200115201815-9680ae9c2515
	github.com/udistrital/utils_oas v0.0.0-20211125230753-1091d2af48e2
)

replace github.com/astaxie/beego v1.12.3 => github.com/udistrital/beego v1.12.4-0.20211126032252-ee78ca48b207
