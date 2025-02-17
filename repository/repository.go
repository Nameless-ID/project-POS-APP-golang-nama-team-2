package repository

import (
	accessrepository "project_pos_app/repository/access_repository"
	authrepository "project_pos_app/repository/auth_repository"
	categoryrepository "project_pos_app/repository/category_repository"
	dashboardrepository "project_pos_app/repository/dashboard_repository"
	"project_pos_app/repository/notification"
	orderrepository "project_pos_app/repository/order_repository"
	productrepository "project_pos_app/repository/product"
	profilesuperadmin "project_pos_app/repository/profile_superadmin"
	reservationrepository "project_pos_app/repository/reservation_repository"
	revenuerepository "project_pos_app/repository/revenue_repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AllRepository struct {
	Category    categoryrepository.CategoryRepository
	Access      accessrepository.AccessRepository
	Auth        authrepository.AuthRepoInterface
	Notif       notification.NotifRepoInterface
	Revenue     revenuerepository.RevenueRepositoryInterface
	Product     productrepository.ProductRepo
	Order       orderrepository.OrderRepository
	Superadmin  profilesuperadmin.SuperadminRepo
	Reservation reservationrepository.RepositoryReservation
	Dashboard   dashboardrepository.RepositoryDashboard
}

func NewAllRepo(DB *gorm.DB, Log *zap.Logger) *AllRepository {
	return &AllRepository{
		Auth:        authrepository.NewManagementVoucherRepo(DB, Log),
		Notif:       notification.NewNotifRepo(DB, Log),
		Revenue:     revenuerepository.NewRevenueRepository(DB, Log),
		Product:     productrepository.NewProductRepo(DB, Log),
		Order:       orderrepository.NewOrderRepo(DB, Log),
		Superadmin:  profilesuperadmin.NewSuperadmin(DB, Log),
		Category:    categoryrepository.NewCategoryRepo(DB, Log),
		Access:      accessrepository.NewAccessRepository(DB, Log),
		Reservation: reservationrepository.NewReservationRepository(DB, Log),
		Dashboard:   dashboardrepository.NewReservationRepository(DB, Log),
	}
}
