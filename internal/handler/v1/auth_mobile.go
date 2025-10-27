package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/internal/service"
)

type AuthMobileHandler struct {
	authMobileService service.AuthMobileService
}

func NewAuthMobileHandler(authMobileService service.AuthMobileService) *AuthMobileHandler {
	return &AuthMobileHandler{authMobileService}
}

func (h *AuthMobileHandler) Login(c *gin.Context) {
	var req struct {
		PermitNumber string `json:"permit_number"`
		NIK          string `json:"nik"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, worker, err := h.authMobileService.Login(req.PermitNumber, req.NIK)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"worker": worker,
	})
}
