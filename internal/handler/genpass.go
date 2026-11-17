package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoand3012/bookmark-management/internal/service"
)

const DefaultLength = 10

type GenPass interface {
	GeneratePassword(length int) (string, error)
}

type genPass struct {
	svc service.GenPass
}

func NewGenPass(svc service.GenPass) *genPass {
	return &genPass{svc: svc}
}

func (g *genPass) GeneratePassword(c *gin.Context) {
	pass, err := g.svc.GeneratePassword(DefaultLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"password": pass})
}
