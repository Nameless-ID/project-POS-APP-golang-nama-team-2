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
