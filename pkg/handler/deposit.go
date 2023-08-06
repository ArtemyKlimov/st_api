package handler

import (
	"net/http"
	"st_api/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getDepositsResponse struct {
	Data []models.Deposit `json:"data"`
}

// @Summary Получить список месторождений в заданном регионе
// @Tags deposits
// @Description Получить список месторождений в заданном регионе
// @ID get-deposits
// @Produce  json
// @Param    regionId    query     int  true  "Id региона"  Format(int)
// @Success 200 {object} getDepositsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /deposits [get]
func (h *Handler) getDeposits(ctx *gin.Context) {
	regionId := ctx.Request.URL.Query().Get("regionId")
	if regionId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "regionId query param is empty")
		return
	}
	regIdint, err := strconv.Atoi(regionId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid regionId param")
		return
	}
	deposits, err := h.services.Deposits.GetInRegion(regIdint)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, deposits)
	return
}
