package models

import (
	"gorm.io/gorm"
)

type Socio struct {
	gorm.Model               // -> el gorm.Model implementa el ID, CreatedAt, UpdatedAt, DeletedAt
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	CorreoElectronico string `json:"email"`
	TipoDocumento     string `json:"typeDocument"`
	NumeroDocumento   string `json:"documentNumber"`
	Usuario           string `json:"user"`
	Contrasenia       string `json:"password"`
}
