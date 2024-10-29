package app

import (
	"database/sql"
	"github.com/RicliZz/app_invest/config"
	authHandler "github.com/RicliZz/app_invest/internal/handlers"
	"github.com/RicliZz/app_invest/internal/repository/AuthRepository"
	"github.com/RicliZz/app_invest/internal/server"
	"github.com/RicliZz/app_invest/internal/services/AuthService"
	"github.com/gin-gonic/gin"
	"log"
)

func Run(configpath string) {
	router := gin.Default()

	var db *sql.DB

	authRepo := authRepository.NewAuthRepository(db)
	authServ := authService.NewAuthService(authRepo)
	authHand := authHandler.NewAuthHandler(authServ)

	authHand.RegisterRoutes(router)

	cfg := config.InitConfig(configpath)
	srv := server.NewAPIServer(cfg, router)
	log.Println("Server success running")
	srv.Start()
}
