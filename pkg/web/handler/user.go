package handler

import (
	"net/http"
	"ship-manage-local/pkg/web/proto"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var err error

	var req = proto.LoginReq{}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(&proto.LoginResp{
			UserID:      1,
			Account:     "admin",
			Password:    "admin",
			RoleID:      1,
			Role:        "admin",
			Permissions: []string{"*.*.*"},
		}),
	)
}
