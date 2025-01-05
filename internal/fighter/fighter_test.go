package fighter

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestFighter(t *testing.T) {

	fighter := Fighter{
		Name:    "Fighter1",
		HP:      100,
		Stamina: 80,
		BaseStats: Stats{
			Attack:  30,
			Defense: 40,
		},
	}

	t.Run("Test recover stamina", func(t *testing.T) {
		fighter.RecoverStamina()
		assert.Assert(t, fighter.Stamina > 80 && fighter.Stamina <= 90)
	})
}
