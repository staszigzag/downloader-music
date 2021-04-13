package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initDownloaderRoutes(api *gin.RouterGroup) {
	downloader := api.Group("/downloader")
	{
		downloader.GET("/ping", h.ping)
	}
}

func (h *Handler) ping(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{"pong"})
}
