package moves

import (
	"errors"
	"fmt"
	"math/rand"
)

var Moves = [3]string{"Rock", "Paper", "Scissors"}

func GetRandMove() int {
	move := rand.Intn(len(Moves))
	return move
}

func GetMoveIndex(move string) (int, error) {
	for i := range Moves {
		if Moves[i] == move {
			return i, nil
		}
	}
	return -1, errors.New("that is not a valid move, Try Rock Paper Or Scissors.")
}

func MatchResult(player int, opponent int) string {
	fmt.Println(Moves[player], " vs ", Moves[opponent])

	if player == opponent {
		return "Tie!"
	}
	if opponent < 2 && player == opponent+1 || opponent == 2 && player == 0 {
		return "Win!"
	}

	return "Lose!"
}
