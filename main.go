package main

import (
	"WhoPoo_Server/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/matches/:gameId", controller.MatchesByGameID)
	e.GET("/matchlists/:name", controller.MatchlistsByName)
	e.Logger.Fatal(e.Start(":3000"))
}
