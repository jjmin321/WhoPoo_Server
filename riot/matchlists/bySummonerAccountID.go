package matchlists

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func ByAccountID(accountID, startIndex, endIndex string) (MatchlistsDto, error) {
	var matchlists MatchlistsDto
	const riotURL = "https://kr.api.riotgames.com/lol/match/v4/matchlists/by-account/"
	_ = godotenv.Load("secret.env")
	riotKey := os.Getenv("RIOT_KEY")
	resp, err := http.Get(riotURL + accountID + "?beginIndex=" + startIndex + "&endIndex=" + endIndex + "&api_key=" + riotKey)
	if err != nil {
		return matchlists, errors.New("API 요청을 통해 정보를 받아오는 중 에러 발생")
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&matchlists)
	if matchlists.Matches == nil {
		return matchlists, errors.New("API 토큰 만료")
	}
	for i, v := range matchlists.Matches {
		unixTime, _ := strconv.Atoi(strconv.Itoa(v.Timestamp)[:10])
		matchlists.Matches[i].Time = time.Unix(int64(unixTime), 0)
	}
	return matchlists, nil
}
