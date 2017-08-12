DROP DATABASE IF EXISTS "ldev-main";
CREATE DATABASE "ldev-main";

DROP USER IF EXISTS "ldev-main";
CREATE USER "ldev-main" WITH PASSWORD 'ldev-main';
GRANT ALL PRIVILEGES ON DATABASE "ldev-main" TO "ldev-main";




DROP DATABASE IF EXISTS "ldev-user";
CREATE DATABASE "ldev-user";

DROP USER IF EXISTS "ldev-user";
CREATE USER "ldev-user" WITH PASSWORD 'ldev-user';
GRANT ALL PRIVILEGES ON DATABASE "ldev-user" TO "ldev-user";

