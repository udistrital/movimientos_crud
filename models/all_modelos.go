package models

import (
	"time"
)

type CuentasMovimientoProcesoExterno struct {
	Cuen_Pre     string
	Mov_Proc_Ext string
	Saldo        float64
	Valor        float64
}

type Cuen_Pre struct {
	RubroId                string
	FuenteFinanciamientoId string
	ActividadId            string
}

type mov_proc_ext_con_hijos struct {
	MovimientoProcesoExterno
	Hijos []MovimientoDetalle
}

type MovimientoProcesoExternoDetallado struct {
	Id                       int             `orm:"column(id);auto"`
	TipoMovimientoId         *TipoMovimiento `orm:"column(tipo_movimiento_id);rel(fk)"`
	ProcesoExterno           int64           `orm:"column(proceso_externo)"`
	MovimientoProcesoExterno int             `orm:"column(movimiento_proceso_externo);null"`
	Activo                   bool            `orm:"column(activo);null"`
	FechaCreacion            time.Time       `orm:"auto_now_add;column(fecha_creacion);null"`
	FechaModificacion        time.Time       `orm:"auto_now;column(fecha_modificacion);null"`
	Detalle                  string          `orm:"column(detalle);type(jsonb);null"`
	MovimientoDetalle        *MovimientoDetalle
}

type RegistrarMovimientoData struct {
	MovimientoDetalle        *MovimientoDetalle
	MovimientoProcesoExterno *MovimientoProcesoExterno
}

type MovimientoDetalleInsertar struct {
	MovimientoProcesoExternoId int
	Valor                      float64
	Descripcion                string
	Activo                     bool
	Saldo                      float64
	Detalle                    string
}
