package api

import (
	"dev-clash/internal/adapters/postgres"
	"dev-clash/internal/adapters/repositories"
	"dev-clash/internal/controllers"
	"dev-clash/internal/controllers/handlers"
	"dev-clash/internal/use-cases/user"
	"dev-clash/pkg/logger"
	"fmt"
	"log"
	"net/http"
)

type API struct {
	config *Config
}

func New(config *Config) *API {
	return &API{
		config: config,
	}
}

func (a *API) Start() error {

	// LOGGER INIT
	if err := logger.InitLogger(a.config.LoggerLevel); err!= nil {
		log.Print("Cannot read logger level")
		return err
	}
	logger.Info("Logger was configure successfully")

	// POSTGRES INIT
	storage, err := postgres.Init(a.config.PostgresURI)
	if err != nil {
		return fmt.Errorf("failed to connect db, %w", err)
	}
	logger.Info("Storage (postgres) was configure successfully")

	// BUSINESS LOGIC INIT
	userRepo := repositories.NewUserRepository(storage.GetDB())
	userService := user.New(userRepo)

	// ROUTER INIT
	handlers := handlers.New(userService)
	router := controllers.InitRoter(handlers)
	logger.Info("Router was configure successfully")


	// SERVER RUNNING
	logger.Info("API STARTED AT PORT", a.config.BindAddr)
	return http.ListenAndServe(a.config.BindAddr, router)
}