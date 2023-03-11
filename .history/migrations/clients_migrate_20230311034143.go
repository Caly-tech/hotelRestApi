package main

import (
	"hotelRestApi/initializer"

	"github.com/gin-gonic/gin"
)

func init(){
	initializer.EnvVariables()
	initializer.ConnectToDB()
}
