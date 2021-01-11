package matches 

type ParticipantsDto struct {
	ParticipantID int `json:"participantId"`
	ChampionID int `json:"championId"`
	TeamID int `json:"teamId"`
	Spell1	int	`json:"spell1Id"`
	Spell2 	int `json:"spell2Id"`
	HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
	Runes []RunesDto `json:"runes"`
	Stats []StatsDto `json:"stats"`
	Timeline []TimelineDto `json:"timeline"`
	Masteries	[]MasteriesDto	`json:"masteries"`
}

type RunesDto struct {
	RuneID	int	`json:"runeId"`
	Rank	int	`json:"rank"`
}

type StatsDto struct {
	item0	int	`json:"item0"`
	item2	int	`json:"item2"`
	TotalUnitsHealed	int	`json:"totalUnitsHealed"`
	item1	int		`json:"item1"`
	LargestMultiKill	int	`json:"largestMultiKill"`
	GoldEarned	int	`json:"goldEarned"`
	FirstInhibitorKill	bool	`json:"firstInhibitorKill"`
	PhysicalDamageTaken	int	`json:"physicalDamageTaken"`
	NodeNeutralizeAssist	int	`json:"nodeNeutralizeAssist"`
	totalPlayerScore	int	`json:"totalPlayerScore"`
	champLevel	int	`json:"champLevel"`
	damageDealtToObjectives	int	`json:"damageDealtToObjectives"`
	totalDamageTaken	int		`json:"totalDamageTaken"`
	neutralMinionsKilled	int	`json:"neutralMinionsKilled"`
	deaths	int		`json:"deaths"`
	tripleKills	int	`json:"tripleKills"`
	magicDamageDealtToChampions	int	`json:"magicDamageDealtToChampions"`
	wardsKilled	int	`json:"wardsKilled"`
	pentaKills	int	`json:"pentaKills"`
	damageSelfMitigated	int	`json:"damageSelfMitigated"`
	largestCriticalStrike	int	`json:"largestCriticalStrike"`
	nodeNeutralize	int	`json:"nodeNeutralize"`
	totalTimeCrowdControlDealt	int	`json:"totalTimeCrowdControlDealt"`
	firstTowerKill	bool	`json:"firstTowerKill"`
	magicDamageDealt	int		`json:"magicDamageDealt"`
	totalScoreRank	int		`json:"totalScoreRank"`
	nodeCapture	int		`json:"nodeCapture"`
	wardsPlaced	int	`json:"wardsPlaced"`
	totalDamageDealt	int	`json:"totalDamageDealt"`
	timeCCingOthers	int		`json:"timeCCingOthers"`
	magicalDamageTaken	int		`json:"magicalDamageTaken"`
	largestKillingSpree	int	`json:"largestKillingSpree"`
	totalDamageDealtToChampions	int	`json:"totalDamageDealtToChampions"`
	physicalDamageDealtToChampions	int		`json:"physicalDamageDealtToChampions"`
	neutralMinionsKilledTeamJungle	int	`json:"neutralMinionsKilledTeamJungle"`
	totalMinionsKilled	int	`json:"totalMinionsKilled"`
	firstInhibitorAssist	bool		`json:"firstInhibitorAssist"`
	visionWardsBoughtInGame	int		`json:"visionWardsBoughtInGame"`
	objectivePlayerScore	int		`json:"objectivePlayerScore"`
	kills	int	`json:"kills"`
	firstTowerAssist	bool	`json:"firstTowerAssist"`
	combatPlayerScore	int		`json:"combatPlayerScore"`
	inhibitorKills	int		`json:"inhibitorKills"`
	turretKills	int	`json:"turretKills"`
	participantId	int		`json:"participantId"`
	trueDamageTaken	int		`json:"trueDamageTaken"`
	firstBloodAssist	bool	`json:"firstBloodAssist"`
	nodeCaptureAssist	int		`json:"nodeCaptureAssist"`
	assists	int		`json:"assists"`
	teamObjective	int		`json:"teamObjective"`
	altarsNeutralized	int	`json:"altarsNeutralized"`
	goldSpent	int		`json:"goldSpent"`
	damageDealtToTurrets	int		`json:"damageDealtToTurrets"`
	altarsCaptured	int		`json:"altarsCaptured"`
	win	bool	`json:"win"`
	totalHeal	int	`json:"totalHeal"`
	unrealKills	int		`json:"unrealKills"`
	visionScore	int		`json:"visionScore"`
	physicalDamageDealt	int		`json:"physicalDamageDealt"`
	firstBloodKill	bool	`json:"firstBloodKill"`
	longestTimeSpentLiving	int		`json:"longestTimeSpentLiving"`
	killingSprees	int		`json:"killingSprees"`
	sightWardsBoughtInGame	int		`json:"sightWardsBoughtInGame"`
	trueDamageDealtToChampions	int		`json:"trueDamageDealtToChampions"`
	neutralMinionsKilledEnemyJungle	int	`json:"neutralMinionsKilledEnemyJungle"`
	doubleKills	int		`json:"doubleKills"`
	trueDamageDealt	int		`json:"trueDamageDealt"`
	quadraKills	int	`json:"quadraKills"`
	item4	int		`json:"item4"`
	item3 int `json:"item3"`
	item5 int `json:"item5"`
	item6 int `json:"item6"`
	playerScore1	int	`json:"playerScore1"`
	playerScore2	int	`json:"playerScore2"`
	playerScore3	int	`json:"playerScore3"`
	playerScore4	int	`json:"playerScore4"`
	playerScore5	int	`json:"playerScore5"`
	playerScore6	int	`json:"playerScore6"`
	playerScore7	int	`json:"playerScore7"`
	playerScore8	int	`json:"playerScore8"`
	playerScore9	int	`json:"playerScore9"`
	perk0	int `json:"perk0"`
	perk0Var1 int `json:"perk0Var1"`
	perk0Var2 int `json:"perk0Var2"`
	perk0Var3 int `json:"perk0Var3"`
	perk1	int `json:"perk1"`
	perk1Var1 int `json:"perk1Var1"`
	perk1Var2 int `json:"perk1Var2"`
	perk1Var3 int `json:"perk1Var3"`
	perk2	int `json:"perk2"`
	perk2Var1 int `json:"perk2Var1"`
	perk2Var2 int `json:"perk2Var2"`
	perk2Var3 int `json:"perk2Var3"`
	perk3	int `json:"perk3"`
	perk3Var1 int `json:"perk3Var1"`
	perk3Var2 int `json:"perk3Var2"`
	perk3Var3 int `json:"perk3Var3"`
	perk4	int `json:"perk4"`
	perk4Var1 int `json:"perk4Var1"`
	perk4Var2 int `json:"perk4Var2"`
	perk4Var3 int `json:"perk4Var3"`
	perk5	int `json:"perk5"`
	perk5Var1 int `json:"perk5Var1"`
	perk5Var2 int `json:"perk5Var2"`
	perk5Var3 int `json:"perk5Var3"`
	perkPrimaryStyle	int	`json:"perkPrimaryStyle"`
	perkSubStyle	int	`json:"perkSubStyle"`
	statPerk0	int	`json:"statPerk0"`
	statPerk1	int	`json:"statPerk1"`
	statPerk2	int	`json:"statPerk2"`
}

type TimelineDto struct {
	participantId	int	`json:"participantId"`
	csDiffPerMinDeltas	map[string]int `json:"csDiffPerMinDeltas"`
	damageTakenPerMinDeltas	map[string]int `json:"damageTakenPerMinDeltas"`
	role	string	`json:"role"`
	damageTakenDiffPerMinDeltas	map[string]int `json:"damageTakenDiffPerMinDeltas"`
	xpPerMinDeltas	map[string]int	`json:"xpPerMinDeltas"`
	xpDiffPerMinDeltas	map[string]int	`json:"xpDiffPerMinDeltas"`
	lane	string		`json:"lane"`
	creepsPerMinDeltas	map[string]int	`json:"creepsPerMinDeltas"`
	goldPerMinDeltas	map[string]int	`json:"goldPerMinDeltas"`
}

type MasteriesDto struct {
	rank int `json:"rank"`
	masteryId int `json:"masteryId"`
}