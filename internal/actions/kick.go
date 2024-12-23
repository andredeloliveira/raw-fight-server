package actions

import (
	"raw-fight-server/internal/fighter"
	"raw-fight-server/internal/game"
)

type Kick struct {
	damageMultiplier float64
}

func NewKick(dm float64) Actions {
	return &Kick{
		damageMultiplier: dm,
	}
}

func (k *Kick) Execute(g *game.Game, attacker *fighter.Fighter, defender *fighter.Fighter) error {
	return nil
}

func (k *Kick) GetStaminaCost(action string, f fighter.Fighter) error {
	return nil
}

func (k *Kick) Register(g *game.Game) error {
	return nil
}
