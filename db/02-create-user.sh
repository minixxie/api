#!/bin/bash

postgis=mdillon/postgis:9.6-alpine

docker run -it --rm --network=host -v "$PWD"/02-create-user.sql:/file.sql $postgis bash -c "psql -h localhost -U postgres -W -f /file.sql"
