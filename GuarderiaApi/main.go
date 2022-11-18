// main.go
package main

import (
	"videotecaapi/controllers"
	"videotecaapi/db"
	"videotecaapi/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Cargo las variables de entorno
	Environment()

	// Inicializo la base de datos
	db.DBConn = db.Init()

	// Cargo las migraciones de las tablas a la base de datos
	db.SeedDatabase(db.DBConn)

	// Configuracion del servidor con los parámetros por defecto.
	router := gin.Default()

	// Inyecto un middlewares con los CORS soportados
	router.Use(middlewares.CORS())

	// Creo las Rutas
	setupRoutes(router)

	// Si se especificó un endpint que no tiene su correspondiente controller,
	// retornar un codigo 404
	//router.NoRoute(func(c *gin.Context) {
	// c.AbortWithStatus(404) // Resource not found -> https://en.wikipedia.org/wiki/HTTP_404
	//})

	router.Run() // listen and serve on 0.0.0.0:8080

	// Cierro la base de datos antes de finalizar la aplicación
	defer db.Close()
}

// Creación de Rutas
func setupRoutes(router *gin.Engine) {

	// para el endpoint /api/ping le indico que controlador va a estar respondiendo a esa petición
	apiController := new(controllers.PingController)
	router.GET("/ping", apiController.Get)

	// Descomentar esta linea si se desea tener las api segurizadas.
	//v1.Use(middlewares.Authentication)

	// Agrego las rutas para la versión v1 de la api del recurso Genero
	v1 := router.Group("/api/v1/genero")
	generoController := new(controllers.GeneroController)
	v1.GET("/", generoController.Find)
	v1.POST("/", generoController.Create)
	v1.GET("/:generoID", generoController.Get)
	v1.PUT("/:generoID", generoController.Update)
	v1.DELETE("/:generoID", generoController.Delete)

	// Agrego las rutas para la versión v1 de la api del recurso Tipo Documento
	v1 = router.Group("/api/v1/tipodocumento")

	TipoDocumentoController := new(controllers.TipoDocumentoController)
	v1.GET("/", TipoDocumentoController.Find)
	v1.POST("/", TipoDocumentoController.Create)
	v1.GET("/:tipoDocumentoID", TipoDocumentoController.Get)
	v1.PUT("/:tipoDocumentoID", TipoDocumentoController.Update)
	v1.DELETE("/:tipoDocumentoID", TipoDocumentoController.Delete)
	v1.GET("/deletes/", TipoDocumentoController.GetDeletes)
	v1.GET("/deletes/:tipoDocumentoID", TipoDocumentoController.GetDeletesByID)
	// Agrego las rutas para la versión v1 de la api del recurso Socio
	v1 = router.Group("/api/v1/socio")

	SocioController := new(controllers.SocioController)
	v1.GET("/", SocioController.Find)
	v1.POST("/", SocioController.Create)
	v1.GET("/:user/:password", SocioController.GetLogin)
	v1.GET("/:user", SocioController.Get)
	v1.PUT("/:socioID", SocioController.Update)
	v1.DELETE("/:socioID", SocioController.Delete)

	v1 = router.Group("/api/v1/pelicula")
	peliculaController := new(controllers.PeliculaController)
	v1.GET("/", peliculaController.Find)
	v1.POST("/", peliculaController.Create)
	v1.GET("/:peliculaID", peliculaController.Get)
	v1.PUT("/:peliculaID", peliculaController.Update)
	v1.DELETE("/:peliculaID", peliculaController.Delete)

	v1 = router.Group("/api/v1/alquiler")
	alquilerController := new(controllers.AlquilerController)
	v1.GET("/", alquilerController.Find)
	v1.POST("/", alquilerController.Create)
	v1.GET("/:alquilerID", alquilerController.Get)
	v1.PUT("/:alquilerID", alquilerController.Update)    // Si se va a permitir modificar un alquiler, hay que analizar bien el caso de uso para determinar que datos se van a cambiar.
	v1.DELETE("/:alquilerID", alquilerController.Delete) // NO debería poder borrarse un alquiler ya creado. SOLO es a fines demostrativo

}

// Environment variables
func Environment() {
	err := godotenv.Load()
	if err != nil {
		err := godotenv.Load("secrets/.env")
		if err != nil {
			panic(".env wasn't found.")
		}
	}
}
