package summoner

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ByName(user string) (Summoner, error){
	var summoner Summoner
	const riotURL = "https://kr.api.riotgames.com/lol/summoner/v4/summoners/by-name/"
	err := godotenv.Load("secret.env")
	riotKey := os.Getenv("RIOT_KEY")
	resp, err := http.Get(riotURL + user + "?api_key=" + riotKey)
	if err != nil {
		return summoner, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&summoner)
	return summoner, nil
}
