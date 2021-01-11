package matches

type ParticipantIdentitiesDto struct {
	ParticipantID int       `json:"participantId"`
	Player        PlayerDto `json:"player"`
}

type PlayerDto struct {
	ProfileIcon       int    `json:"profileIcon"`
	AccountID         string `json:"accountId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	CurrentAccountID  string `json:"currentAccountId"`
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	SummonerID        string `json:"summonerId"`
	PlatformID        string `json:"platformId"`
}
