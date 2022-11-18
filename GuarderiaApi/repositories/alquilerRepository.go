package repositories

import (
	"videotecaapi/db"
	"videotecaapi/models"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

// ModelRepository ...
type AlquilerRepository struct{}

func (rep AlquilerRepository) Find(context *gin.Context) paginate.Page {
	db := db.DBConn
	pg := paginate.New()

	model := db.Preload("Socio.TipoDocumento").Preload("Peliculas.Genero").Model(&models.Alquiler{})

	return pg.With(model).Request(context.Request).Response(&[]models.Alquiler{})
}

func (rep AlquilerRepository) Get(id int) *models.Alquiler {

	entity := new(models.Alquiler)

	db := db.DBConn
	db.Preload("Socio.TipoDocumento").Preload("Peliculas.Genero").First(&entity, id)

	return entity
}

func (rep AlquilerRepository) Insert(entity models.Alquiler) uint {
	db := db.DBConn

	db.Create(&entity)

	return entity.ID
}

func (rep AlquilerRepository) Update(ID int, entity models.Alquiler) int {

	entityToUpdate := new(models.Alquiler)

	db := db.DBConn
	db.First(&entityToUpdate, ID)

	result := db.Model(&entityToUpdate).Updates(models.Alquiler{
		Importe:     entity.Importe,
		FechaCierre: entity.FechaCierre,
		Abonado:     entity.Abonado,
		SocioID:     entity.SocioID,
		Peliculas:   entity.Peliculas,
	})

	return int(result.RowsAffected)
}

func (rep AlquilerRepository) Delete(ID int) int {

	entityToDelete := new(models.Alquiler)

	db := db.DBConn
	db.First(&entityToDelete, ID)

	result := db.Delete(&entityToDelete)

	return int(result.RowsAffected)
}
