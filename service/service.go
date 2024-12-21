package service

import (
	"project_pos_app/repository"
	authservice "project_pos_app/service/auth_service"
	exampleservice "project_pos_app/service/example_service"
	notifservice "project_pos_app/service/notif_service"
	orderservice "project_pos_app/service/order_service"
	productservice "project_pos_app/service/product_service"
	revenueservice "project_pos_app/service/revenue_service"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AllService struct {
	DB      *gorm.DB
	Example exampleservice.ExampleService
	Auth    *authservice.AuthService
	Notif   notifservice.NotifServiceInterface
	Revenue revenueservice.RevenueServiceInterface
	Product productservice.ProductService
	Order   orderservice.OrderService
}

func NewAllService(repo *repository.AllRepository, redis *redis.Client, log *zap.Logger) *AllService {
	return &AllService{
		Example: exampleservice.NewExampleService(repo, log),
		Auth:    authservice.NewAuthService(&repo.Auth, log, redis, repo.DB),
		Notif:   notifservice.NewNotifService(repo, log),
		Revenue: revenueservice.NewRevenueService(repo, log),
		Product: productservice.NewProductService(repo, log),
		Order:   orderservice.NewOrderService(repo, log),
	}
}
