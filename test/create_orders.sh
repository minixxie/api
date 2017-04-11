#!/bin/bash

choice=$(expr $RANDOM % 3 + 1) # 1-3

case "$choice" in
	1)
		title="Send document to Solicitor"
		orderCategoryId=3
	;;
	2)
		title="Buy fried rice"
		orderCategoryId=2
	;;
	3)
		title="Buy instant noodles"
		orderCategoryId=2
	;;
esac


postData=$(cat<<EOF
{
    "title": "$title",
    "orderCategoryId": $orderCategoryId
}
EOF
)

curl -w "\nHTTP %{http_code} time:%{time_total}s\n" -v \
    -X POST \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$postData" \
    http://127.0.0.1:8080/v1/orders
