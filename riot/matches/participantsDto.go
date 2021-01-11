package matches

type ParticipantsDto struct {
	ParticipantID             int          `json:"participantId"`
	ChampionID                int          `json:"championId"`
	TeamID                    int          `json:"teamId"`
	Spell1                    int          `json:"spell1Id"`
	Spell2                    int          `json:"spell2Id"`
	HighestAchievedSeasonTier string       `json:"highestAchievedSeasonTier"`
	Runes                     RunesDto     `json:"runes"`
	Stats                     StatsDto     `json:"stats"`
	Timeline                  TimelineDto  `json:"timeline"`
	Masteries                 MasteriesDto `json:"masteries"`
}

type RunesDto struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

type StatsDto struct {
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	GoldEarned                      int  `json:"goldEarned"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	NodeNeutralizeAssist            int  `json:"nodeNeutralizeAssist"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	ChampLevel                      int  `json:"champLevel"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	Deaths                          int  `json:"deaths"`
	TripleKills                     int  `json:"tripleKills"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	WardsKilled                     int  `json:"wardsKilled"`
	PentaKills                      int  `json:"pentaKills"`
	DamageSelfMitigated             int  `json:"damageSelfMitigated"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	NodeNeutralize                  int  `json:"nodeNeutralize"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	NodeCapture                     int  `json:"nodeCapture"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	Kills                           int  `json:"kills"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	TurretKills                     int  `json:"turretKills"`
	ParticipantID                   int  `json:"participantId"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	NodeCaptureAssist               int  `json:"nodeCaptureAssist"`
	Assists                         int  `json:"assists"`
	TeamObjective                   int  `json:"teamObjective"`
	AltarsNeutralized               int  `json:"altarsNeutralized"`
	GoldSpent                       int  `json:"goldSpent"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	AltarsCaptured                  int  `json:"altarsCaptured"`
	Win                             bool `json:"win"`
	TotalHeal                       int  `json:"totalHeal"`
	UnrealKills                     int  `json:"unrealKills"`
	VisionScore                     int  `json:"visionScore"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
	KillingSprees                   int  `json:"killingSprees"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	DoubleKills                     int  `json:"doubleKills"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	QuadraKills                     int  `json:"quadraKills"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore6                    int  `json:"playerScore6"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore8                    int  `json:"playerScore8"`
	PlayerScore9                    int  `json:"playerScore9"`
	Perk0                           int  `json:"perk0"`
	Perk0Var1                       int  `json:"perk0Var1"`
	Perk0Var2                       int  `json:"perk0Var2"`
	Perk0Var3                       int  `json:"perk0Var3"`
	Perk1                           int  `json:"perk1"`
	Perk1Var1                       int  `json:"perk1Var1"`
	Perk1Var2                       int  `json:"perk1Var2"`
	Perk1Var3                       int  `json:"perk1Var3"`
	Perk2                           int  `json:"perk2"`
	Perk2Var1                       int  `json:"perk2Var1"`
	Perk2Var2                       int  `json:"perk2Var2"`
	Perk2Var3                       int  `json:"perk2Var3"`
	Perk3                           int  `json:"perk3"`
	Perk3Var1                       int  `json:"perk3Var1"`
	Perk3Var2                       int  `json:"perk3Var2"`
	Perk3Var3                       int  `json:"perk3Var3"`
	Perk4                           int  `json:"perk4"`
	Perk4Var1                       int  `json:"perk4Var1"`
	Perk4Var2                       int  `json:"perk4Var2"`
	Perk4Var3                       int  `json:"perk4Var3"`
	Perk5                           int  `json:"perk5"`
	Perk5Var1                       int  `json:"perk5Var1"`
	Perk5Var2                       int  `json:"perk5Var2"`
	Perk5Var3                       int  `json:"perk5Var3"`
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
	PerkSubStyle                    int  `json:"perkSubStyle"`
	StatPerk0                       int  `json:"statPerk0"`
	StatPerk1                       int  `json:"statPerk1"`
	StatPerk2                       int  `json:"statPerk2"`
}

type TimelineDto struct {
	ParticipantID               int                `json:"participantId"`
	CsDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	XpPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	XpDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	Lane                        string             `json:"lane"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
}

type MasteriesDto struct {
	Rank      int `json:"rank"`
	MasteryID int `json:"masteryId"`
}
