package handler

import (
	"fuel-price/pkg/ds"
	"fuel-price/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	R    *gin.Engine
	repo *repository.Repository
}

type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {
	return &Handler{
		R: c.R,
		repo: repository.NewRepository(&repository.RepoConfig{
			DS: c.DS,
		}),
	}
}

func (h *Handler) Register() {
	// Dashboard
	// dashboardHandler := NewDashboardHandler(h)
	// dashboardHandler.register()

}
