package actions

import "math/rand/v2"

type Block struct{}
type Dodge struct{}
type Jump struct{}
type Duck struct{}

type Punch struct {
	damageMultiplier float64
}

func CalculateBaseDamage(attackStat int) int {
	minDamage := int(float64(attackStat) * 0.5)
	maxDamage := attackStat

	if minDamage < 1 {
		minDamage = 1
	}

	return rand.IntN(maxDamage-minDamage+1) + minDamage

}

func CalculateActualDamage(modifiedDamage float64, defenseStat int) int {
	damageReduction := float64(defenseStat) * 0.10
	actualDamage := modifiedDamage * (1 - damageReduction)

	if actualDamage < 1 {
		actualDamage = 1
	}

	return int(actualDamage)
}
