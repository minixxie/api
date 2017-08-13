#!/bin/bash

docker exec -it postgres bash -c "PGPASSWORD=ldev-api psql -U ldev-api \"ldev-user\""
