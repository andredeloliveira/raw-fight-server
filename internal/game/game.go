package game

import "raw-fight-server/internal/fighter"

type Game struct {
	Player  *fighter.Fighter
	Player2 *fighter.Fighter
	Round   int
}
