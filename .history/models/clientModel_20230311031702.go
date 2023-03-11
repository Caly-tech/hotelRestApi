package models

import (
	"time"

	"gorm.io/gorm"
)

// table Clients
type Clients struct {
	gorm.Model
	Name string 
	Prenom string 
	telephone int
	Reservationclients []Reservation
}

// table categories 
type Categories struct {
	gorm.Model
	Nomcategorie string
	Tarifs int
}

// table Chambres 
type Chambres struct {
	gorm.Model
	Etatchambre string
	Nbrelits int 
	Description string 
	Capacites int
	Services string 
	Surfaces int 
}

// table Reservation 

type Reservation struct {
	gorm.Model
	Datedentrer time.Time
	Datedesortie time.Time
	Nuite int
}

// table Facturation 
type Facturation struct {
	gorm.Model
	Date time.Time
}




