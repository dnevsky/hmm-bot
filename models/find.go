package models

type Find struct {
	ID      int      `json:"id"`
	LS      int      `json:"ls"`
	SF      int      `json:"sf"`
	LV      int      `json:"lv"`
	Players []string `json:"players"`
	Client  string   `json:"client"`
	IP      string   `json:"ip"`
	TS      int      `json:"ts"`
}
