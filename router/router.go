package router

import (
	"project_pos_app/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx *infra.IntegrationContext) *gin.Engine {
	r := gin.Default()

	// Swagger Documentation Route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Authentication Routes
	AuthRoutes(r, ctx)

	// Notification Routes
	NotificationRoutes(r, ctx)
	return r
}

func AuthRoutes(r *gin.Engine, ctx *infra.IntegrationContext) {
	authRoute := r.Group("/api/auth")
	{

		authRoute.POST("/login", ctx.Ctl.Auth.Login)                  // Login
		authRoute.GET("/check-email", ctx.Ctl.Auth.CheckEmail)        // Check Email
		authRoute.GET("/validate-otp", ctx.Ctl.Auth.ValidateOTP)      // Validate OTP
		authRoute.POST("/reset-password", ctx.Ctl.Auth.ResetPassword) // Reset Password
	}
}

func StaffRoutes(r *gin.Engine, ctx *infra.IntegrationContext) {
	staffController := controller.NewStaffController(staffService, logger)

	// Setup routes
	router := gin.Default()
	router.GET("/staffs", staffController.GetAllStaff)
	router.GET("/staffs/:id", staffController.GetStaffByID)
	router.POST("/staffs", staffController.CreateStaff)
	router.PUT("/staffs/:id", staffController.UpdateStaff)
	router.DELETE("/staffs/:id", staffController.DeleteStaff)
}
func NotificationRoutes(r *gin.Engine, ctx *infra.IntegrationContext) {
	notifRoute := r.Group("/api/notifications")
	{
		notifRoute.POST("", ctx.Ctl.Notif.CreateNotifications)                     // Create Notification
		notifRoute.GET("", ctx.Ctl.Notif.GetAllNotifications)                      // Get All Notifications
		notifRoute.GET("/:id", ctx.Ctl.Notif.GetNotificationByID)                  // Get Notification by ID
		notifRoute.PUT("/:id", ctx.Ctl.Notif.UpdateNotification)                   // Update Notification
		notifRoute.DELETE("/:id", ctx.Ctl.Notif.DeleteNotification)                // Delete Notification
		notifRoute.PUT("/mark-all-read", ctx.Ctl.Notif.MarkAllNotificationsAsRead) // Mark All as Read
	}
}

func RevenueRoutes(r *gin.Engine, ctx *infra.IntegrationContext) {
	revenueRoute := r.Group("/api")
	{
		revenueRoute.GET("/revenue/month", ctx.Ctl.Revenue.GetMonthlyRevenue)
		revenueRoute.GET("/revenue/products", ctx.Ctl.Revenue.GetProductRevenues)
		revenueRoute.GET("/revenue/status", ctx.Ctl.Revenue.GetTotalRevenueByStatus)
	}
}

func ProductRoutes(r *gin.Engine, ctx *infra.IntegrationContext) {
	productRoute := r.Group("/api")
	{
		productRoute.GET("/products", ctx.Ctl.Product.GetAllProducts)
		productRoute.GET("/products/:id", ctx.Ctl.Product.GetProductByID)
		productRoute.POST("/products", ctx.Ctl.Product.CreateProduct)
		productRoute.PUT("/products/:id", ctx.Ctl.Product.UpdateProduct)
		productRoute.DELETE("/product/:id", ctx.Ctl.Product.DeleteProduct)
	}
}
