package handler

import (
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/repository"
	"fuel-price/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type stationHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newStationHandler(h *Handler) *stationHandler {
	return &stationHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *stationHandler) register() {

	group := ctr.R.Group("/api/stations")
	group.POST("", ctr.getStations)
	group.POST("/create", ctr.createStation)
	group.POST("/update", ctr.updateStation)
	group.POST("/delete", ctr.deleteStation)
}

func (ctr *stationHandler) getStations(c *gin.Context) {
	req := dto.SearchStation{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	list, total, err := ctr.repo.Station.List(c.Request.Context(), &req)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}

	data := gin.H{
		"list":  list,
		"total": total,
	}
	res := utils.GenerateSuccessResponse(data)

	c.JSON(res.HttpStatusCode, res)
}

func (ctr stationHandler) createStation(c *gin.Context) {
	req := dto.CreateStation{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	// hash password
	stations := model.Station{}
	if err := copier.Copy(&stations, &req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	err := ctr.repo.Station.Create(c.Request.Context(), &stations)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *stationHandler) updateStation(c *gin.Context) {
	req := dto.UpdateStation{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	updateFields := model.UpdateFields{
		Field: "id",
		Value: req.ID,
		Data:  map[string]any{},
	}
	updateFields.Data["name"] = req.Name

	if err := ctr.repo.Station.Update(c.Request.Context(), &updateFields); err != nil {
		res := utils.GenerateServerError(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *stationHandler) deleteStation(c *gin.Context) {
	req := &dto.ReqByIDs{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	ids := utils.IdsIntToInCon(req.IDS)
	err := ctr.repo.Station.Delete(c.Request.Context(), ids)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}
