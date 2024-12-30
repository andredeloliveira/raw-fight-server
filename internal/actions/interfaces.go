package actions

import (
	"raw-fight-server/internal/fighter"
	"raw-fight-server/internal/game"
)

type Actions interface {
	Execute(g *game.Game, attacker *fighter.Fighter, defender *fighter.Fighter) error
	GetStaminaCost() int
	Register(g *game.Game) error
	Error(msg string) error
}
