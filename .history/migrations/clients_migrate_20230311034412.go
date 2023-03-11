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
}
