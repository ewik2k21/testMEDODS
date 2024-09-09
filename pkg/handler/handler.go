package handler

import (
	"testMEDODS/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth", h.tokenValidation)
	{
		auth.POST("/add-user", h.addUser)
		auth.POST("getToken", h.getPairTokens)
		auth.POST("/refresh-pair", h.refreshPairTokens)
	}

	return router
}
