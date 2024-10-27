package app

import (
	"github.com/RicliZz/app_invest/config"
	"github.com/RicliZz/app_invest/internal/handler"
	"github.com/RicliZz/app_invest/internal/server"
	"log"
)

func Run(configpath string) {
	cfg := config.InitConfig(configpath)
	handlers := handler.NewHandler()
	srv := server.NewServer(cfg, handlers.InitRoutes())
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
