#!/bin/bash

postgis=mdillon/postgis:9.6-alpine

docker run -it --rm --network=host -v "$PWD"/01-create-db.sql:/file.sql $postgis bash -c "psql -h localhost -U postgres -W -f /file.sql"
