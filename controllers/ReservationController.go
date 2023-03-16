package controllers

import(
	"github.com/gin-gonic/gin"
	"hotelRestApi/initializer"
	"hotelRestApi/models"
	"time"
)

// cette fonction d'ajouter une reservation 
func CreateReservation(c *gin.Context){
	var bodyReservations struct {
		DateDentrer time.Time
		DateSortie time.Time
		Nuite int
		ClientsID uint
		ChambresID uint
	}
	c.Bind(&bodyReservations)
	reservations := models.Reservations{DateDentrer:bodyReservations.DateDentrer, DateSortie:bodyReservations.DateSortie, Nuite:bodyReservations.Nuite, ClientsID:bodyReservations.ClientsID, ChambresID:bodyReservations.ChambresID}
	resultReservations := initializer.DB.Create(&reservations)
	

	if resultReservations.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"reservations": &reservations,
	})
}