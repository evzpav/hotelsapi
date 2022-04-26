CREATE USER "hotelscrud";
ALTER USER "hotelscrud" WITH ENCRYPTED PASSWORD 'hotelscrudpass';
CREATE DATABASE "hotelscrud";
GRANT ALL PRIVILEGES ON DATABASE "hotelscrud" TO "hotelscrud";
\c "hotelscrud"
CREATE EXTENSION pgcrypto;