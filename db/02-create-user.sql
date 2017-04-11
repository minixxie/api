DROP USER IF EXISTS "ldev-api";
CREATE USER "ldev-api" WITH PASSWORD 'ldev-api';
GRANT ALL PRIVILEGES ON DATABASE "ldev-main" TO "ldev-api";
