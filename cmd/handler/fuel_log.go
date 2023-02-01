package handler

import (
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/repository"
	"fuel-price/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type fuel_logHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newfuel_logHandler(h *Handler) *fuel_logHandler {
	return &fuel_logHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *fuel_logHandler) register() {

	group := ctr.R.Group("/api/fuel_logs")
	group.POST("", ctr.getFuelLog)
	group.POST("/create", ctr.createFuelLog)
	group.POST("/update", ctr.updateFuelLog)
	group.POST("/delete", ctr.deleteFuelLog)
}

func (ctr *fuel_logHandler) getFuelLog(c *gin.Context) {
	req := dto.SearchFuelLog{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	list, total, err := ctr.repo.FuelLog.List(c.Request.Context(), &req)
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

func (ctr fuel_logHandler) createFuelLog(c *gin.Context) {
	req := dto.CreateFuelLog{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	// hash password
	fuel_log := model.FuelLog{}
	if err := copier.Copy(&fuel_log, &req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	err := ctr.repo.FuelLog.Create(c.Request.Context(), &fuel_log)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *fuel_logHandler) updateFuelLog(c *gin.Context) {
	req := dto.UpdateFuelLog{}
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

	if err := ctr.repo.FuelLog.Update(c.Request.Context(), &updateFields); err != nil {
		res := utils.GenerateServerError(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *fuel_logHandler) deleteFuelLog(c *gin.Context) {
	req := &dto.ReqByIDs{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	ids := utils.IdsIntToInCon(req.IDS)
	err := ctr.repo.FuelLog.Delete(c.Request.Context(), ids)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}
