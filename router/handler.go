package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tjtc/proto/model"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Code:    1,
		Message: "pong",
	})
}

func login(c *gin.Context) {
	request := model.UserLoginRequest{}
	//返回错误，并在header里面写400的状态码.根据Body数据类型, 将数据赋值到指定的结构体变量中. (类似于序列化和反序列化);
	//ShouldBindJSON()只会返回错误信息，不会往header里面写400的错误状态码
	c.BindJSON(&request)
	logrus.Info("login request:%v", request)
	//都数据库
	c.JSON(http.StatusOK, model.Response{
		Code:    1,
		Data:    request,
		Message: "pong",
	})
}
