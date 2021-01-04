package main

import (
	"WhoPoo_Server/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/match/:name", controller.MatchByName)
	e.Logger.Fatal(e.Start(":3000"))
}
