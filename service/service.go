package service

import (
	"project_pos_app/repository"
	authservice "project_pos_app/service/auth_service"
	exampleservice "project_pos_app/service/example_service"
	notifservice "project_pos_app/service/notif_service"
	orderservice "project_pos_app/service/order_service"
	productservice "project_pos_app/service/product_service"
	reservationservice "project_pos_app/service/reservation_service"
	revenueservice "project_pos_app/service/revenue_service"

	"go.uber.org/zap"
)

type AllService struct {
	Example     exampleservice.ExampleService
	Auth        authservice.AuthService
	Notif       notifservice.NotifServiceInterface
	Revenue     revenueservice.RevenueServiceInterface
	Product     productservice.ProductService
	Order       orderservice.OrderService
	Reservation reservationservice.ServiceReservation
}

func NewAllService(repo *repository.AllRepository, log *zap.Logger) *AllService {
	return &AllService{
		Example:     exampleservice.NewExampleService(repo, log),
		Auth:        authservice.NewManagementVoucherService(repo, log),
		Notif:       notifservice.NewNotifService(repo, log),
		Revenue:     revenueservice.NewRevenueService(repo, log),
		Product:     productservice.NewProductService(repo, log),
		Order:       orderservice.NewOrderService(repo, log),
		Reservation: reservationservice.NewRevenueService(repo, log),
	}
}
