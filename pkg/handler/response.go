package handler

import (
	"net/http"
	"st_api/pkg/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type getResponseResponse struct {
	Data []models.Response `json:"data"`
}

// @Summary Получить расчет по мониторингу ММП
// @Tags response
// @Description Получить расчет по мониторингу ММП
// @ID get-response
// @Produce  json
// @Param    regionId    query     int  true  "ID of the region where minerals are extracted"  Format(int)
// @Param    mineralId      query     int  true  "ID of the mineral"  Format(int)
// @Success 200 {object} getResponseResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /response [get]
func (h *Handler) getResponse(ctx *gin.Context) {
	regionId := ctx.Request.URL.Query().Get("regionId")
	mineralId := ctx.Request.URL.Query().Get("mineralId")
	if regionId == "" || mineralId == "" {
		newErrorResponse(ctx, http.StatusBadRequest, "regionId and mineralId query params can not be empty")
		return
	}
	regIdint, err := strconv.Atoi(regionId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid regionId param")
		return
	}
	mineralIdint, err := strconv.Atoi(mineralId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid mineralId param")
		return
	}
	tempResp, err := h.services.Response.GetResponseByRegionForMineral(regIdint, mineralIdint)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var response []models.Response
	for _, r := range tempResp {
		var resp models.Response
		resp.T_trend = r.T_trend
		resp.G_temp = r.G_temp
		resp.Response = formatResponse(r.Response, r.Oopt, r.PermDistr, r.CoastDist, r.CoastRate, r.YedDist)
		resp.Oopt = r.Oopt
		resp.CoastDist = r.CoastDist
		resp.CoastRate = r.CoastRate
		resp.DepositName = r.DepositName
		resp.YedDist = r.YedDist
		resp.PermDistr = r.PermDistr
		response = append(response, resp)
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func formatResponse(response, oopt, permDistr string, coastDist, coastRate, yedDist float32) string {
	resultString := response
	if coastDist >= 1 && coastRate >= 2 {
		resultString += ". Месторождение расположено в зоне с активно отступающими морскими берегами"
	}
	if yedDist <= 10 {
		resultString += ". Месторождение расположено в пределах распространения сильнольдистых мерзлых отложений"
	}
	if oopt != "" {
		resultString += ". Месторождение расположено в пределах ООПТ " + oopt + "."
	}

	resultString += getPermDistrValue(permDistr)
	return resultString

}

func getPermDistrValue(in string) string {
	switch strings.ToLower(in) {
	case "cont":
		return "Сплошное распространение (более 90% площади)"
	case "discon":
		return "Прерывистое распространение (от 50 до 90% площади)"
	case "spora":
		return "Прерывистое распространение (от 10 до 50% площади)"
	case "isol":
		return "Островное распространение (менее 10% площади)"
	default:
		return ""
	}
}
