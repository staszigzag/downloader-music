package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/staszigzag/downloader-music/internal/delivery/http/v1"
	"github.com/staszigzag/downloader-music/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		// gin.Logger(),
		corsMiddleware,
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Init router
	h.initAPIRoutes(router)
	return router
}

func (h *Handler) initAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// Implementation of handler versioning
	handlerV1 := v1.NewHandler(h.services)
	handlerV1.Init(api)
}
