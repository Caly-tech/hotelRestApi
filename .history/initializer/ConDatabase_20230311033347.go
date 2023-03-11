package initializer

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )
  var DB *gorm.DB 
  func main() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/hotelRestApi?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	
  }