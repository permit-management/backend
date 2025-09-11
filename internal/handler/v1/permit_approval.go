package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/service"
)

type PermitApprovalHandler struct {
	service service.PermitApprovalService
}

func NewPermitApprovalHandler(s service.PermitApprovalService) *PermitApprovalHandler {
	return &PermitApprovalHandler{service: s}
}

func (h *PermitApprovalHandler) ApprovePermit(c *gin.Context) {
	var approval domain.PermitApproval
	if err := c.ShouldBindJSON(&approval); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.ApprovePermit(&approval); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Permit approval processed successfully"})
}
