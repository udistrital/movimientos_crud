INSERT INTO movimientos.tipo_movimiento (
    nombre, descripcion, acronimo, activo, fecha_creacion, fecha_modificacion, parametros
)
SELECT 'AdicionCuenPre', 'Adición de movimiento de una Cuenta Presupuestal cuando se publica un Plan de Adquisiciones', 'ad_cuenpre', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, null
WHERE NOT EXISTS (
    SELECT acronimo FROM movimientos.tipo_movimiento WHERE acronimo = 'ad_cuenpre'
) LIMIT 1;
