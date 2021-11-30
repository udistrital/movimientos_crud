package models

type CuentasMovimientoProcesoExterno struct {
	Cuen_Pre     string
	Mov_Proc_Ext string
}

type Cuen_Pre struct {
	RubroId                string
	FuenteFinanciamientoId string
	ActividadId            string
}
