package types

type Match struct {
	Id       string `json:"id"`
	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`
	DateTime string `json:"dateTime"`
	Stadium  string `json:"stadium"`
	Status   string `json:"status"`
}
