package league

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func BySummonerID(summonerID string) ([]LeagueDto, error) {
	var league []LeagueDto
	const riotURL = "https://kr.api.riotgames.com/lol/league/v4/entries/by-summoner/"
	_ = godotenv.Load("secret.env")
	riotKey := os.Getenv("RIOT_KEY")
	resp, err := http.Get(riotURL + summonerID + "?api_key=" + riotKey)
	if err != nil {
		return league, errors.New("API 요청을 통해 정보를 받아오는 중 에러 발생")
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&league)
	if league == nil {
		return league, errors.New("API 토큰 만료")
	} else if len(league) == 0 {
		return league, errors.New("현재 소속되어 있는 리그 및 티어 정보가 없습니다")
	}
	return league, nil
}
