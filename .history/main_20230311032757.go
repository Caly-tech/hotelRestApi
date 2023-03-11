package main

// import "fmt"


import (
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

    "log"
)

func init(){

}


func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Api en cour ...........",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
