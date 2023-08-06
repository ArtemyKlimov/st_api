package handler

import (
	"net/http"
	"st_api/pkg/models"

	"github.com/gin-gonic/gin"
)

type getAllRegionsResponse struct {
	Data []models.Region `json:"data"`
}

// @Summary Получить список наблюдаемых регионов
// @Tags regions
// @Description Получить список наблюдаемых регионов
// @ID get-all-regions
// @Produce  json
// @Success 200 {object} getAllRegionsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /regions [get]
func (h *Handler) getAllRegions(c *gin.Context) {

	regions, err := h.services.Regions.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, regions)

}
