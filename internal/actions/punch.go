package actions

import (
	"errors"
	"raw-fight-server/internal/fighter"
	"raw-fight-server/internal/game"
)

const (
	PunchStaminaCost = 8
)

type Punch struct {
	damageMultiplier float64
}

func NewPunch(dm float64) Actions {
	return &Punch{
		damageMultiplier: dm,
	}
}

func (p *Punch) Error(msg string) error {
	return errors.New(msg)
}

func (p *Punch) Execute(
	g *game.Game,
	attacker *fighter.Fighter,
	defender *fighter.Fighter,
) error {
	if attacker.Stamina <= 0 || attacker.Stamina < p.GetStaminaCost() {
		return p.Error("Not enough stamina")
	}
	baseDamage := CalculateBaseDamage(attacker.BaseStats.Attack)
	modifiedDamage := float64(baseDamage) * p.damageMultiplier
	actualDamage := CalculateActualDamage(modifiedDamage, defender.BaseStats.Defense)
	defender.HP -= actualDamage
	attacker.Stamina -= p.GetStaminaCost()

	return nil
}

func (p *Punch) GetStaminaCost() int {
	return PunchStaminaCost
}

func (p *Punch) Register(g *game.Game) error {
	return nil
}
