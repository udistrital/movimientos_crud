-- Se revierte ajuste al modelo de datos

ALTER TABLE movimientos.tipo_movimiento
    ALTER COLUMN fecha_creacion TYPE TIMESTAMP WITHOUT TIME ZONE,
    ALTER COLUMN fecha_modificacion TYPE TIMESTAMP WITHOUT TIME ZONE;

ALTER TABLE movimientos.movimiento_proceso_externo
    ALTER COLUMN fecha_creacion TYPE TIMESTAMP WITHOUT TIME ZONE,
    ALTER COLUMN fecha_modificacion TYPE TIMESTAMP WITHOUT TIME ZONE;

ALTER TABLE movimientos.movimiento_detalle
    ALTER COLUMN fecha_creacion TYPE TIMESTAMP WITHOUT TIME ZONE,
    ALTER COLUMN fecha_modificacion TYPE TIMESTAMP WITHOUT TIME ZONE;

-- Se revierte la función de las fechas automaticas

DROP FUNCTION trigger_set_fecha_modificacion() cascade;
DROP FUNCTION trigger_set_fecha_creacion() cascade;
