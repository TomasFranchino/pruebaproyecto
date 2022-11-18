package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"videotecaapi/models"
	"videotecaapi/repositories"
	"videotecaapi/utils"
)

type PeliculaController struct{}

func (controller PeliculaController) Get(context *gin.Context) {
	ID := context.Param("peliculaID")

	id, err := strconv.Atoi(ID) // se convierte un string a int
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.PeliculaRepository)
	entity := entityRep.Get(id)

	if entity.ID == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, entity)
	}
}

func (controller PeliculaController) Find(context *gin.Context) {

	entityRep := new(repositories.PeliculaRepository)
	entities := entityRep.Find(context)
	context.JSON(http.StatusOK, entities)
}

func (controller PeliculaController) Create(context *gin.Context) {

	entity := new(models.Pelicula)

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error()))
		return
	}

	rep := new(repositories.PeliculaRepository)
	id := rep.Insert(*entity)

	context.JSON(http.StatusCreated, id)
}

func (controller PeliculaController) Update(context *gin.Context) {

	entity := new(models.Pelicula)

	ID := context.Param("peliculaID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	if err := context.BindJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, utils.Error(err.Error())) // NO MOSTRAR ERRORES DIRECTOS del SISTEMA. (NO SE HACE!!!!)
		return
	}

	rep := new(repositories.PeliculaRepository)
	rowAffected := rep.Update(id, *entity)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, rowAffected)
	}
}

func (controller PeliculaController) Delete(context *gin.Context) {
	ID := context.Param("peliculaID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.Error("ID is missing."))
		return
	}

	entityRep := new(repositories.PeliculaRepository)
	rowAffected := entityRep.Delete(id)

	if rowAffected == 0 {
		context.JSON(http.StatusNotFound, "")
	} else {
		context.JSON(http.StatusOK, "")
	}
}
