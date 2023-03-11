package main

import (
	"hotelRestApi/initializer"
)

func init(){
	initializer.EnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate()
}
