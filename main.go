package main

import (
	"sosialMedia/configs"
	"sosialMedia/routes"
)

func main() {
	configs.InitEnv()
	configs.InitDB()
	configs.InitMigrate()

	r := routes.InitRoutes()
	r.Run()
}
