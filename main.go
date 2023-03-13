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

	/* les routes concernant les chambres et leurs categories */
	r.POST("/createCategories", controllers.CreateCategoriesController) // route pour ajouter une categorie
	r.POST("/createChambres", controllers.CreateChambresController) // route pour ajouter une Chambre 
	r.GET("/showChambres", controllers.ShowChambresController) // route pour Lister les chambres
	r.GET("/showChambres/:id", controllers.ShowIDChambresController) //route pour lister une chambre en fonction de son id 
	r.PUT("/showChambres/:id", controllers.UpdateChambresController)// route pour modifier les informations d'une chambre

	/* les routes concernant les clients */
	r.POST("/createClients", controllers.CreateClientsController) // route pour ajouter un client
	r.GET("/showClients", controllers.ShowClientsController) // route pour Lister les clients
	r.GET("/showClients/:id", controllers.ShowIDClientsController) //route pour lister un client en fonction de son id 

	r.Run() // listen and serve on 127.0.0.0:3000
}
