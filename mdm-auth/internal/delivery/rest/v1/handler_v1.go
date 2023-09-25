package v1

import (
	"github.com/gin-gonic/gin"
	"mdm-auth/internal/services"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.InitAuthRoutes(v1)
	}
}
