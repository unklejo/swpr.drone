-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

-- This is test table. Remove this table and replace with your own tables. 
CREATE TABLE test (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) UNIQUE NOT NULL
);

-- Table to store estate information
CREATE TABLE IF NOT EXISTS estates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    width INTEGER NOT NULL,
    length INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
	updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table to store tree information within estates
CREATE TABLE IF NOT EXISTS trees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    estate_id UUID NOT NULL REFERENCES estates(id) ON DELETE CASCADE,
    x_coordinate INTEGER NOT NULL,
    y_coordinate INTEGER NOT NULL,
    height INTEGER NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
	updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (estate_id, x_coordinate, y_coordinate) -- Ensure no duplicate trees in the same plot
);

-- Indexes to improve query performance
CREATE INDEX IF NOT EXISTS idx_trees_estate_id ON trees (estate_id);

-- Table to store drone plan
CREATE TABLE drone_plans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    estate_id UUID REFERENCES estates(id) ON DELETE CASCADE,
    distance INTEGER NOT NULL,
    UNIQUE (estate_id)
);
