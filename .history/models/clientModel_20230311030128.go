package models 

import "gorm.io/gorm"

type Clients struct {
	gorm.Model
	Name string 
	Pre
}