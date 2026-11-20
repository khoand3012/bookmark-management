package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khoand3012/bookmark-management/internal/service"
)

const DefaultLength = 10

type genPass struct {
	svc service.GenPass
}

func NewGenPass(svc service.GenPass) *genPass {
	return &genPass{svc: svc}
}

// GenPassResponse is the success body of GET /genpass.
type GenPassResponse struct {
	Password string `json:"password" example:"aB3dE9xYz1"`
}

// ErrorResponse is the error body returned by the API.
type ErrorResponse struct {
	Error string `json:"error" example:"Failed to generate password"`
}

// GeneratePassword godoc
// @Summary      Generate a random password
// @Description  Generates a random 10-character password made of letters (A-Z, a-z) and digits (0-9), using a cryptographically secure random source.
// @Tags         password
// @Produce      json
// @Success      200  {object}  GenPassResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /genpass [get]
func (g *genPass) GeneratePassword(c *gin.Context) {
	pass, err := g.svc.GeneratePassword(DefaultLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"password": pass})
}
