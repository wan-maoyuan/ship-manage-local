package router

import (
	"ship-manage-local/pkg/web/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func registerUser(engine *gin.RouterGroup) {
	user := engine.Group("/user")

	user.POST("/login", handler.Login) // 用户登录
	user.POST("/create", nil)          // 用户注册
	user.POST("/delete", nil)          // 删除用户信息
	user.POST("/update", nil)          // 更新用户信息

	logrus.Info("user 路由注册成功！！！")
}
