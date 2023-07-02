package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.POST("/user/login", login)
	engine.GET("/ping", ping)
}
