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

func setupGame() (fighter.Fighter, fighter.Fighter, game.Game) {
	fighter1 := fighter.Fighter{
		Name:    "Fighter 1",
		HP:      80,
		Stamina: 80,
		BaseStats: fighter.Stats{
			Attack:  50,
			Defense: 50,
		},
	}
	fighter2 := fighter.Fighter{
		Name:    "Fighter 2",
		HP:      80,
		Stamina: 80,
		BaseStats: fighter.Stats{
			Attack:  40,
			Defense: 40,
		},
	}
	game := game.Game{
		Player:  &fighter1,
		Player2: &fighter2,
		Round:   1,
	}
	return fighter1, fighter2, game
}

func TestKick(t *testing.T) {
	damageMultiplier := 5.
	attacker, defender, game := setupGame()
	kick := NewKick(damageMultiplier)

	t.Run("Should Execute Kick without any problems", func(t *testing.T) {
		defenderHPBeforeKick := defender.HP
		attackerStaminaBeforeKick := attacker.Stamina
		err := kick.Execute(&game, &attacker, &defender)
		assert.Equal(t, err, nil)
		assert.Less(t, defender.HP, defenderHPBeforeKick)
		assert.Less(t, attacker.Stamina, attackerStaminaBeforeKick)
	})

	t.Run("Should not Kick if there is no stamina", func(t *testing.T) {
		attacker.Stamina = 4
		defenderHPBeforeKick := defender.HP
		err := kick.Execute(&game, &attacker, &defender)
		assert.Equal(t, err.Error(), "Not enough stamina")
		assert.Equal(t, defender.HP, defenderHPBeforeKick)
	})

	t.Run("Should test stamina cost", func(t *testing.T) {
		staminaCost := kick.GetStaminaCost()
		// Stamina cost of kick is 10
		assert.Equal(t, staminaCost, 10)
	})
}

func TestPunch(t *testing.T) {
	attacker, defender, game := setupGame()
	damageMultiplier := 10.
	punch := NewPunch(damageMultiplier)

	t.Run("Should punch without any problems", func(t *testing.T) {
		defenderHPBeforePunch := defender.HP
		attackerStaminaBeforePunch := attacker.Stamina
		err := punch.Execute(&game, &attacker, &defender)
		assert.Equal(t, err, nil)
		assert.Less(t, defender.HP, defenderHPBeforePunch)
		assert.Less(t, attacker.Stamina, attackerStaminaBeforePunch)
	})

	t.Run("Should not punch if there is no stamina", func(t *testing.T) {
		defenderHPBeforePunch := defender.HP
		attacker.Stamina = 5
		err := punch.Execute(&game, &attacker, &defender)
		assert.Equal(t, err.Error(), "Not enough stamina")
		assert.Equal(t, defender.HP, defenderHPBeforePunch)
	})
}

func TestBlock(t *testing.T) {
	attacker, defender, game := setupGame()
	block := NewBlock()

	t.Run("Should block without any problems", func(t *testing.T) {
		defenderHPBeforeBlock := defender.HP
		err := block.Execute(&game, &attacker, &defender)
		assert.Equal(t, err, nil)
		assert.Equal(t, defenderHPBeforeBlock, defender.HP)
	})

}
