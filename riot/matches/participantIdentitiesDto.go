package matches

type ParticipantIdentitiesDto struct {
	ParticipantID int         `json:"participantId"`
	Player        []PlayerDto `json:"player"`
}

type PlayerDto struct {
	profileIcon       int    `json:"profileIcon"`
	accountId         string `json:"accountId"`
	matchHistoryUri   string `json:"matchHistoryUri"`
	currentAccountId  string `json:"currentAccountId"`
	currentPlatformId string `json:"currentPlatformId"`
	summonerName      string `json:"summonerName"`
	summonerId        string `json:"summonerId"`
	platformId        string `json:"platformId"`
}
