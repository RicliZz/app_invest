package app

import (
	"fmt"
	docs "github.com/RicliZz/app_invest/cmd/docs"
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
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Run(configpath string) {
	//config
	cfg := config.InitConfig(configpath)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.Username, os.Getenv("DB_PASSWORD"), cfg.Db.Name, cfg.Db.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database")

	//redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "ricliznedokyrill",
		DB:       0,
	})
	log.Println("Successfully connected to the redis server")

	//REPOSITORIES
	authRepo := authRepository.NewAuthRepository(db)
	userRepo := UserRepository.NewUserRepositoryImpl(db)
	userDetailsRepo := UserDetailsRepository.NewUserDetailsRepositoryImpl(db)
	startUpRepo := StartUpRepository.NewStartUpRepository(db)

	//SERVICES
	authServ := authService.NewAuthService(authRepo, userRepo, userDetailsRepo, rdb)
	profileServ := profileService.NewProfileService(userRepo, userDetailsRepo)
	startUpServ := StartUpService.NewStartUpService(startUpRepo)

	//HANDLERS
	authHand := authHandler.NewAuthHandler(authServ)
	profileHand := profileHandler.NewProfileHandler(profileServ)
	startUpHand := startUpHandler.NewStartUpHandler(startUpServ)

	//SWAG AND DEFAULT_ROUTER
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"
	api := router.Group("/api/v1/")

	//REGISTER_ROUTES
	authHand.RegisterRoutes(api)
	profileHand.RegisterRoutes(api)
	startUpHand.RegisterRoutes(api)

	srv := server.NewAPIServer(cfg, router)
	log.Println("Server success running")
	srv.Start()
}
