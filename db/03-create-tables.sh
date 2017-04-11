#!/bin/bash

scriptPath=$(cd $(dirname $0) && pwd)

cd "$scriptPath" && ./psql.sh 03-create-tables.sql
