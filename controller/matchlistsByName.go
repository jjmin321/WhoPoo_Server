package controller

import (
	"WhoPoo_Server/riot/matchlists"
	"WhoPoo_Server/riot/summoner"

	"github.com/labstack/echo"
)

func MatchlistsByName(c echo.Context) error {
	name, startIndex, endIndex := c.Param("name"), c.QueryParam("startIndex"), c.QueryParam("endIndex")
	summoner, _ := summoner.ByName(name)
	matchlists, err := matchlists.ByAccountID(summoner.AccountID, startIndex, endIndex)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"result": matchlists,
	})
}
