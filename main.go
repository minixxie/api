package main

import "log"
import "gopkg.in/labstack/echo.v3"
import "database/sql"
import _ "github.com/lib/pq"
// import _ "gopkg.in/go-sql-driver/mysql.v1"
import "api/endpoints"

func main() {
    db, err := sql.Open("postgres", "postgres://ldev-api:ldev-api@postgres/ldev-main?sslmode=disable")
    // db, err := sql.Open("mysql", "ldev-api:ldev-api@tcp(mariadb10:3306)/ldev-main-db?charset=utf8")
    if err != nil {
        log.Fatal("Cannot connect to DB", err)
    }
    log.Printf("main()...")

    e := echo.New()

    e.POST("/v1/orders", endpoints.CreateOrder(db))
    e.GET("/v1/orders", endpoints.GetOrders(db))

    log.Printf("listing at 0.0.0.0:80...")
    e.Logger.Fatal(e.Start(":80"))
}
