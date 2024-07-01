package moves

import (
	"errors"
	"fmt"
	"math/rand"
	"rpsq/beastiary"
	"rpsq/events"
	"rpsq/rooms"
	"rpsq/utils"
	"time"
)

type moveEventFunc func(level int, opponent *beastiary.Beast, room *rooms.Room)

var MoveIndices = map[string]int{"Rock": 0, "Paper": 1, "Scissors": 2}
var Moves = []string{"Rock", "Paper", "Scissors"}
var ErrInvalidMove = errors.New("invalid move")

var MoveOutcomeEvents = map[int]moveEventFunc{
	0: func(level int, opponent *beastiary.Beast, room *rooms.Room) {
		events.PrintDeathEvent(level, opponent, room)
	},
	1: func(level int, opponent *beastiary.Beast, room *rooms.Room) {
		events.PrintTriumphEvent(opponent, room)

	},
	2: func(level int, opponent *beastiary.Beast, room *rooms.Room) {
		events.PrintTieEvent()
	},
}

func GetRandMove() int {
	return rand.Intn(len(Moves))
}

func GetWeightedMove(attack int) int {
	weight := rand.Intn(2)
	if weight == 0 {
		return attack
	}
	return GetRandMove()
}

func GetMatchResult(playerMove int, opponentMove int) (int, error) {
	//check move set length can include move pointers
	if playerMove > len(Moves)-1 || opponentMove > len(Moves)-1 {
		return 0, ErrInvalidMove
	}

	fmt.Println(Moves[playerMove], " vs ", Moves[opponentMove])

	if playerMove == opponentMove {
		return 2, nil
	}
	if opponentMove < 2 && playerMove == opponentMove+1 || opponentMove == 2 && playerMove == 0 {
		return 1, nil
	}

	return 0, nil
}

func SelectMove() (string, error) {
	fmt.Println("\nRock, Paper, Scissors.... SHOOT!")
	time.Sleep(1 * time.Second)
	move, err := utils.Prompter.SelectPrompt(utils.LivePrompter{}, "Choose your weapon...", Moves)
	if err != nil {
		return "error", err
	}
	return move, nil
}
