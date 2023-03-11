package main

import (
	"hotelRestApi/initializer"
	"hotelRestApi/models"
)

func init(){
	initializer.EnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.Clients{})
	initializer.DB.AutoMigrate(&models.Categories{})
	initializer.DB.AutoMigrate(&models.Clients{})
	initializer.DB.AutoMigrate(&models.Clients{})
	initializer.DB.AutoMigrate(&models.Clients{})
	initializer.DB.AutoMigrate(&models.Clients{})
}
