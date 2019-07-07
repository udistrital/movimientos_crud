package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoDetalleController"],
        beego.ControllerComments{
            Method: "RegistrarMultiple",
            Router: `/registrar_multiple`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:MovimientoProcesoExternoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_crud/controllers:TipoMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
