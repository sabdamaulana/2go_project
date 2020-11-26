package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"project_2go/controllers"
	"project_2go/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello this is sab")
	})
	e.GET("/mahasiswa", controllers.FetchAllMahasiswa, middleware.IsAuthenticated)
	e.POST("/mahasiswa", controllers.StoreMahasiswa,middleware.IsAuthenticated)
	e.PUT("/mahasiswa", controllers.UpdateMahasiswa,middleware.IsAuthenticated)
	e.DELETE("/mahasiswa", controllers.DeleteMahasiswa,middleware.IsAuthenticated)
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)


	return e
}