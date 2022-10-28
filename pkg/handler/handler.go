package handler

import (
	"github.com/gin-gonic/gin"
	"test_task/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		transactions := api.Group("/transactions")
		{
			transactions.POST("/upload", h.uploadTransactions)
			transactions.GET("/", h.getAllTransactions)
			transactions.GET("/:id", h.getTransactionById)
		}
	}

	return router
}