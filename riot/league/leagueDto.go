package league

type LeagueDto struct {
	LeagueID     string        `json:"leagueId"`
	SummonerID   string        `json:"summonerId"`
	SummonerName string        `json:"summonerName"`
	QueueType    string        `json:"queueType"`
	Tier         string        `json:"tier"`
	Rank         string        `json:"rank"`
	LeaguePoints int           `json:"leaguePoints"`
	Wins         int           `json:"wins"`
	Losses       int           `json:"losses"`
	HotStreak    bool          `json:"hotStreak"`
	Veteran      bool          `json:"veteran"`
	FreshBlood   bool          `json:"freshBlood"`
	Inactive     bool          `json:"inactive"`
	MiniSeries   MiniSeriesDto `json:"miniSeries"`
}

type MiniSeriesDto struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}
