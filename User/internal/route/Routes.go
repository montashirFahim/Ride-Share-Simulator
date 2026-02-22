package route

import (
	handler "User/internal/handler/api/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes(r chi.Router, authUser, authPass string,
	RiderRegisterHandler *handler.RiderRegisterHandler,
	DriverRegisterHandler *handler.DriverRegisterHandler,
	statusHandler *handler.StatusHandler,
	infoHandler *handler.InfoHandler,
	driverOnlineHandler *handler.DriverOnlineHandler) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Logger)
		// Public routes
		r.Post("/riders", RiderRegisterHandler.RegisterRider)
		r.Post("/drivers", DriverRegisterHandler.RegisterDriver)
		r.Put("/drivers/{id}/status", statusHandler.UpdateStatus)

		// Protected routes (with Basic Auth middleware)
		r.Group(func(r chi.Router) {
			r.Use(middleware.BasicAuth("Restricted", map[string]string{authUser: authPass}))

			r.Get("/users/{id}/info", infoHandler.GetUserInfo)

			r.Get("/drivers/online", driverOnlineHandler.ListOnlineDrivers)
		})
	})
}
