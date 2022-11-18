package models

import (
	"gorm.io/gorm"
)

type TipoDocumento struct {
	gorm.Model        // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	Tipo       string `gorm:"uniqueIndex:idx_tipodocumento" json:"tipo"`
	Nombre     string `json:"nombre"`
}
