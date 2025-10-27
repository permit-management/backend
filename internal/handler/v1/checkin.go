package v1

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/internal/service"
	"github.com/permit-management/backend/pkg/logger"
)

type CheckInHandler struct {
	service service.CheckInService
}

func NewCheckInHandler(s service.CheckInService) *CheckInHandler {
	return &CheckInHandler{service: s}
}

func (h *CheckInHandler) CheckIn(c *gin.Context) {
	permitID 	:= c.PostForm("permit_id")
	workerName 	:= c.PostForm("worker_name")
	Nik 		:= c.PostForm("nik")

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "photo is required"})
		return
	}

	filename := time.Now().Format("20060102_150405") + filepath.Ext(file.Filename)
	savePath := "uploads/checkin/" + filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		logger.Log().Error("failed to save uploaded file: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save photo"})
		return
	}

	photoURL := "uploads/checkin/" + filename

	checkIn, err := h.service.CreateCheckIn(parseUint(permitID), workerName, Nik, photoURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   checkIn,
	})
}

func parseUint(s string) uint {
	var id uint
	fmt.Sscanf(s, "%d", &id)
	return id
}
