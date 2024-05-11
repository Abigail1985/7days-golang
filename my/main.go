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
	engine.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	engine.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"filepath": c.Param("filepath")})
	})
	//注意这里9999前有一个冒号，如果没加就跑不成功
	engine.Run(":9999")

}
