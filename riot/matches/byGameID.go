package matches

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ByGameID(gameID string) (Matches, error) {
	var matches Matches
	const riotURL = "https://kr.api.riotgames.com/lol/match/v4/matches/"
	_ = godotenv.Load("secret.env")
	riotKey := os.Getenv("RIOT_KEY")
	resp, err := http.Get(riotURL + gameID + "?api_key=" + riotKey)
	if err != nil {
		return matches, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&matches)
	if matches.Teams == nil {
		return matches, errors.New("API 토큰 만료")
	}
	return matches, nil
}
