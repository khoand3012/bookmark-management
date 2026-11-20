package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoand3012/bookmark-management/internal/service"
)

type healthCheck struct {
	svc service.HealthCheck
}

func NewHealthCheck(svc service.HealthCheck) *healthCheck {
	return &healthCheck{svc: svc}
}

// CheckHealth godoc
// @Summary      Health check
// @Description  Reports that the service is up, along with its service name and instance ID.
// @Tags         health
// @Produce      json
// @Success      200  {object}  service.HealthCheckResult
// @Router       /health-check [get]
func (h *healthCheck) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.CheckHealth())
}
