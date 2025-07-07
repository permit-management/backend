package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/permit-management/backend/internal/domain"
	"github.com/permit-management/backend/internal/repository"
	"github.com/permit-management/backend/internal/service"
	"github.com/permit-management/backend/pkg/setting"
	"gorm.io/gorm"
)

type DepartementHandler struct {
	service service.DepartementService
}

func NewDepartementHandler(db *gorm.DB, _ *setting.Configuration) *DepartementHandler {
	repo := repository.NewDepartementRepository(db)
	svc := service.NewDepartementService(repo)
	return &DepartementHandler{service: svc}
}

func (h *DepartementHandler) Create(c *gin.Context) {
	var dept domain.Departement
	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&dept); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dept)
}

func (h *DepartementHandler) List(c *gin.Context) {
	depts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, depts)
}

func (h *DepartementHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	dept, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, dept)
}

func (h *DepartementHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var dept domain.Departement
	if err := c.ShouldBindJSON(&dept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dept.ID = uint(id)

	// 🔥 gunakan update khusus tanpa overwrite created_at
	if err := h.service.UpdateWithoutCreatedAt(&dept); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dept)
}

func (h *DepartementHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":      1,
			"message":   err.Error(),
			"data":      struct{}{},
			"timestamp": time.Now(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"message":   "Success",
		"data":      struct{}{},
		"timestamp": time.Now(),
	})
}
