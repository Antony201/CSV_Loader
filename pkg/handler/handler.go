package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Antony201/CsvLoader/pkg/service"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/Antony201/CsvLoader/docs"

)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger endpoint

	v1 := router.Group("/api/v1")
	{
		transactions := v1.Group("/transactions")
		{
			transactions.POST("/upload", h.uploadTransactions)
			transactions.GET("/", h.getTransactions)
		}
	}

	return router
}