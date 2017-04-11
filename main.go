package main

import "gopkg.in/kataras/iris.v5"

type ApiJson struct {
    Success bool        `json:"success"`
    Msg   string        `json:"msg"`
    Count int           `json:"count"`
}


func main() {
    app := iris.New()

    app.Get("/v1/hello", func(ctx *iris.Context) {
        ctx.JSON(iris.StatusOK, ApiJson{Success: true, Msg:"Hello", Count:18})
    })

    app.Listen("0.0.0.0:80")
}
