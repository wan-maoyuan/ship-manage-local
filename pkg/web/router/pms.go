package router

import (
	"ship-manage-local/pkg/web/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func registerPMS(engine *gin.RouterGroup) {
	user := engine.Group("/pms")

	user.GET("/cat_log", handler.QueryPMSCatLog)
	user.GET("/group", handler.QueryPMSGroup)
	user.GET("/equipment", handler.QueryPMSEquipment)
	user.GET("/component", handler.QueryPMSComponent)

	//user.GET("/job_order", handler.QueryPMSJobOrder)
	//user.GET("/work_done", handler.QueryPMSWorkDone)
	//user.GET("/over_due_order", handler.QueryPMSWorkDone)

	logrus.Info("pms 路由注册成功！！！")
}
