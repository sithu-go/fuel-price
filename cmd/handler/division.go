package handler

import (
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/repository"
	"fuel-price/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type divisionHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newDivisionHandler(h *Handler) *divisionHandler {
	return &divisionHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *divisionHandler) register() {

	group := ctr.R.Group("/api/divisions")
	group.POST("", ctr.getDivisions)
	group.POST("/create", ctr.createDivision)
	group.POST("/update", ctr.updateDivision)
	group.POST("/delete", ctr.deleteDivision)
}

func (ctr *divisionHandler) getDivisions(c *gin.Context) {
	req := dto.SearchDivision{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	list, total, err := ctr.repo.Division.List(c.Request.Context(), &req)
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

func (ctr divisionHandler) createDivision(c *gin.Context) {
	req := dto.CreateDivision{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	// hash password
	divisions := model.Division{}
	if err := copier.Copy(&divisions, &req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	err := ctr.repo.Division.Create(c.Request.Context(), &divisions)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *divisionHandler) updateDivision(c *gin.Context) {
	req := dto.UpdateDivision{}
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

	if err := ctr.repo.Division.Update(c.Request.Context(), &updateFields); err != nil {
		res := utils.GenerateServerError(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *divisionHandler) deleteDivision(c *gin.Context) {
	req := &dto.ReqByIDs{}
	if err := c.ShouldBind(&req); err != nil {
		res := utils.GenerateValidationErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	ids := utils.IdsIntToInCon(req.IDS)
	err := ctr.repo.Division.Delete(c.Request.Context(), ids)
	if err != nil {
		res := utils.GenerateGormErrorResponse(err)
		c.JSON(res.HttpStatusCode, res)
		return
	}
	res := utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}
