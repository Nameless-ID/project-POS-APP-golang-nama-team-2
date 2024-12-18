package controller

import (
	authcontroller "project_pos_app/controller/auth_controller"
	examplecontroller "project_pos_app/controller/example_controller"
	notifcontroller "project_pos_app/controller/notif_controller"
	"project_pos_app/database"
	"project_pos_app/service"

	"go.uber.org/zap"
)

// AllController mengelola semua controller dalam aplikasi
type AllController struct {
	Example examplecontroller.ExampleController
	Auth    authcontroller.AuthHandler
	Notif   notifcontroller.NotifController
}

// NewAllController membuat instance AllController dengan dependensi yang diperlukan
func NewAllController(service *service.AllService, log *zap.Logger, cache *database.Cache) *AllController {
	return &AllController{
		Example: examplecontroller.NewExampleController(service, log),
		Auth:    authcontroller.NewAuthHandler(service, log, cache), // Menggunakan nama yang konsisten
		Notif:   notifcontroller.NewNotifController(service, log),
	}
}
