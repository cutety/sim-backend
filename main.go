package main

import (
	_ "sim-backend/docs"
	"sim-backend/extension"
	"sim-backend/routes"
	"sim-backend/utils/viper"
)
// @title Swagger API
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
func main() {
	viper.InitViper()
	extension.InitDB()
	routes.InitRouter()
}
