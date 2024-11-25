package main

import (
	"github.com/RicliZz/app_invest/internal/app"
)

const config_path = "config.yml"

func main() {
	app.Run(config_path)
}
