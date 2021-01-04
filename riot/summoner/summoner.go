package summoner

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountID"`
	PuUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummoerLevel  int    `json:"summonerLevel"`
}
