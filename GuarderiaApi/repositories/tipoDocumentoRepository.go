package repositories

import (
	"videotecaapi/db"
	"videotecaapi/models"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

// ModelRepository ...
type TipoDocumentoRepository struct{}

func (rep TipoDocumentoRepository) Find(context *gin.Context) paginate.Page {
	db := db.DBConn
	pg := paginate.New()

	model := db.Joins("Tipo_Documentos").Model(&models.TipoDocumento{})

	return pg.With(model).Request(context.Request).Response(&[]models.TipoDocumento{})
}

func (rep TipoDocumentoRepository) Get(id int) *models.TipoDocumento {

	entity := new(models.TipoDocumento)

	db := db.DBConn
	db.First(&entity, id)

	return entity
}

func (rep TipoDocumentoRepository) Insert(entity models.TipoDocumento) (uint, error) {
	db := db.DBConn

	result := db.Create(&entity)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return entity.ID, nil
	}
}

func (rep TipoDocumentoRepository) Update(ID int, entity models.TipoDocumento) (int, error) {

	entityToUpdate := new(models.TipoDocumento)

	db := db.DBConn
	db.First(&entityToUpdate, ID)

	result := db.Model(&entityToUpdate).Updates(map[string]interface{}{"tipo": entity.Tipo, "nombre": entity.Nombre})

	if result.Error != nil {
		return 0, result.Error
	} else {
		return int(result.RowsAffected), nil
	}
}

func (rep TipoDocumentoRepository) Delete(ID int) (int, error) {

	entityToDelete := new(models.TipoDocumento)

	db := db.DBConn
	db.First(&entityToDelete, ID)

	result := db.Delete(&entityToDelete)

	if result.Error != nil {
		return 0, result.Error
	} else {
		return int(result.RowsAffected), nil
	}
}

func (rep TipoDocumentoRepository) GetByDocumentName(documentName string) *models.TipoDocumento {

	entity := new(models.TipoDocumento)

	db := db.DBConn

	// Get first matched record
	db.Where("tipo = ?", documentName).First(&entity)
	// SELECT * FROM tipodocumentos WHERE nombre = 'algo_a_buscar' ORDER BY id LIMIT 1;

	return entity
}

func (rep TipoDocumentoRepository) GetDeletes(context *gin.Context) paginate.Page {

	db := db.DBConn

	pg := paginate.New()

	model := db.Unscoped().Where("deleted_at IS NOT null").Model(&models.TipoDocumento{})

	return pg.Response(model, context.Request, &[]models.TipoDocumento{})
}

func (rep TipoDocumentoRepository) GetDeletesByID(ID string) *models.TipoDocumento {

	entity := new(models.TipoDocumento)

	db := db.DBConn

	db.Unscoped().Where("deleted_at IS NOT null AND id = ? ", ID).First(&entity)

	return entity
}
