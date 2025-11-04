package main

import (
	"vehicle-service-api/config"
	"vehicle-service-api/routers"
)

func main() {
	config.InitDB()

	r := routers.InitRouter()
	r.Run(":8080")
}
