-- ============================================================
<<<<<<< HEAD
-- 001_initial.sql (SAFE REFACTOR)
=======
-- 001_initial.sql
-- Initial database setup 
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
-- ============================================================

-- =====================
-- 1. Create Tables
-- =====================

CREATE TABLE IF NOT EXISTS rees (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
<<<<<<< HEAD
    symbol TEXT UNIQUE NOT NULL,
    category TEXT NOT NULL,
    market_price NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
=======
    market_price NUMERIC NOT NULL,
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
<<<<<<< HEAD
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE,
    password TEXT NOT NULL,
    role TEXT DEFAULT 'trader',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
=======
    username TEXT NOT NULL
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
);

CREATE TABLE IF NOT EXISTS inventories (
    id SERIAL PRIMARY KEY,
    element_id INT NOT NULL,
<<<<<<< HEAD
    quantity NUMERIC NOT NULL,
    warehouse_location TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
=======
    quantity INT NOT NULL
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    element_id INT NOT NULL,
    type TEXT NOT NULL,
    quantity NUMERIC NOT NULL,
<<<<<<< HEAD
    price NUMERIC NOT NULL,
    warehouse_location TEXT, -- initially nullable for safe migration
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- =====================
-- 2. Safe Schema Updates
-- =====================

-- Drop old constraints
=======
    price NUMERIC NOT NULL
);

-- =====================
-- 2. Drop old constraints/indexes if they exist
-- =====================

>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
ALTER TABLE rees DROP CONSTRAINT IF EXISTS rees_name_key;

ALTER TABLE transactions DROP CONSTRAINT IF EXISTS chk_transaction_type;
ALTER TABLE transactions DROP CONSTRAINT IF EXISTS chk_quantity_positive;
ALTER TABLE transactions DROP CONSTRAINT IF EXISTS chk_price_positive;
<<<<<<< HEAD
ALTER TABLE transactions DROP CONSTRAINT IF EXISTS chk_warehouse_not_empty;

ALTER TABLE inventories DROP CONSTRAINT IF EXISTS chk_inventory_quantity_non_negative;
=======

ALTER TABLE inventories DROP CONSTRAINT IF EXISTS chk_inventory_quantity_non_negative;

>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
ALTER TABLE rees DROP CONSTRAINT IF EXISTS chk_market_price_positive;

ALTER TABLE inventories DROP CONSTRAINT IF EXISTS fk_inventory_element;
ALTER TABLE transactions DROP CONSTRAINT IF EXISTS fk_transaction_user;
ALTER TABLE transactions DROP CONSTRAINT IF EXISTS fk_transaction_element;

<<<<<<< HEAD
-- Add REE extra fields
ALTER TABLE rees ADD COLUMN IF NOT EXISTS form TEXT DEFAULT 'oxide';
ALTER TABLE rees ADD COLUMN IF NOT EXISTS price_unit TEXT DEFAULT 'USD/kg';
ALTER TABLE rees ADD COLUMN IF NOT EXISTS purity NUMERIC DEFAULT 99.9;

-- Add warehouse column safely (NO default hack)
ALTER TABLE transactions ADD COLUMN IF NOT EXISTS warehouse_location TEXT;

-- =====================
-- 3. Backfill Data (IMPORTANT)
-- =====================

-- Prevent NULLs for existing rows
UPDATE transactions
SET warehouse_location = 'LEGACY'
WHERE warehouse_location IS NULL;

-- =====================
-- 4. Enforce NOT NULL AFTER BACKFILL
-- =====================

ALTER TABLE transactions
ALTER COLUMN warehouse_location SET NOT NULL;

-- =====================
-- 5. Indexes
=======
-- =====================
-- 3. Indexes
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
-- =====================

CREATE UNIQUE INDEX IF NOT EXISTS idx_rees_name_unique
ON rees(name)
WHERE deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_inventory_element_id ON inventories(element_id);
<<<<<<< HEAD
CREATE INDEX IF NOT EXISTS idx_inventory_location ON inventories(warehouse_location);

CREATE INDEX IF NOT EXISTS idx_transaction_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transaction_element_id ON transactions(element_id);

-- 🔥 NEW: important for queries
CREATE INDEX IF NOT EXISTS idx_transaction_element_location
ON transactions(element_id, warehouse_location);

-- =====================
-- 6. Constraints
=======
CREATE INDEX IF NOT EXISTS idx_transaction_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transaction_element_id ON transactions(element_id);

-- =====================
-- 4. Constraints
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
-- =====================

ALTER TABLE transactions
ADD CONSTRAINT chk_transaction_type
CHECK (type IN ('buy', 'sell'));

ALTER TABLE transactions
ADD CONSTRAINT chk_quantity_positive
CHECK (quantity > 0);

ALTER TABLE transactions
ADD CONSTRAINT chk_price_positive
CHECK (price > 0);

<<<<<<< HEAD
-- Better constraint than just <> ''
ALTER TABLE transactions
ADD CONSTRAINT chk_warehouse_not_empty
CHECK (length(trim(warehouse_location)) > 0);

=======
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
ALTER TABLE inventories
ADD CONSTRAINT chk_inventory_quantity_non_negative
CHECK (quantity >= 0);

ALTER TABLE rees
ADD CONSTRAINT chk_market_price_positive
CHECK (market_price > 0);

<<<<<<< HEAD
CREATE UNIQUE INDEX IF NOT EXISTS idx_inventory_unique
ON inventories (element_id, warehouse_location)
WHERE deleted_at IS NULL;

-- =====================
-- 7. Foreign Keys
=======
-- =====================
-- 5. Foreign Keys
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
-- =====================

ALTER TABLE inventories
ADD CONSTRAINT fk_inventory_element
FOREIGN KEY (element_id) REFERENCES rees(id);

ALTER TABLE transactions
ADD CONSTRAINT fk_transaction_user
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE transactions
ADD CONSTRAINT fk_transaction_element
<<<<<<< HEAD
FOREIGN KEY (element_id) REFERENCES rees(id);
=======
FOREIGN KEY (element_id) REFERENCES rees(id);
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
