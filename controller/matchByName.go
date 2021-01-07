package controller

import (
	"WhoPoo_Server/riot/match"
	"WhoPoo_Server/riot/summoner"

	"github.com/labstack/echo"
)

func MatchByName(c echo.Context) error {
	name, startIndex, endIndex := c.Param("name"), c.QueryParam("startIndex"), c.QueryParam("endIndex")
	summoner, _ := summoner.ByName(name)
	match, err := match.ByAccountID(summoner.AccountID, startIndex, endIndex)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": 200,
		"result": match,
	})
}
