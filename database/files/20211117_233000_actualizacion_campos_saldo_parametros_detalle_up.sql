ALTER TABLE movimientos.movimiento_detalle ADD COLUMN saldo NUMERIC(20,7);

ALTER TABLE movimientos.movimiento_proceso_externo ADD COLUMN detalle jsonb;

ALTER TABLE movimientos.tipo_movimiento ALTER COLUMN parametros SET DATA TYPE jsonb;
