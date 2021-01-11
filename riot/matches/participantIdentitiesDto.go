package matches

type ParticipantIdentitiesDto struct {
	ParticipantID int         `json:"participantId"`
	Player        []PlayerDto `json:"player"`
}

type PlayerDto struct {
	ProfileIcon       int    `json:"profileIcon"`
	AccountId         string `json:"accountId"`
	MatchHistoryUri   string `json:"matchHistoryUri"`
	CurrentAccountId  string `json:"currentAccountId"`
	CurrentPlatformId string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	SummonerId        string `json:"summonerId"`
	PlatformId        string `json:"platformId"`
}
