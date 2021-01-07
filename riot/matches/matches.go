package matches

type Matches struct {
	GameID       int     `json:"gameId"`
	PlatformID   string  `json:"platformId"`
	GameCreation int     `json:"gameCreation"`
	GameDuration int     `json:"gameDuration"`
	QueueID      int     `json:"queueId"`
	MapID        int     `json:"mapId"`
	SeasonID     int     `json:"seasonId"`
	GameVersion  string  `json:"gameVersion"`
	GameMode     string  `json:"gameMode"`
	GameType     string  `json:"gameType"`
	Teams        []Teams `json:"teams"`
}

type Teams struct {
	TeamID               int    `json:"teamId"`
	Win                  string `json:"win"`
	FirstBlood           bool   `json:"firstBlood"`
	FirstTower           bool   `json:"firstTower"`
	FirstInhibitor       bool   `json:"firstInhibitor"`
	FirstBaron           bool   `json:"firstBaron"`
	FirstDragon          bool   `json:"firstDragon"`
	FirstRiftHerald      bool   `json:"firstRiftHerald"`
	TowerKills           int    `json:"towerKills"`
	InhibitorKills       int    `json:"inhibitorKills"`
	BaronKills           int    `json:"baronKills"`
	DragonKills          int    `json:"dragonKills"`
	VilemawKills         int    `json:"vilemawKills"`
	RiftHeraldKills      int    `json:"riftHeraldKills"`
	DominionVictoryScore int    `json:"dominionVictoryScore"`
	Bans                 []Bans `json:"bans"`
}

type Bans struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}
