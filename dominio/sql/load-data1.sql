-- -----------------------------------------------------------------------------
-- productos
-- -----------------------------------------------------------------------------
INSERT INTO productos (id, nombre, precio, tipo) VALUES (1, 'Arroz con pollo', 15, 'Segundo');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (2, 'Estofado de pollo', 14, 'Segundo');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (3, 'Coca Cola', 5, 'Bebida');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (4, 'Arroz con leche', 4, 'Postre');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (5, 'Inka Cola', 6, 'Bebida');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (6, 'Arroz Chaufa', 16, 'Segundo');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (7, 'Mazamorra', 7.5, 'Postre');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (8, 'Ceviche', 20, 'Segundo');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (9, 'Limonada', 5.5, 'Bebida');
INSERT INTO productos (id, nombre, precio, tipo) VALUES (10, 'Aguadito', 18, 'Sopa');

SELECT setval('productos_id_seq', (SELECT MAX(id) FROM productos));
