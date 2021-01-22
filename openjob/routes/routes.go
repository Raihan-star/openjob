package routes

import (
	"net/http"
	"openjob/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HALLO SELAMAT DATANG DI WEB OPEN DONE JOB RAIHAN DAN AGUNG ")
	})
	e.GET("/openjob", controllers.FetchAllOpenjob)
	e.POST("/post_openjob", controllers.StoreOpenjob)
	e.PUT("/put_openjob", controllers.UpdateOpenjob)
	e.DELETE("/delete_openjob", controllers.DeleteOpenjob)
	//e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	//e.POST("/login", controllers.CheckLogin)

	return e
}
