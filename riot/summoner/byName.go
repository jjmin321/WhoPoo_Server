package summoner

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ByName(user string) (Summoner, error) {
	var summoner Summoner
	const riotURL = "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-name/"
	err := godotenv.Load("secret.env")
	riotKey := os.Getenv("RIOT_KEY")
	resp, err := http.Get(riotURL + user + "?api_key=" + riotKey)
	if err != nil {
		return summoner, errors.New("API 요청을 통해 정보를 받아도는 중 에러 발생")
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&summoner)
	if summoner.ID == "" {
		return summoner, errors.New("API 토큰 만료")
	}
	return summoner, nil
}
