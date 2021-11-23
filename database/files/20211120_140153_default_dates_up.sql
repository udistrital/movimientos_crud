ALTER TABLE movimientos.tipo_movimiento
    ALTER COLUMN fecha_creacion TYPE TIMESTAMP WITH TIME ZONE,
    ALTER COLUMN fecha_modificacion TYPE TIMESTAMP WITH TIME ZONE;

ALTER TABLE movimientos.movimiento_proceso_externo
    ALTER COLUMN fecha_creacion TYPE TIMESTAMP WITH TIME ZONE,
    ALTER COLUMN fecha_modificacion TYPE TIMESTAMP WITH TIME ZONE;

ALTER TABLE movimientos.movimiento_detalle
    ALTER COLUMN fecha_creacion TYPE TIMESTAMP WITH TIME ZONE,
    ALTER COLUMN fecha_modificacion TYPE TIMESTAMP WITH TIME ZONE;

-- Se declaran funciones para fechas automaticas

CREATE OR REPLACE FUNCTION trigger_set_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
  NEW.fecha_creacion = OLD.fecha_creacion;
  NEW.fecha_modificacion = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION trigger_set_fecha_creacion()
RETURNS TRIGGER AS $$
BEGIN
  NEW.fecha_creacion = NOW();
  NEW.fecha_modificacion = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Se asocia update e insert de cada tabla con un trigger a la funcion de las fechas

CREATE TRIGGER set_fecha_modificacion_tipo_movimiento
BEFORE UPDATE ON movimientos.tipo_movimiento
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_fecha_modificacion();

CREATE TRIGGER set_fecha_creacion_tipo_movimiento
BEFORE INSERT ON movimientos.tipo_movimiento
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_fecha_creacion();

CREATE TRIGGER set_fecha_modificacion_movimiento_proceso_externo
BEFORE UPDATE ON movimientos.movimiento_proceso_externo
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_fecha_modificacion();

CREATE TRIGGER set_fecha_creacion_movimiento_proceso_externo
BEFORE INSERT ON movimientos.movimiento_proceso_externo
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_fecha_creacion();

CREATE TRIGGER set_fecha_modificacion_movimiento_detalle
BEFORE UPDATE ON movimientos.movimiento_detalle
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_fecha_modificacion();

CREATE TRIGGER set_fecha_creacion_movimiento_detalle
BEFORE INSERT ON movimientos.movimiento_detalle
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_fecha_creacion();
