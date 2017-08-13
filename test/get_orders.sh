#!/bin/bash

curl -w "\nHTTP %{http_code} time:%{time_total}s\n" -v \
    -X GET \
    -H "Content-Type: application/json; charset=utf-8" \
    http://127.0.0.1:8080/v1/orders
