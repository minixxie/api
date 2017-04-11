CREATE USER 'ldev-api'@'%' IDENTIFIED BY 'ldev-api';
GRANT ALL PRIVILEGES ON `ldev-main-db`.* TO 'ldev-api'@'%';
FLUSH PRIVILEGES;
