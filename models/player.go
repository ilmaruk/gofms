package models

type Player struct {
	Name            string `json:"name"`
	Age             int    `json:"age"`
	Nationality     string `json:"nationality"`
	PreferredSide   string `json:"preferredSide"`
	Stopping        int    `json:"stopping"`
	Tackling        int    `json:"tackling"`
	Passing         int    `json:"passing"`
	Shooting        int    `json:"shooting"`
	Stamina         int    `json:"stamina"`
	Aggression      int    `json:"aggression"`
	StoppingAbility int    `json:"kab"`
	TacklingAbility int    `json:"tab"`
	PassingAbility  int    `json:"pab"`
	ShootingAbility int    `json:"sab"`
	Gam             int    `json:"gam"`
	Sav             int    `json:"sav"`
	Ktk             int    `json:"ktk"`
	Kps             int    `json:"kps"`
	Sht             int    `json:"sht"`
	Gls             int    `json:"gls"`
	Ass             int    `json:"ass"`
	DP              int    `json:"dp"`
	Inj             int    `json:"inj"`
	Sus             int    `json:"sus"`
	Fitness         int    `json:"fitness"`
}

// Simply setting defaults.
func NewPlayer() Player {
	return Player{StoppingAbility: 300, TacklingAbility: 300, PassingAbility: 300, ShootingAbility: 300, Fitness: 100}
}
