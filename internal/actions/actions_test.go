package actions

import (
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
	})
}
