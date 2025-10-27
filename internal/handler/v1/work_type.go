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

type WorkTypeHandler struct {
	service service.WorkTypeService
}

func NewWorkTypeHandler(db *gorm.DB, _ *setting.Configuration) *WorkTypeHandler {
	repo := repository.NewWorkTypeRepository(db)
	svc := service.NewWorkTypeService(repo)
	return &WorkTypeHandler{service: svc}
}

func (h *WorkTypeHandler) Create(c *gin.Context) {
	var workType domain.WorkType
	if err := c.ShouldBindJSON(&workType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&workType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "WorkType created", "data": workType})
}

func (h *WorkTypeHandler) List(c *gin.Context) {
	workTypes, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": workTypes})
}

func (h *WorkTypeHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	workType, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "WorkType not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": workType})
}

func (h *WorkTypeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var workType domain.WorkType
	if err := c.ShouldBindJSON(&workType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workType.ID = uint(id)
	if err := h.service.Update(&workType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "WorkType updated", "data": workType})
}

func (h *WorkTypeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "WorkType deleted"})
}