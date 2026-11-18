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

func (h *healthCheck) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.CheckHealth())
}
