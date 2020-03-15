package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"advance-go/internal/config"
)

type Handler struct {
	conf *config.Config
}

func NewHandler(conf *config.Config) *Handler {
	return &Handler{conf: conf}
}

func (h *Handler) GetProjectCode(c *gin.Context) {
	if h.conf.Maintenance {
		c.JSON(http.StatusForbidden, gin.H{"maintenance": true})
		return
	}
	c.JSON(http.StatusOK, gin.H{"project_code": h.conf.ProjectCode})
}
