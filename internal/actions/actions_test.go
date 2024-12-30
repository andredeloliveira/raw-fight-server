package actions

import (
	"raw-fight-server/internal/fighter"
	"raw-fight-server/internal/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDamages(t *testing.T) {
	t.Run("Should calculate base damage", func(t *testing.T) {
		attackStat := 2
		damage := CalculateBaseDamage(attackStat)
		assert.GreaterOrEqual(t, damage, 1)
	})

	t.Run("Should calculate actual damage", func(t *testing.T) {
		modifiedDamage := 30.5
		defenseStat := 1
		damage := CalculateActualDamage(modifiedDamage, defenseStat)
		assert.GreaterOrEqual(t, damage, 3)
		// this is newly saved
	})
}

func TestKick(t *testing.T) {

	damageMultiplier := 5.
	attacker := fighter.Fighter{
		Name:    "Fighter 1",
		HP:      80,
		Stamina: 80,
		BaseStats: fighter.Stats{
			Attack:  50,
			Defense: 50,
			Dodge:   25,
		},
	}
	defender := fighter.Fighter{
		Name:    "Fighter 2",
		HP:      80,
		Stamina: 80,
		BaseStats: fighter.Stats{
			Attack:  40,
			Defense: 40,
			Dodge:   20,
		},
	}
	game := game.Game{
		Player:  &attacker,
		Player2: &defender,
		Round:   1,
	}

	kick := NewKick(damageMultiplier)

	t.Run("Should Execute Kick without any problems", func(t *testing.T) {
		err := kick.Execute(&game, &attacker, &defender)
		assert.Equal(t, err, nil)
		assert.LessOrEqual(t, defender.HP, 80)
	})

	t.Run("Should not Kick if there is no stamina", func(t *testing.T) {
		attacker.Stamina = 0
		defenderHPBeforeKick := defender.HP
		err := kick.Execute(&game, &attacker, &defender)
		assert.Equal(t, err.Error(), "Not enough stamina")
		assert.Equal(t, defender.HP, defenderHPBeforeKick)
	})
}
