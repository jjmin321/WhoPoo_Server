package summoner

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type Body struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountID"`
	PuUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummoerLevel  int    `json:"summonerLevel"`
}

func ByName(c echo.Context) error {
	var body Body
	const riotURL = "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-name/"
	err := godotenv.Load("secret.env")
	user := c.Param("user")
	riotKey := os.Getenv("RIOT_KEY")
	resp, err := http.Get(riotURL + user + "?api_key=" + riotKey)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  500,
			"message": "API 토큰 만료",
		})
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&body)
	return c.JSON(200, map[string]interface{}{
		"result": body,
	})
}
