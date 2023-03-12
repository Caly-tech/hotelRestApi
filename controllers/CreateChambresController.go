
package controllers 
import(
	"github.com/gin-gonic/gin"
	"hotelRestApi/initializer"
	"hotelRestApi/models"
)

// function qui permet d'ajouter une categorie au niveau de la base de données
func CreateCategoriesController(c *gin.Context) {
	var bodyCat struct {
		NomCategorie string
		Tarifs int
		Chambres[] models.Chambres
	}
	c.Bind(&bodyCat)

	categorie := models.Categories{NomCategorie:bodyCat.NomCategorie, Tarifs:bodyCat.Tarifs, ChambresCategories:bodyCat.Chambres}
	categorieResult := initializer.DB.Create(&categorie)

	if categorieResult.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"categories": categorie,
	})
}

// function qui permet de lister les chambres au niveau de la base données
func ShowChambresController(c *gin.Context) {
	chambres := []models.Chambres{}
	initializer.DB.Find(&chambres)
	// result := initializer.DB.Table("chambres").Select("chambre.* , categories.*").Joins("Join categories on chambres.categories_id=categories.id").Scan(&chambres)
	c.JSON(200, gin.H{
		"Chambres": &chambres,
	})
}

// la fonction qui permet d'ajouter une chambre au niveau de la base de données
func CreateChambresController(c *gin.Context) {
	var bodyChambres struct {
	EtatChambre string
	NbreLits int
	Description string
	Capacites int 
	Services string 
	Surfaces int 
	CategoriesID uint
	ReservationChambres[] models.Reservations
	ImageChambres[] models.Images
	}
	c.Bind(&bodyChambres)

	chambres := models.Chambres{
				EtatChambre:bodyChambres.EtatChambre, 
				NbreLits:bodyChambres.NbreLits, 
				Description:bodyChambres.Description, 
				Capacites:bodyChambres.Capacites, 
				Services:bodyChambres.Services, 
				Surfaces:bodyChambres.Surfaces, 
				CategoriesID:bodyChambres.CategoriesID, 
				ReservationChambres:bodyChambres.ReservationChambres, 
				ImageChambres:bodyChambres.ImageChambres}
				
	Chambresresult := initializer.DB.Create(&chambres)

	if Chambresresult.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Chambres": chambres,
	})
	 

}