package controller

import (
	"WhoPoo_Server/riot/matches"

	"github.com/labstack/echo"
)

func MatchesByGameID(c echo.Context) error {
	gameID := c.Param("gameId")
	matches, err := matches.ByGameID(gameID)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"result": matches,
	})
}
