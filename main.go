package main

import (
	"api-goklaster/config"
	v1 "api-goklaster/v1"
)

func main() {
	config.InitDB()
	app := v1.GinRoutes()
	app.Run(":666")
}
