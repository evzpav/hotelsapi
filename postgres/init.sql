CREATE USER "hotelsapi";
ALTER USER "hotelsapi" WITH ENCRYPTED PASSWORD 'hotelsapipass';
CREATE DATABASE "hotelsapi";
GRANT ALL PRIVILEGES ON DATABASE "hotelsapi" TO "hotelsapi";
\c "hotelsapi"
CREATE EXTENSION pgcrypto;