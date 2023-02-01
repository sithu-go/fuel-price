package handler

import (
	"fuel-price/pkg/dto"
	"fuel-price/pkg/repository"
	"fuel-price/pkg/utils"

	"github.com/gin-gonic/gin"
)

type fuelHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newFuelHandler(h *Handler) *fuelHandler {

	return &fuelHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *fuelHandler) register() {
	group := ctr.R.Group("/api/fuel")

	group.GET("/prices", ctr.getFuelPirces)
}

func (ctr *fuelHandler) getFuelPirces(c *gin.Context) {
	req := dto.FuelPriceFilter{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}

	prices, total, err := ctr.repo.Fuel.GetFuelPrices(&req)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}

	res := utils.GenerateSuccessResponse(gin.H{
		"list":  prices,
		"total": total,
	})

	c.JSON(res.HttpStatusCode, res)
}
