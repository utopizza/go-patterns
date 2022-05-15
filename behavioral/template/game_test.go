package template

import "testing"

func TestGame(t *testing.T) {
	gameA := &GameA{
		Name: "NBA2k12",
	}
	gameB := &GameB{
		Name: "FIFA",
	}

	gamePlayer := &GamePlayer{}
	gamePlayer.Game = gameA
	gamePlayer.Play()

	gamePlayer.Game = gameB
	gamePlayer.Play()
}
