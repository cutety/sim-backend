package main

import (
	"sim-backend/extension"
	"sim-backend/routes"
	"sim-backend/utils/viper"
)

func main() {
	viper.InitViper()
	extension.InitDB()
	routes.InitRouter()
}
