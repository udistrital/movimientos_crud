module github.com/udistrital/movimientos_crud

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/imdario/mergo v0.3.12
	github.com/lib/pq v1.10.5
	github.com/prometheus/common v0.33.0 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/udistrital/utils_oas v0.0.0-20220415063412-5dae4bd58180
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/astaxie/beego v1.12.3 => github.com/udistrital/beego v1.12.4-0.20211126032252-ee78ca48b207
