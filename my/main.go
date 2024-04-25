package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.NewEngine()
	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>hello this is a html page<h1>")
	})
	engine.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello, name is %s, path is %s", ctx.Query("name"), ctx.Path)
	})
	engine.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})
	//注意这里9999前有一个冒号，如果没加就跑不成功
	engine.Run(":9999")

}
