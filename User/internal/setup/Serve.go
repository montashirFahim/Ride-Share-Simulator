package setup

import (
	"User/internal/config"
	"User/internal/redis"
	"User/internal/route"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

func RunServer(cmd *cobra.Command, args []string) error {
	config.LoadConfig()

	userRepo, err := ConnectDB(config.Cfg.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := redis.InitRedis(config.Cfg.RedisHost, config.Cfg.RedisPort); err != nil {
		log.Printf("Warning: Redis unavailable: %v", err)
	} else {
		log.Println("Redis cache enabled")
	}

	httpHandler := InitHandler(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	route.RegisterRoutes(r, config.Cfg.AuthUser, config.Cfg.AuthPass, httpHandler.RiderRegister, httpHandler.DriverRegister, httpHandler.Status, httpHandler.Info, httpHandler.DriverOnline)

	//return from
	StartServer(r, config.Cfg.ServerPort)
	return nil
}

func StartServer(router chi.Router, port string) {
	log.Println("Starting server at port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
