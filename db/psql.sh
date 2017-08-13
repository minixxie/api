#!/bin/bash

if [ "$1" == "" ]
then
    echo >&2 "usage: $0 file.sql"
    exit 1
fi

postgis=mdillon/postgis:9.6-alpine

docker run -t --rm --network=host -v "$PWD"/"$1":/file.sql $postgis bash -c "PGPASSWORD=ldev-api psql -h localhost -U ldev-api -f /file.sql \"ldev-main\""
