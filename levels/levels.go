package levels

import (
	"errors"
	"fmt"
	"os"
	"phptogo/beastiary"
	"phptogo/events"
	"phptogo/moves"
	"phptogo/rooms"
	"phptogo/utils"
	"time"
)

type Level struct {
	Number   int
	opponent *beastiary.Beast
	room     *rooms.Room
}

var ErrInvalidOutcome = errors.New("invalid outcome")

var currentLevel = 1
var replayLevel = false

func Levels() {
	for {

		if currentLevel > 10 {
			events.ExitEvent()
			break
		}

		level := createLevel(currentLevel)
		if !replayLevel {
			fmt.Println("\n\nLevel", level.Number, level.room.Name)
			time.Sleep(2 * time.Second)
			fmt.Println(level.room.Enter)
			time.Sleep(2 * time.Second)
			fmt.Println(level.opponent.EntryMessage)
			time.Sleep(2 * time.Second)
		}
		result := level.play()

		switch result {
		case 2:
			events.TieEvent()
			replayLevel = true
			fmt.Println("Go Again!")
		case 1:
			events.TriumphEvent()
			fmt.Println("You defeated", level.opponent.Name)
			time.Sleep(2 * time.Second)
			fmt.Println(level.opponent.DefeatMessage)
			time.Sleep(2 * time.Second)
			fmt.Println(level.room.Leave)
			time.Sleep(2 * time.Second)
			currentLevel++
		case 0:
			events.DeathEvent(level.Number)
			time.Sleep(500 * time.Millisecond)
			fmt.Println("You Died. You got to level", level.Number)
			time.Sleep(2 * time.Second)
			fmt.Println(level.opponent.WinMessage)
			time.Sleep(2 * time.Second)
			fmt.Println(level.room.Defeat)
			return
		default:
			// Something has gone terribly wrong
			fmt.Println(ErrInvalidOutcome.Error())
			os.Exit(1)
		}

	}
}

func (level Level) play() int {
	replayLevel = false
	var playerMove string
	fmt.Println("\n\nRock, Paper, Scissors.... SHOOT!")
	time.Sleep(1 * time.Second)
	playerMove = moves.ChooseMove()

	moveIndex, err := moves.GetMoveIndex(playerMove)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	switch difficulty {
	case "Medium":
		return moves.MatchResult(moveIndex, moves.GetRandMove())
	case "Hard":
		return moves.MatchResult(moveIndex, moves.WeightedMove(level.opponent.Attack))
	default:
		return moves.MatchResult(moveIndex, level.opponent.Attack)
	}
}

func createLevel(level int) Level {
	beast := &beastiary.Beastiary[level]
	return Level{
		Number:   currentLevel,
		opponent: beast,
		room:     &rooms.Dungeon[level-1],
	}
}

// Difficulty
var difficulty = "Easy"

func SetDifficulty() {
	d, err := utils.SelectPrompt("Set Difficulty:", []string{
		"Easy",
		"Medium",
		"Hard",
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	difficulty = d
}
