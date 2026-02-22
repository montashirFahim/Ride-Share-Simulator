package setup

import (
	handler "User/internal/handler/api/v1"
	"User/internal/repository"
)

type Handlers struct {
	RiderRegister  *handler.RiderRegisterHandler
	DriverRegister *handler.DriverRegisterHandler
	Status         *handler.StatusHandler
	Info           *handler.InfoHandler
	DriverOnline   *handler.DriverOnlineHandler
}

func InitHandler(userRepo repository.UserRepository) *Handlers {
	userQuery := repository.NewUserQuery(userRepo)

	return &Handlers{
		RiderRegister:  handler.NewRiderRegisterHandler(userRepo),
		DriverRegister: handler.NewDriverRegisterHandler(userRepo),
		Status:         handler.NewStatusHandler(userRepo),
		Info:           handler.NewInfoHandler(userQuery),
		DriverOnline:   handler.NewDriverOnlineHandler(userRepo),
	}
}
