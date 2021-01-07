package matchlists

type Matchlists struct {
	Matches    []Matches `json:"matches"`
	StartIndex int       `json:"startIndex"`
	EndIndex   int       `json:"endIndex"`
	TotalGames int       `json:"totalGames"`
}

type Matches struct {
	PlatformID string `json:"platformId"`
	GameID     int    `json:"gameId"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Season     int    `json:"season"`
	Timestamp  int    `json:"timestamp"`
	Role       string `json:"role"`
	Lane       string `json:"lane"`
}
