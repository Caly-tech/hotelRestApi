package controllers

import(
	"github.com/gin-gonic/gin"
	"hotelRestApi/initializer"
	"hotelRestApi/models"
	"time"
)

// cette fonction d'ajouter une reservation 
func CreateReservationsController(c *gin.Context){
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

func ShowReservationsController(c *gin.Context) {
	reservations := models.Reservations{}

	initializer.DB.Find(&reservations)

	c.JSON(200, gin.H{
		"reservations": &reservations,
	})
}

func ShowIDReservationsController(c *gin.Context) {
	id := c.Param("id")

	reservations := []models.Reservations{}
	initializer.DB.Find(&reservations, id)

	c.JSON(200, gin.H{
		"reservations": &reservations,
	})
}

func UpdateReservationsController(c *gin.Context) {

	var bodyReservations struct {
		DateDentrer time.Time
		DateSortie time.Time
		Nuite int
		ClientsID uint
		ChambresID uint
	}
	c.Bind(&bodyReservations)
	id := c.Param("id")
	
	reservations := []models.Reservations{}
	initializer.DB.Find(&reservations, id)

	initializer.DB.Model(&reservations).Updates(models.Reservations{DateDentrer:bodyReservations.DateDentrer, DateSortie:bodyReservations.DateSortie, Nuite:bodyReservations.Nuite, ClientsID:bodyReservations.ClientsID, ChambresID:bodyReservations.ChambresID})

	c.JSON(200, gin.H{
		"reservations": &reservations,
	})

}