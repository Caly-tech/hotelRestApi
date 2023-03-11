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
	initializer.DB.AutoMigrate(&models.Chambres{})
	initializer.DB.AutoMigrate(&models.Reservations{})
	initializer.DB.AutoMigrate(&models.Facturations{})
	initializer.DB.AutoMigrate(&models.{})
}
