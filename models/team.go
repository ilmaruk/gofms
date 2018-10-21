package models

type Team struct {
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

func (t Team) GetPlayer(i int) Player {
	return t.Players[i]
}

func (t *Team) AddPlayer(p Player) {
	t.Players = append(t.Players, p)
}
