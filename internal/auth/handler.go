package auth

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService Service
}

func NewHandler(authService Service) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) RegisterUser(c *gin.Context) {

}

func (h *Handler) LoginUser(c *gin.Context) {

}
