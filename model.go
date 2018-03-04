package r6

import "net/http"

const Endpoint = "https://api.r6stats.com/api/v1"

type client struct {
	httpClient http.Client
}

type PlayerResponse struct {
	Player Player `json:"player"`
}

type Player struct {
	Username string `json:"username"`
	Platform string `json:"platform"`
	UbisoftId string `json:"ubisoft_id"`
	IndexedAt string `json:"indexed_at"`
	UpdatedAt string `json:"updated_at"`
	Stats Stats `json:"stats"`
}

type Stats struct {
	Ranked GameTypeStats `json:"ranked"`
	Casual GameTypeStats `json:"casual"`
	Overall OverallStats `json:"overall"`
	Progression ProgressionStats `json:"progression"`
}

type GameTypeStats struct {
	HasPlayed bool `json:"has_played"`
	Wins int `json:"wins"`
	Losses int `json:"losses"`
	WinLossRatio float64 `json:"wlr"`
	Kills int `json:"kills"`
	Deaths int `json:"deaths"`
	KillDeathRatio float64 `json:"kd"`
	Playtime float64 `json:"playtime"`
}

type OverallStats struct {
	Revives int `json:"revives"`
	Suicides int `json:"suicides"`
	ReinforcementsDeployed int `json:"reinforcements_deployed"`
	BarricadesBuilt int `json:"barricades_built"`
	StepsMoved int `json:"steps_moved"`
	BulletsFired int `json:"bullets_fired"`
	Headshots int `json:"headshots"`
	MeleeKills int `json:"melee_kills"`
	PenetrationKills int `json:"penetration_kills"`
	Assists int `json:"assists"`
}

type ProgressionStats struct {
	Level int `json:"level"`
	Xp int `json:"xp"`
}
