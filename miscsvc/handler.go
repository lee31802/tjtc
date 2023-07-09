package miscsvc

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tjtc/proto/model"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Code:    1,
		Message: "pong",
	})
}

func Login(c *gin.Context) {
	request := model.UserLoginRequest{}
	//返回错误，并在header里面写400的状态码.根据Body数据类型, 将数据赋值到指定的结构体变量中. (类似于序列化和反序列化);
	//ShouldBindJSON()只会返回错误信息，不会往header里面写400的错误状态码
	c.BindJSON(&request)
	logrus.Info("login request:%v", request)
	account, err := loginSvc(request.UserId, request.Password, request.Type)
	if err != nil {
		logrus.Error("Login loginSvc err :%v", err)
	}
	logrus.Info("login account:%v", account)
	c.JSON(http.StatusOK, model.Response{
		Code:    1,
		Data:    request,
		Message: "pong",
	})
	return
}
