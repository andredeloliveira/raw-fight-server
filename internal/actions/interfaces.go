package actions

import (
	"raw-fight-server/internal/fighter"
	"raw-fight-server/internal/game"
)

type Actions interface {
	Execute(g *game.Game, attacker *fighter.Fighter, defender *fighter.Fighter) error
	GetStaminaCost(action string, f fighter.Fighter) error
	Register(g *game.Game) error
}
