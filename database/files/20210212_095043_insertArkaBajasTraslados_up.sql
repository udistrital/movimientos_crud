-- PARTE 1 - Resetear secuencia dañada por registros con IDs quemados
-- Lo siguiente debería ser una forma segura de resetear el serial
-- y permitir que de ahora en más no se admita quemar los IDs
-- (o por lo menos para movimientos.tipo_movimiento)
-- (Referencia: https://stackoverflow.com/a/244265/3180052)
-- (Otra ref: https://hcmc.uvic.ca/blogs/index.php/how_to_fix_postgresql_error_duplicate_ke?blog=22)
-- Equivale a (también funciona pero no es tan seguro):
-- ALTER SEQUENCE movimientos.tipo_movimiento RESTART WITH 14
-- Otras formas de alterar secuencias:
-- https://stackoverflow.com/questions/8745051/postgres-manually-alter-sequence
BEGIN;
LOCK TABLE movimientos.tipo_movimiento IN EXCLUSIVE MODE;
SELECT setval(
    'movimientos.tipo_movimiento_id_seq',
    COALESCE((SELECT MAX(id)+1 FROM movimientos.tipo_movimiento), 31),
    false);
COMMIT;

-- PARTE 2 - Registros en sí
INSERT INTO movimientos.tipo_movimiento (acronimo, nombre, descripcion, activo)
VALUES
('b_arka', 'Baja', 'Baja de bienes y servicios de Arka II', true),
('tr_arka', 'Traslado', 'Traslado de bienes y servicios de Arka II',  true);
