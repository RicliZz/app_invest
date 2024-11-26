package main

import (
	"github.com/RicliZz/app_invest/internal/app"
)

const config_path = "config.yml"

// @title HaveIdea
// @version 1.0
// @description This is a sample server for demonstrating JWT with Swagger.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app.Run(config_path)
}
