package models

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
