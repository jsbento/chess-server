package types

type SearchGamesReq struct {
	Ids       *[]string `json:"ids"`
	PlayerId  *string   `json:"playerId"`
	PlayerIds *[]string `json:"playerIds"`
	Result    *string   `json:"result"`
	Limit     int       `json:"limit"`
	Offset    int       `json:"offset"`
}
