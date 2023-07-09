package router

import (
	"github.com/gin-gonic/gin"
	"tjtc/miscsvc"
)

func InitRouter(engine *gin.Engine) {
	engine.POST("/user/login", miscsvc.Login)
	engine.GET("/ping", miscsvc.Ping)
}
