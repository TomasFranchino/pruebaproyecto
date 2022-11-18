package repositories

import (
	"videotecaapi/db"
	"videotecaapi/models"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

// ModelRepository ...
type SocioRepository struct{}

func (rep SocioRepository) Find(context *gin.Context) paginate.Page {
	db := db.DBConn
	pg := paginate.New()

	model := db.Joins("Socios").Model(&models.Socio{})

	return pg.With(model).Request(context.Request).Response(&[]models.Socio{})
}

func (rep SocioRepository) Get(id int) *models.Socio {

	entity := new(models.Socio)

	db := db.DBConn

	db.First(&entity, id)

	return entity
}

func (rep SocioRepository) Insert(entity models.Socio) uint {
	db := db.DBConn

	db.Create(&entity)

	return entity.ID
}

func (rep SocioRepository) Update(ID int, entity models.Socio) int {

	entityToUpdate := new(models.Socio)

	db := db.DBConn

	result := db.Model(&entityToUpdate).Updates(models.Socio{
		Nombre:            entity.Nombre,
		Apellido:          entity.Apellido,
		CorreoElectronico: entity.CorreoElectronico,
		TipoDocumento:     entity.TipoDocumento,
		NumeroDocumento:   entity.NumeroDocumento,
	})

	return int(result.RowsAffected)
}

func (rep SocioRepository) Delete(ID int) int {

	entityToDelete := new(models.Socio)

	db := db.DBConn
	db.First(&entityToDelete, ID)

	result := db.Delete(&entityToDelete)

	return int(result.RowsAffected)
}

func (rep SocioRepository) GetLogin(user string, password string) *models.Socio {

	entity := new(models.Socio)

	db := db.DBConn

	db.Where("usuario = ?", user).Where("contrasenia", password).First(&entity)

	return entity
}
