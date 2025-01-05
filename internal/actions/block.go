package actions

import (
	"raw-fight-server/internal/fighter"
	"raw-fight-server/internal/game"

	"github.com/pkg/errors"
)

const (
	BlockStaminaCost = 4
)

type Block struct{}

func NewBlock() Actions {
	return &Block{}
}

func (b *Block) Error(msg string) error {
	return errors.New(msg)
}

func (b *Block) Execute(g *game.Game, attacker *fighter.Fighter, defender *fighter.Fighter) error {
	if defender.Stamina <= 0 || defender.Stamina < b.GetStaminaCost() {
		return b.Error("Not enough stamina")
	}
	defender.Stamina = b.GetStaminaCost()
	return nil
}

func (b *Block) GetStaminaCost() int {
	return BlockStaminaCost
}

func (b *Block) Register(g *game.Game) error {
	return nil
}
