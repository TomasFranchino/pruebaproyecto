package dtos

type NuevoSocioDTO struct {
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	CorreoElectronico string `json:"email"`
	TipoDocumento     string `json:"documentType"`
	NumeroDocumento   string `json:"documentNumber"`
	Usuario           string `json:"user"`
	Contrasenia       string `json:"password"`
}

type SelectSocioDTO struct {
	ID                int    `json:"id"`
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	CorreoElectronico string `json:"email"`
	TipoDocumento     string `json:"ducumentType"`
	NumeroDocumento   string `json:"documentNumber"`
}

type ModificarSocioDTO struct {
	Nombre            string `json:"firstName"`
	Apellido          string `json:"lastName"`
	CorreoElectronico string `json:"email"`
	TipoDocumento     string `json:"documentType"`
	NumeroDocumento   string `json:"documentNumber"`
}
