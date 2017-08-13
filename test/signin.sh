#!/bin/bash

choice=$(expr $RANDOM % 3 + 1) # 1-3

postData=$(cat<<EOF
{
    "phone": "+85260000000",
    "password": "hello123"
}
EOF
)

curl -w "\nHTTP %{http_code} time:%{time_total}s\n" -v \
    -X POST \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$postData" \
    http://127.0.0.1:8080/v1/signin
