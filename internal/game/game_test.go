package game

import (
	"raw-fight-server/internal/fighter"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The tests cover the following scenarios:
// 1. The game should not end if both fighters have HP remaining.
// 2. The game should end if either fighter's HP reaches 0.
// 3. The game should end when the maximum number of rounds (MaxRounds) is reached.
//

// The tests also verify that the game correctly sets the game state (GameOver) and
// potentially determines the winner based on HP and stamina in different scenarios.
func TestGameStates(t *testing.T) {
	fighter1 := fighter.Fighter{
		Name:    "player1",
		HP:      100,
		Stamina: 500,
		BaseStats: fighter.Stats{
			Attack:  200,
			Defense: 100,
		},
	}

	fighter2 := fighter.Fighter{
		Name:    "player2",
		HP:      100,
		Stamina: 500,
		BaseStats: fighter.Stats{
			Attack:  200,
			Defense: 100,
		},
	}
	game := Game{
		Player:  &fighter1,
		Player2: &fighter2,
		Round:   1,
	}

	t.Run("Should not end the game if both fighters have HP", func(t *testing.T) {
		gameStatus := game.CheckGame()
		assert.Equal(t, gameStatus, GameOnGoing)
	})

	t.Run("Should end a game if the HP of either fighters are 0", func(t *testing.T) {
		game.Round = 10 // what if I remove this set, would keep game as 500??
		game.Player.HP = 0
		gameStatus := game.CheckGame()
		assert.Equal(t, game.Status, GameOver)
		assert.Equal(t, gameStatus, GameOver)
	})

	t.Run("Should end a game if the round is MaxRounds returning the fighter with most life", func(t *testing.T) {
		game.Round = MaxRounds
		gameStatus := game.CheckGame()
		assert.Equal(t, game.Status, GameOver)
		assert.Equal(t, gameStatus, GameOver)
	})

}
