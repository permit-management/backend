package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
	"github.com/permit-management/backend/internal/service"
	"github.com/permit-management/backend/pkg/setting"
	"gorm.io/gorm"
)

type PermitHandler struct {
	permitService service.PermitService
}

func NewPermitHandler(db *gorm.DB, _ *setting.Configuration) *PermitHandler {
	repo := repository.NewPermitRepository(db)
	svc := service.NewPermitService(repo)
	return &PermitHandler{permitService: svc}
}

func (h *PermitHandler) CreatePermit(c *gin.Context) {
	var permit domain.Permit
	if err := c.ShouldBindJSON(&permit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.permitService.CreatePermit(&permit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Permit created", "data": permit})
}

func (h *PermitHandler) GetAllPermits(c *gin.Context) {
	permits, err := h.permitService.GetAllPermits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": permits})
}

func (h *PermitHandler) GetPermitByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	permit, err := h.permitService.GetPermitByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permit not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": permit})
}

func (h *PermitHandler) UpdatePermit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var permit domain.Permit
	if err := c.ShouldBindJSON(&permit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	permit.ID = uint(id)
	if err := h.permitService.UpdatePermit(&permit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Permit updated", "data": permit})
}

func (h *PermitHandler) DeletePermit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.permitService.DeletePermit(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Permit deleted"})
}
