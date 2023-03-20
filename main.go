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
	r.POST("/createCategories", controllers.CreateCategoriesController) 
	r.GET("/showCategories", controllers.ShowCategoriesController)	
	r.POST("/createChambres", controllers.CreateChambresController) 
	r.GET("/showChambres", controllers.ShowChambresController) 
	r.GET("/showChambres/:id", controllers.ShowIDChambresController) 
	r.PUT("/updateChambres/:id", controllers.UpdateChambresController)

	/* les routes concernant les clients */
	r.POST("/createClients", controllers.CreateClientsController) 
	r.GET("/showClients", controllers.ShowClientsController) 
	r.GET("/showClients/:id", controllers.ShowIDClientsController) 
	r.PUT("/updateClients/:id", controllers.UpdateClientsController)

	/* les routes concernant les clients */
	r.POST("/createReservation", controllers.CreateReservationsController)
	r.GET("/showReservations", controllers.ShowReservationsController)  
	r.GET("/showReservations/:id", controllers.ShowIDReservationsController)
	r.PUT("/updateReservations/:id", controllers.UpdateReservationsController)

	r.Run() // listen and serve on 127.0.0.0:3000
}
