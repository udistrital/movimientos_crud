DELETE FROM movimientos.tipo_movimiento WHERE acronimo = 'b_arka'
AND EXISTS (SELECT 1 FROM movimientos.tipo_movimiento WHERE acronimo = 'b_arka');
DELETE FROM movimientos.tipo_movimiento WHERE acronimo = 'tr_arka'
AND EXISTS (SELECT 1 FROM movimientos.tipo_movimiento WHERE acronimo = 'tr_arka');
