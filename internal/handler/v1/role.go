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

type RoleHandler struct {
	service service.RoleService
}

// NewRoleHandler inisialisasi RoleHandler
func NewRoleHandler(db *gorm.DB, _ *setting.Configuration) *RoleHandler {
	repo := service.NewRoleService(repository.NewRoleRepository(db))
	return &RoleHandler{service: repo}
}

// Create membuat role baru
func (h *RoleHandler) Create(c *gin.Context) {
	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    0,
		"message": "Role created successfully",
		"data":    role,
	})
}

// List mengambil semua role
func (h *RoleHandler) List(c *gin.Context) {
	roles, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

// Get mengambil role berdasarkan ID
func (h *RoleHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// Update memperbarui role berdasarkan ID
func (h *RoleHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	role.ID = uint(id)

	if err := h.service.Update(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"message":   "Role updated successfully",
		"data":      role,
		"timestamp": time.Now(),
	})
}

// Delete menghapus role berdasarkan ID
func (h *RoleHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"message":   "Role deleted successfully",
		"timestamp": time.Now(),
	})
}