package fighter

import (
	"math/rand"
)

type ActionType string

const (
	Kick      ActionType = "kick"
	Punch     ActionType = "punch"
	Block     ActionType = "block"
	DodgeBack ActionType = "dodge_back"
	Jump      ActionType = "jump"
	Duck      ActionType = "duck"

	BaseStaminaRecover = 10
)

type Stats struct {
	Attack  int
	Defense int
	Dodge   int
}

type Fighter struct {
	Name      string
	HP        int
	Stamina   int
	BaseStats Stats
}

func (f *Fighter) RecoverStamina() int {
	if f.Stamina > 100 {
		f.Stamina = 100
		return 0
	}

	f.Stamina = f.Stamina + rand.Intn(BaseStaminaRecover)
	return f.Stamina
}
