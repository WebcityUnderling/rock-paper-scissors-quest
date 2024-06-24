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
)

// Level Setup
type Level struct {
	Number   int
	opponent *beastiary.Beast
	room     *rooms.Room
}

func createLevel(level int) Level {
	return Level{
		Number:   currentLevel,
		opponent: &beastiary.Beastiary[level],
		room:     &rooms.Dungeon[level],
	}
}

func ResetLevels() {
	currentLevel = 0
	replayLevel = false
}

var currentLevel = 0
var replayLevel = false
var levelCap = len(rooms.Dungeon)
var ErrInvalidOutcome = errors.New("invalid match outcome")

// Main Gameplay Loop
func Levels() {
	for {
		// If you surpass the last level, break, exit event
		if currentLevel >= levelCap {
			events.PrintExitEvent()
			break
		}
		level := createLevel(currentLevel)

		// Only execute enter room event if this is the first time you've entered a room
		if !replayLevel {
			events.PrintEnterRoomEvent(level.Number+1, level.opponent, level.room)
		}

		// Play a round of RPS
		result, err := level.play()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resulteventfunc, ok := moves.MoveOutcomeEvents[result]
		if ok {
			resulteventfunc(level.Number+1, level.opponent, level.room)
		}
		switch result {
		case 0:
			return
		case 1:
			currentLevel++
			replayLevel = false
		default:
			replayLevel = true
		}
	}
}

// RPS Logic
func (level Level) play() (int, error) {
	var playerMove string
	playerMove = moves.SelectMove()
	moveIndex := moves.MoveIndices[playerMove]

	difficultyFunc, ok := difficultyMoves[difficulty]
	if ok {
		return difficultyFunc(moveIndex, level), nil
	}
	return -1, ErrInvalidDifficulty
}

// Difficulty Settings
type DifficultyFunc func(moveIndex int, level Level) int

var difficultyMoves = map[string]DifficultyFunc{
	"Easy": func(moveIndex int, level Level) int {
		return moves.GetMatchResult(moveIndex, level.opponent.Attack)
	},
	"Medium": func(moveIndex int, level Level) int {
		return moves.GetMatchResult(moveIndex, moves.GetRandMove())
	},
	"Hard": func(moveIndex int, level Level) int {
		return moves.GetMatchResult(moveIndex, moves.GetWeightedMove(level.opponent.Attack))
	},
}

var difficulties = []string{"Easy", "Medium", "Hard"}
var difficulty = difficulties[0]
var ErrInvalidDifficulty = errors.New("invalid difficulty")

func SetDifficulty() {
	d, err := utils.SelectPrompt("Set Difficulty:", difficulties)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	difficulty = d
}
