CREATE DATABASE backend-golang-gin;

CREATE USER postgres WITH ENCRYPTED PASSWORD '1234';
ALTER ROLE postgres SET client_encoding TO 'utf8';
ALTER ROLE postgres SET default_transaction_isolation TO 'read committed';
ALTER ROLE postgres SET timezone TO 'UTC';

GRANT ALL PRIVILEGES ON DATABASE backend-golang-gin TO postgres;

\c backend-golang-gin;

Create TABLE IF NOT EXISTS tasks (
    id_task VARCHAR(25) NOT NULL,
    name_task VARCHAR(50) NOT NULL,
    status_task VARCHAR(50) NOT NULL
);