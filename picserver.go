package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"pictureServer/controller"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.POST("/upload", controller.UploadFile)
	e.Static("/files", "files")
	e.Logger.Fatal(e.Start(":1323"))

}
