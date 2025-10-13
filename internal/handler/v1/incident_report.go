package v1

import (
	"net/http"
	"path/filepath"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/internal/service"
)

type IncidentReportHandler struct {
	service *service.IncidentReportService
}

func NewIncidentReportHandler(s *service.IncidentReportService) *IncidentReportHandler {
	return &IncidentReportHandler{service: s}
}

func (h *IncidentReportHandler) Create(c *gin.Context) {
	permitIDStr := c.PostForm("permit_id")
	description := c.PostForm("description")
	dateStr := c.PostForm("date")

	// parse permit_id
	var permitID uint
	_, err := fmt.Sscan(permitIDStr, &permitID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permit_id"})
		return
	}

	// parse date (support multiple formats)
	var date *time.Time
	if dateStr != "" {
		// coba format RFC3339 (ex: 2025-10-09T09:00:00Z)
		t, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			// fallback ke yyyy-mm-dd
			t2, err2 := time.Parse("2006-01-02", dateStr)
			if err2 == nil {
				date = &t2
			}
		} else {
			date = &t
		}
	}

	// upload photo
	file, err := c.FormFile("photo")
	var filename string
	if err == nil && file != nil {
		filename = filepath.Base(file.Filename)
		savePath := filepath.Join("uploads", "incident_report", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload photo"})
			return
		}
	}

	// call service
	err = h.service.Create(permitID, description, filename, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "incident report created successfully"})
}
