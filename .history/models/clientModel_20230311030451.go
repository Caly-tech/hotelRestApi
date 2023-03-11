package models 

import "gorm.io/gorm"

// table Clients
type Clients struct {
	gorm.Model
	Name string 
	Prenom string 
	telephone int
}

// table Chambres 

type Chambres struct {
	gorm.Model
	Etatchambre string
	Nbrelits int 
	Description string 
	Capacites string
	Services string 
	
}

