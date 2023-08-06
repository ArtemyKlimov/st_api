package handler

import (
	"net/http"
	"st_api/pkg/service"

	_ "st_api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(port string) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	v1 := router.Group("/api/v1")
	{
		regions := v1.Group("/regions")
		{
			regions.GET("", h.getAllRegions)
		}
		mineraltypes := v1.Group("/mineraltypes")
		{
			mineraltypes.GET("", h.getMineraltypes)
		}
		minerals := v1.Group("/minerals")
		{
			minerals.GET("", h.getMinerals)
		}
		deposits := v1.Group("/deposits")
		{
			deposits.GET("", h.getDeposits)
		}
		response := v1.Group("/response")
		{
			response.GET("", h.getResponse)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + port)
	//router.RunTLS(":8080", "public.pem", "private.key")
	return router
}
