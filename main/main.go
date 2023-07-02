package main

import (
	"github.com/gin-gonic/gin"
	"tjtc/router"
)

func main() {
	// 初始化gin框架
	r := gin.Default()
	// 124
	r.LoadHTMLFiles("html/login.html")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(200, "login.html", "flysnow_org")
	})
	router.InitRouter(r)
	r.Run(":8081")
}
