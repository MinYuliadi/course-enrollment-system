package main

import (
	"course-enrollment-system/config"
	"course-enrollment-system/routers"
)

func main() {
	config.InitEnv()
	config.InitDB()

	r := routers.InitRouter()
	r.Run(":8080")
}
