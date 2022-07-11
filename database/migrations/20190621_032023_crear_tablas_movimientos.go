package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablasMovimientos_20190621_032023 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablasMovimientos_20190621_032023{}
	m.Created = "20190621_032023"

	migration.Register("CrearTablasMovimientos_20190621_032023", m)
}

// Run the migrations
func (m *CrearTablasMovimientos_20190621_032023) Up() {
	m.SQL(" CREATE SEQUENCE movimientos.movimiento_id_seq " +
		"INCREMENT BY 1 " +
		"MINVALUE 1 " +
		"MAXVALUE 2147483647 " +
		"START WITH 1 " +
		"CACHE 1 " +
		"NO CYCLE " +
		"OWNED BY NONE ")
	m.SQL("CREATE TABLE movimientos.movimiento_proceso_externo(" +
		"id integer NOT NULL DEFAULT nextval('movimientos.movimiento_id_seq'::regclass)," +
		"tipo_movimiento_id integer NOT NULL," +
		"proceso_externo bigint NOT NULL," +
		"movimiento_proceso_externo integer," +
		"CONSTRAINT pk_movimiento PRIMARY KEY (id));")
	m.SQL("CREATE SEQUENCE movimientos.movimiento_detalle_id_seq " +
		"INCREMENT BY 1 " +
		"MINVALUE 1 " +
		"MAXVALUE 2147483647 " +
		"START WITH 1 " +
		"CACHE 1 " +
		"NO CYCLE " +
		"OWNED BY NONE;")
	m.SQL("CREATE TABLE movimientos.movimiento_detalle(" +
		"id integer NOT NULL DEFAULT nextval('movimientos.movimiento_detalle_id_seq'::regclass)," +
		"movimiento_id bigint NOT NULL," +
		"valor numeric(20,7) NOT NULL," +
		"fecha_registro date NOT NULL," +
		"descripcion character varying," +
		"CONSTRAINT pk_movimiento_detalle PRIMARY KEY (id));")
	m.SQL("CREATE SEQUENCE movimientos.tipo_movimiento_id_seq " +
		"INCREMENT BY 1 " +
		"MINVALUE 1 " +
		"MAXVALUE 2147483647 " +
		"START WITH 1 " +
		"CACHE 1 " +
		"NO CYCLE ")
	m.SQL("CREATE TABLE movimientos.tipo_movimiento(" +
		"id integer NOT NULL DEFAULT nextval('movimientos.tipo_movimiento_id_seq'::regclass)," +
		"nombre character varying(20) NOT NULL," +
		"descripcion character varying," +
		"acronimo character varying(10) NOT NULL," +
		"CONSTRAINT pk_tipo_movimiento PRIMARY KEY (id));")
	m.SQL("ALTER TABLE movimientos.movimiento_proceso_externo ADD CONSTRAINT fk_movimiento_tipo_movimiento FOREIGN KEY (tipo_movimiento_id)" +
		"REFERENCES movimientos.tipo_movimiento (id) MATCH FULL " +
		"ON DELETE RESTRICT ON UPDATE RESTRICT;")
	m.SQL("ALTER TABLE movimientos.movimiento_detalle ADD CONSTRAINT fk_movimiento_detalle_movimiento FOREIGN KEY (movimiento_id)" +
		"REFERENCES movimientos.movimiento_proceso_externo (id) MATCH FULL " +
		"ON DELETE RESTRICT ON UPDATE RESTRICT;")
}

// Reverse the migrations
func (m *CrearTablasMovimientos_20190621_032023) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
