package game

import "raw-fight-server/internal/fighter"

const (
	MaxRounds = 500
)

type GameState int

// iota does a magic where all the consts
// start from 0 to ...n (the number of const)
const (
	GameOver GameState = iota
	GameOnGoing
)

type Game struct {
	Player  *fighter.Fighter
	Player2 *fighter.Fighter
	Round   int
	Status  GameState
	Winner  *fighter.Fighter
}

func (g *Game) CheckGame() GameState {
	if g.Round == MaxRounds {
		g.Status = GameOver
		return GameOver
	}

	switch {
	case g.Player.HP == 0 && g.Player2.HP > 0:
		g.Winner = g.Player
	case g.Player.HP == 0 && g.Player2.HP > 0:
		g.Winner = g.Player2
	default:
		return GameOnGoing
	}
	g.Status = GameOver
	return GameOver
}
