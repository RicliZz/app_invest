package main

import "github.com/RicliZz/app_invest/internal/app"

const configpath = "config.yml"

func main() {
	app.Run(configpath)
}
