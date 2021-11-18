ALTER TABLE movimientos.movimiento_detalle DROP COLUMN IF EXISTS saldo;

ALTER TABLE movimientos.movimiento_proceso_externo DROP COLUMN IF EXISTS detalle;

ALTER TABLE movimientos.tipo_movimiento ALTER COLUMN parametros SET DATA TYPE json;