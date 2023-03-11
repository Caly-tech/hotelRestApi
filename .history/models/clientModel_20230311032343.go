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
	NomCategorie string
	Tarifs int
	ChambresCategories []Chambres
}

// table Chambres 
type Chambres struct {
	gorm.Model
	EtatChambre string
	NbreLits int 
	Description string 
	Capacites int
	Services string 
	Surfaces int 
	CategoriesID uint
	ReservationChambres []Reservations
}

// table Reservation 
type Reservations struct {
	gorm.Model
	DateDentrer time.Time
	DateSortie time.Time
	Nuite int
	ClientsID uint
	ChambresID uint
}

// table Facturation 
type Facturations struct {
	gorm.Model
	Date time.Time
	ClientsID uint
}

// table images 

type Images struct {
	gorm.Model
	NomImages string
	
}




