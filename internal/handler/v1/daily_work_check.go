package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/internal/service"
)

type DailyWorkCheckHandler struct {
	service service.DailyWorkCheckService
}

func NewDailyWorkCheckHandler(service service.DailyWorkCheckService) *DailyWorkCheckHandler {
	return &DailyWorkCheckHandler{service}
}

func (h *DailyWorkCheckHandler) MarkDone(c *gin.Context) {
	var input service.DailyWorkCheckInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.MarkDone(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Activity marked as done",
	})
}

// --- NEW ---
func (h *DailyWorkCheckHandler) GetActivitiesByWorker(c *gin.Context) {
	permitIDStr := c.Param("permit_id")
	nik := c.Param("nik")

	permitID, err := strconv.Atoi(permitIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permit_id"})
		return
	}

	activities, err := h.service.GetActivitiesByWorker(uint(permitID), nik)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}