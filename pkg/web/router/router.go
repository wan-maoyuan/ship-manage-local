package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterRouter(engine *gin.Engine) {
	api := engine.Group("/api")

	registerUser(api)
	registerPMS(api)

	logrus.Info("所有路由注册成功！！！")
}
