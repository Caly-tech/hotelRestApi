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
	ReservationClients []Reservations
	FacturationClients []Facturations
}

// table categories 
type Categories struct {
	gorm.Model
	Nomcategorie string
	Tarifs int
	ChambresCategories []Chambres
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

type Reservations struct {
	gorm.Model
	Datedentrer time.Time
	Datedesortie time.Time
	Nuite int
	ClientsID uint
}

// table Facturation 
type Facturations struct {
	gorm.Model
	Date time.Time
	ClientsID uint
}




