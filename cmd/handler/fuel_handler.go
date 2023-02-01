package handler

import (
	"fuel-price/pkg/repository"

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

}
