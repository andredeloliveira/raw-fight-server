package actions

import (
	"raw-fight-server/internal/fighter"
	"raw-fight-server/internal/game"

	"github.com/pkg/errors"
)

type Kick struct {
	damageMultiplier float64
}

func NewKick(dm float64) Actions {
	return &Kick{
		damageMultiplier: dm,
	}
}

func (k *Kick) Error(msg string) error {
	return errors.New(msg)
}

func (k *Kick) Execute(
	g *game.Game,
	attacker *fighter.Fighter,
	defender *fighter.Fighter,
) error {
	if attacker.Stamina <= 0 {
		return k.Error("Not enough stamina")
	}
	baseDamage := CalculateBaseDamage(attacker.BaseStats.Attack)
	modifiedDamage := float64(baseDamage) * k.damageMultiplier
	actualDamage := CalculateActualDamage(modifiedDamage, defender.BaseStats.Defense)
	defender.HP -= actualDamage
	attacker.Stamina -= k.GetStaminaCost()
	return nil
}

// Kick will have a fixed stamina cost
func (k *Kick) GetStaminaCost() int {
	return 10
}

func (k *Kick) Register(g *game.Game) error {
	return nil
}
