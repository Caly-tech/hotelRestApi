package models 

import "gorm.io/gorm"

//table 
type Clients struct {
	gorm.Model
	Name string 
	Prenom string 
	telephone int
}


