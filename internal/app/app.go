package app

import (
	"fmt"
	"github.com/RicliZz/app_invest/config"
	"github.com/RicliZz/app_invest/internal/handlers/authHandler"
	"github.com/RicliZz/app_invest/internal/handlers/profileHandler"
	"github.com/RicliZz/app_invest/internal/handlers/startUpHandler"
	"github.com/RicliZz/app_invest/internal/repository/AuthRepository"
	"github.com/RicliZz/app_invest/internal/repository/StartUpRepository"
	"github.com/RicliZz/app_invest/internal/repository/UserDetailsRepository"
	"github.com/RicliZz/app_invest/internal/repository/UserRepository"
	"github.com/RicliZz/app_invest/internal/server"
	"github.com/RicliZz/app_invest/internal/services/AuthService"
	profileService "github.com/RicliZz/app_invest/internal/services/ProfileService"
	"github.com/RicliZz/app_invest/internal/services/StartUpService"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Run(configpath string) {
	router := gin.Default()
	cfg := config.InitConfig(configpath)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.Username, os.Getenv("DB_PASSWORD"), cfg.Db.Name, cfg.Db.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database")

	//REPOSITORIES
	authRepo := authRepository.NewAuthRepository(db)
	userRepo := UserRepository.NewUserRepositoryImpl(db)
	userDetailsRepo := UserDetailsRepository.NewUserDetailsRepositoryImpl(db)
	startUpRepo := StartUpRepository.NewStartUpRepository(db)

	//SERVICES
	authServ := authService.NewAuthService(authRepo, userRepo, userDetailsRepo)
	profileServ := profileService.NewProfileService(userRepo, userDetailsRepo)
	startUpServ := StartUpService.NewStartUpService(startUpRepo)

	//HANDLERS
	authHand := authHandler.NewAuthHandler(authServ)
	profileHand := profileHandler.NewProfileHandler(profileServ)
	startUpHand := startUpHandler.NewStartUpHandler(startUpServ)

	//REGISTER_ROUTES
	authHand.RegisterRoutes(router)
	profileHand.RegisterRoutes(router)
	startUpHand.RegisterRoutes(router)

	srv := server.NewAPIServer(cfg, router)
	log.Println("Server success running")
	srv.Start()
}
