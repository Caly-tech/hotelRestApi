package initializer

import (
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  var DB *gorm.DB 
  func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("fail to connect to database")
	}
  }