package matches

type MatchDto struct {
	GameID                int                        `json:"gameId"`
	PlatformID            string                     `json:"platformId"`
	GameCreation          int                        `json:"gameCreation"`
	GameDuration          int                        `json:"gameDuration"`
	QueueID               int                        `json:"queueId"`
	MapID                 int                        `json:"mapId"`
	SeasonID              int                        `json:"seasonId"`
	GameVersion           string                     `json:"gameVersion"`
	GameMode              string                     `json:"gameMode"`
	GameType              string                     `json:"gameType"`
	Teams                 []TeamsDto                 `json:"teams"`
	Participants          []ParticipantsDto          `json:"participants"`
	ParticipantIdentities []ParticipantIdentitiesDto `json:"participantIdentities"`
}
