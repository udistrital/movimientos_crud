package main

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/customerror"

	_ "github.com/udistrital/movimientos_crud/routers"
	"github.com/udistrital/utils_oas/xray"
)

func main() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+
		":"+url.QueryEscape(beego.AppConfig.String("PGpass"))+
		"@"+beego.AppConfig.String("PGurls")+
		":"+beego.AppConfig.String("PGport")+
		"/"+beego.AppConfig.String("PGdb")+
		"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschemas")+"")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		orm.Debug = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&customerror.CustomErrorController{})
	// logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/beego/movimientos_crud/movimientos_crud.log"}`)

	//Prueba de auditoria
	xray.InitXRay()
	auditoria.InitMiddleware()
	apistatus.Init()
	beego.Run()
}
