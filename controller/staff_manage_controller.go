package controller

import (
	"net/http"
	"project_pos_app/model"
	"project_pos_app/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Staff = model.Staff

type StaffController struct {
	Service service.StaffService
	Log     *zap.Logger
}

// NewStaffController creates a new instance of StaffController
func NewStaffController(service service.StaffService, log *zap.Logger) *StaffController {
	return &StaffController{
		Service: service,
		Log:     log,
	}
}

// GetAllStaff handles GET /staffs
func (c *StaffController) GetAllStaff(ctx *gin.Context) {
	staffs, err := c.Service.GetAll()
	if err != nil {
		c.Log.Error("Failed to retrieve staff list", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve staff list"})
		return
	}
	ctx.JSON(http.StatusOK, staffs)
}

// GetStaffByID handles GET /staffs/:id
func (c *StaffController) GetStaffByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Log.Warn("Invalid staff ID", zap.String("id", ctx.Param("id")), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	staff, err := c.Service.GetByID(id)
	if err != nil {
		c.Log.Warn("Staff not found", zap.Int("ID", id), zap.Error(err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		return
	}

	ctx.JSON(http.StatusOK, staff)
}

// CreateStaff handles POST /staffs
func (c *StaffController) CreateStaff(ctx *gin.Context) {
	var staff Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		c.Log.Warn("Invalid request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.Service.Create(&staff)
	if err != nil {
		c.Log.Error("Failed to create staff", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create staff"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Staff created successfully", "staff": staff})
}

// UpdateStaff handles PUT /staffs/:id
func (c *StaffController) UpdateStaff(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Log.Warn("Invalid staff ID", zap.String("id", ctx.Param("id")), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	var staff Staff
	if err := ctx.ShouldBindJSON(&staff); err != nil {
		c.Log.Warn("Invalid request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	staff.ID = id
	err = c.Service.Update(&staff)
	if err != nil {
		c.Log.Error("Failed to update staff", zap.Int("ID", id), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update staff"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Staff updated successfully", "staff": staff})
}

// DeleteStaff handles DELETE /staffs/:id
func (c *StaffController) DeleteStaff(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.Log.Warn("Invalid staff ID", zap.String("id", ctx.Param("id")), zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid staff ID"})
		return
	}

	err = c.Service.Delete(id)
	if err != nil {
		c.Log.Error("Failed to delete staff", zap.Int("ID", id), zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete staff"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Staff deleted successfully"})
}
