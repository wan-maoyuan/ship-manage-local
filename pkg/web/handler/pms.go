package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ship-manage-local/pkg/database"
	"ship-manage-local/pkg/web/proto"
)

func QueryPMSCatLog(c *gin.Context) {
	catLogList, err := database.QueryPMSCatLog("")
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(catLogList),
	)
}

func QueryPMSGroup(c *gin.Context) {
	var err error

	var req = proto.QueryPMSGroupReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	list, err := database.QueryPMSGroup(req.CatLogID, req.GroupID)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(list),
	)
}

func QueryPMSEquipment(c *gin.Context) {
	var err error

	var req = proto.QueryPMSEquipmentReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	list, err := database.QueryPMSEquipment(req.CatLogID, req.GroupID, req.EquipmentID)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(list),
	)
}

func QueryPMSComponent(c *gin.Context) {
	var err error

	var req = proto.QueryPMSComponentReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	list, err := database.QueryPMSComponent(req.CatLogID, req.GroupID, req.EquipmentID, req.ComponentID)
	if err != nil {
		c.JSON(
			http.StatusOK,
			proto.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		proto.NewSuccesResponse(list),
	)
}
