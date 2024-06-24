package moves

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"phptogo/utils"
	"time"
)

var MoveIndices = map[string]int{"Rock": 0, "Paper": 1, "Scissors": 2}
var Moves = []string{"Rock", "Paper", "Scissors"}

func GetRandMove() int {
	return rand.Intn(len(Moves))
}

func WeightedMove(attack int) int {
	weight := rand.Intn(2)
	if weight == 0 {
		return attack
	}
	return GetRandMove()
}

func GetMoveIndex(move string) (int, error) {
	if index, ok := MoveIndices[move]; ok {
		return index, nil
	}
	return -1, errors.New("invalid move")
}

func MatchResult(playerMove int, opponentMove int) int {
	fmt.Println(Moves[playerMove], " vs ", Moves[opponentMove])

	if playerMove == opponentMove {
		return 2
	}
	if opponentMove < 2 && playerMove == opponentMove+1 || opponentMove == 2 && playerMove == 0 {
		return 1
	}

	return 0
}

func ChooseMove() string {
	time.Sleep(1 * time.Second)
	move, err := utils.SelectPrompt("Choose your weapon...", Moves)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return move
}
