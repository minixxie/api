package main

import "gopkg.in/labstack/echo.v3"
import "database/sql"
import _ "github.com/lib/pq"
// import _ "gopkg.in/go-sql-driver/mysql.v1"
import "api/lib"
import "api/endpoints"
// import "api/entity"
import "log"

func main() {
    config, err := lib.LoadConfig()
    if err != nil {
        log.Fatal("Cannot load config file:", err)
    }
    log.Printf("main DB: %s", config.Dbs.Main)
    log.Printf("user DB: %s", config.Dbs.User)

    // db, err := sql.Open("mysql", "ldev-api:ldev-api@tcp(mariadb10:3306)/ldev-main-db?charset=utf8")
    mainDB, err := sql.Open("postgres", config.Dbs.Main)
    if err != nil {
        log.Fatal("Cannot connect to DB (main):", err)
    }
    userDB, err := sql.Open("postgres", config.Dbs.User)
    if err != nil {
        log.Fatal("Cannot connect to DB (user):", err)
    }

    e := echo.New()

    e.POST("/v1/signup", endpoints.Signup(userDB))
    e.POST("/v1/signin", endpoints.Signin(userDB))

    e.POST("/v1/orders", endpoints.CreateOrder(mainDB))
    e.GET("/v1/orders", endpoints.GetOrders(mainDB))

    log.Printf("listing at 0.0.0.0:80...")
    e.Logger.Fatal(e.Start(":80"))
}
