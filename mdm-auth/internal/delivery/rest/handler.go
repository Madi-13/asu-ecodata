package rest

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin/binding"
	_ "github.com/go-playground/validator/v10"
	"mdm-auth/config"
	v1 "mdm-auth/internal/delivery/rest/v1"
	"mdm-auth/internal/services"
	"net/http"
	_ "net/http"
	_ "time"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("validFavRestsSelectionTypes", domain.ValidFavRestsSelectionTypes)
	//}

	router.Use(
		gzip.Gzip(gzip.DefaultCompression),
	)

	h.initAPI(router)

	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
