package main

// import "fmt"

import (
	"hotelRestApi/controllers"
	"hotelRestApi/initializer"

	"github.com/gin-gonic/gin"
)

func init(){
	initializer.EnvVariables()
	initializer.ConnectToDB()
}


func main() {

	r := gin.Default()
	r.POST("/createCategories", controllers.CreateCategoriesController) // route pour ajouter une categorie
	r.GET("/showChambres", controllers.ShowChambresController) // route pour Lister les chambres
	r.POST("/createChambres", controllers.CreateChambresController) // route pour ajouter une Chambre 


	r.Run() // listen and serve on 127.0.0.0:3000
}
