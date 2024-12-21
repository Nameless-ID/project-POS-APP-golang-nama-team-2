package controller

import (
	authcontroller "project_pos_app/controller/auth_controller"
	examplecontroller "project_pos_app/controller/example_controller"
	notifcontroller "project_pos_app/controller/notif_controller"
	ordercontroller "project_pos_app/controller/order_controller"
	productcontroller "project_pos_app/controller/product_controller"
	revenuecontroller "project_pos_app/controller/revenue_controller"
	"project_pos_app/database"
	"project_pos_app/service"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AllController mengelola semua controller dalam aplikasi
type AllController struct {
	Example     examplecontroller.ExampleController
	Auth        authcontroller.AuthHandler
	Notif       notifcontroller.NotifController
	Revenue     revenuecontroller.RevenueController
	Product     productcontroller.ProductController
	Order       ordercontroller.OrderController
	RedisClient *redis.Client
	DB          *gorm.DB
}

// NewAllController membuat instance AllController dengan dependensi yang diperlukan
func NewAllController(service *service.AllService, log *zap.Logger, cache *database.Cache, redisClient *redis.Client, db *gorm.DB) *AllController {
	return &AllController{
		Example: examplecontroller.NewExampleController(service, log),
		Auth:    *authcontroller.NewAuthHandler(service, log, cache, redisClient, db), // Perbaikan pada pemanggilan konstruktor
		Notif:   notifcontroller.NewNotifController(service, log),
		Revenue: revenuecontroller.NewRevenueController(service, log),
		Product: *productcontroller.NewProductController(service, log), // Perbaikan jika diperlukan
		Order:   ordercontroller.NewOrderController(service, log),
	}
}
