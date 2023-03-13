package controllers 

import (
	"github.com/gin-gonic/gin"
	"hotelRestApi/models"
	"hotelRestApi/initializer"
)

func CreateClientsController(c *gin.Context) {

	var bodyClients struct {
		Name string 
		Prenom string 
		telephone int
		ReservationClients[]	models.Reservations
		FacturationClients[]	models.Facturations
	}

	c.Bind(&bodyClients)

	newClients := models.Clients{Name:bodyClients.Name,Prenom:bodyClients.Prenom,ReservationClients: bodyClients.ReservationClients,FacturationClients: bodyClients.FacturationClients}
	clientsResult := initializer.DB.Create(&newClients)

	if clientsResult.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"Client": newClients,
	})

}

func ShowClientsController(c *gin.Context) {
	clients := []models.Clients{}
	initializer.DB.Find(&clients)
	c.JSON(200, gin.H{
		"Chambres": &clients,
	})
}

func ShowIDClientsController(c *gin.Context) {
	id := c.Param("id")

	clients := []models.Clients{}
	initializer.DB.Find(&clients, id)

	c.JSON(200, gin.H{
		"Chambres": &clients,
	})
}
