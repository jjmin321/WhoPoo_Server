package match

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ByAccountID(accountID, startIndex, endIndex string) (Match, error) {
	var match Match
	const riotURL = "https://kr.api.riotgames.com/lol/match/v4/matchlists/by-account/"
	err := godotenv.Load("secret.env")
	riotKey := os.Getenv("RIOT_KEY")
	resp, err := http.Get(riotURL + accountID + "?beginIndex=" + startIndex + "&endIndex=" + endIndex + "&api_key=" + riotKey)
	if err != nil {
		return match, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&match)
	return match, nil
}
