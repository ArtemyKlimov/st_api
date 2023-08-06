package handler

import (
	"net/http"
	"st_api/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getMineralsResponse struct {
	Data []models.Minerals `json:"data"`
}

// @Summary Получить список полезных ископаемых указанного типа в заданном регионе
// @Tags minerals
// @Description Получить список полезных ископаемых указанного типа в заданном регионе
// @ID get-minerals
// @Produce  json
// @Param    regionId    query     int  false  "ID региона"  Format(int)
// @Param    typeId      query     int  false  "ID типа полезных ископаемых"  Format(int)
// @Success 200 {object} getMineralsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /minerals [get]
func (h *Handler) getMinerals(ctx *gin.Context) {
	regionId := ctx.Request.URL.Query().Get("regionId")
	typeId := ctx.Request.URL.Query().Get("typeId")
	if regionId == "" && typeId == "" {
		minerals, err := h.services.Minerals.GetAll()
		if err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, minerals)
		return
	}
	if regionId != "" && typeId == "" {
		id, err := strconv.Atoi(regionId)
		if err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, "invalid regionId param")
			return
		}
		minerals, err := h.services.Minerals.GetAllInRegion(id)
		if err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, minerals)
		return
	}
	if regionId == "" && typeId != "" {
		id, err := strconv.Atoi(typeId)
		if err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, "invalid typeId param")
			return
		}
		minerals, err := h.services.Minerals.GetByType(id)
		if err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, minerals)
		return
	}
	regIdint, err := strconv.Atoi(regionId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid regionId param")
		return
	}
	typeIdint, err := strconv.Atoi(typeId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid typeId param")
		return
	}
	minerals, err := h.services.Minerals.GetAllInRegionByType(regIdint, typeIdint)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, minerals)
	return

}
