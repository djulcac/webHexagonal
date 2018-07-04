-- Drop foreign keys
-- Example:
-- ALTER TABLE transactions
--   DROP CONSTRAINT transactions_account_id_foreign;

-- Drop tables
DROP TABLE IF EXISTS productos;

-- Create tables
CREATE TABLE productos (
  id SERIAL PRIMARY KEY,
  nombre VARCHAR(64) NOT NULL,
  precio FLOAT NOT NULL,
  tipo VARCHAR(64)
);
ALTER TABLE productos
  ADD CONSTRAINT productos_name_unique UNIQUE (nombre);

-- Add foreign keys
-- Example:
-- ALTER TABLE transactions
--   ADD COLUMN account_id INTEGER NOT NULL,
--   ADD COLUMN category_id INTEGER NOT NULL;
