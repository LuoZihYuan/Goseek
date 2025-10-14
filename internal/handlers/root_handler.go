package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

// GetHealth handles GET /health
// @Summary Get health condition
// @Description Get health of server
// @ID GetHealth
// @Tags (Root)
// @Produce json
// @Success 200
// @Failure 503
// @Router /health [get]
func (h *RootHandler) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
