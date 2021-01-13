package controller

import (
	"WhoPoo_Server/riot/league"
	"WhoPoo_Server/riot/summoner"

	"github.com/labstack/echo"
)

func LeagueByName(c echo.Context) error {
	name := c.Param("name")
	summoner, _ := summoner.ByName(name)
	league, err := league.BySummonerID(summoner.ID)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"result": league,
	})
}
