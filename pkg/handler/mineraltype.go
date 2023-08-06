package handler

import (
	"net/http"
	"st_api/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getMineraltypesResponse struct {
	Data []models.MineralTypes `json:"data"`
}

// @Summary Получить список типов полезных ископаемых
// @Tags mineral types
// @Description Получить список типов полезных ископаемых
// @ID get-mineral-types
// @Produce  json
// @Param    regionId    query     string  false  "ID региона"  Format(int)
// @Success 200 {object} getMineraltypesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mineraltypes [get]
func (h *Handler) getMineraltypes(ctx *gin.Context) {
	var mineraltypes []models.MineralTypes
	regionId := ctx.Request.URL.Query().Get("regionId")
	if regionId == "" {
		mineraltypes, err := h.services.MineralTypes.GetAll()
		if err != nil {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, mineraltypes)
		return
	}
	id, err := strconv.Atoi(regionId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid regionId param")
		return
	}
	mineraltypes, err = h.services.MineralTypes.GetAllInRegion(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, mineraltypes)

}
