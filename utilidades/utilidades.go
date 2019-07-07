package utilidades

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/responseformat"
)

// ErrorResponse Devuelve una respuesta de tipo responseFormat cuando se realiza un panic en el proceso
func ErrorResponse(c beego.Controller) {
	if r := recover(); r != nil {
		beego.Error(r)
		responseformat.SetResponseFormat(&c, r, "E", 500)
	}
}
