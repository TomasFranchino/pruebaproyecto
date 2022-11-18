package models

import (
	"time"

	"gorm.io/gorm"
)

type Pelicula struct {
	gorm.Model              // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	Nombre        string    `json:"nombre"`
	Idioma        string    `json:"idioma"`
	FechaAlquiler time.Time `json:"fechaAlquiler"`
	Productora    string    `json:"productora"`
	Actores       string    `json:"actores"`
	PaisDeOrigen  string    `json:"paisDeOrigen"`
	GeneroID      int       `json:"generoTypeID"`
	Genero        Genero
}
