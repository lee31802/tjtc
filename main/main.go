package main

import (
	"github.com/gin-gonic/gin"
	"tjtc/router"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("html/login.html")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(200, "login.html", "flysnow_org")
	})
	router.InitRouter(r)
	r.Run(":8081")
}
