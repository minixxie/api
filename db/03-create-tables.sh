#!/bin/bash

scriptPath=$(cd $(dirname $0) && pwd)

postgis=postgres:9.6-alpine

cd "$scriptPath"

docker run -t --rm --network=host -v "$PWD"/03-create-tables-main.sql:/file.sql $postgis bash -c "PGPASSWORD=ldev-api psql -h localhost -U ldev-api -f /file.sql \"ldev-main\""
docker run -t --rm --network=host -v "$PWD"/03-create-tables-user.sql:/file.sql $postgis bash -c "PGPASSWORD=ldev-api psql -h localhost -U ldev-api -f /file.sql \"ldev-user\""