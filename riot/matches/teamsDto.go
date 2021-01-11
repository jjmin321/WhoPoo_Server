package matches

type TeamsDto struct {
	TeamID               int       `json:"teamId"`
	Win                  string    `json:"win"`
	FirstBlood           bool      `json:"firstBlood"`
	FirstTower           bool      `json:"firstTower"`
	FirstInhibitor       bool      `json:"firstInhibitor"`
	FirstBaron           bool      `json:"firstBaron"`
	FirstDragon          bool      `json:"firstDragon"`
	FirstRiftHerald      bool      `json:"firstRiftHerald"`
	TowerKills           int       `json:"towerKills"`
	InhibitorKills       int       `json:"inhibitorKills"`
	BaronKills           int       `json:"baronKills"`
	DragonKills          int       `json:"dragonKills"`
	VilemawKills         int       `json:"vilemawKills"`
	RiftHeraldKills      int       `json:"riftHeraldKills"`
	DominionVictoryScore int       `json:"dominionVictoryScore"`
	Bans                 []BansDto `json:"bans"`
}

type BansDto struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}
