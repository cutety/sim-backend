package main

import (
	_ "sim-backend/docs"
	"sim-backend/extension"
	"sim-backend/routes"
	"sim-backend/utils/viper"
)

func main() {
	viper.InitViper()
	extension.InitDB()
	routes.InitRouter()
}
