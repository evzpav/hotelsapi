CREATE USER "hotelsapi";
ALTER USER "hotelsapi" WITH ENCRYPTED PASSWORD 'hotelsapipass';
CREATE DATABASE "hotelsapi";
GRANT ALL PRIVILEGES ON DATABASE "hotelsapi" TO "hotelsapi";
\c "hotelsapi"
CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS hotels (
    id SERIAL PRIMARY KEY,
	name TEXT,
	active BOOLEAN,
    city TEXT,
    nr_of_employees INT,
    revenue DECIMAL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);