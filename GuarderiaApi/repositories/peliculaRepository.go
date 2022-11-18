package repositories

import (
	"videotecaapi/db"
	"videotecaapi/models"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

// ModelRepository ...
type PeliculaRepository struct{}

func (rep PeliculaRepository) Find(context *gin.Context) paginate.Page {
	db := db.DBConn
	pg := paginate.New()

	model := db.Preload("Genero").Joins("Peliculas").Model(&models.Pelicula{})

	return pg.With(model).Request(context.Request).Response(&[]models.Pelicula{})
}

func (rep PeliculaRepository) Get(id int) *models.Pelicula {

	entity := new(models.Pelicula)

	db := db.DBConn
	db.Preload("Genero").First(&entity, id)

	return entity
}

func (rep PeliculaRepository) Insert(entity models.Pelicula) uint {
	db := db.DBConn

	db.Create(&entity)

	return entity.ID
}

func (rep PeliculaRepository) Update(ID int, entity models.Pelicula) int {

	entityToUpdate := new(models.Pelicula)

	db := db.DBConn
	db.First(&entityToUpdate, ID)

	result := db.Model(&entityToUpdate).Updates(models.Pelicula{
		Nombre:        entity.Nombre,
		Idioma:        entity.Idioma,
		FechaAlquiler: entity.FechaAlquiler,
		Productora:    entity.Productora,
		Actores:       entity.Actores,
		PaisDeOrigen:  entity.PaisDeOrigen,
		GeneroID:      entity.GeneroID,
	})

	return int(result.RowsAffected)
}

func (rep PeliculaRepository) Delete(ID int) int {

	entityToDelete := new(models.Pelicula)

	db := db.DBConn
	db.First(&entityToDelete, ID)

	result := db.Delete(&entityToDelete)

	return int(result.RowsAffected)
}
