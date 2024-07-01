package levels

import (
	"errors"
	"fmt"
	"rpsq/beastiary"
	"rpsq/events"
	"rpsq/moves"
	"rpsq/rooms"
	"rpsq/utils"
	"sync"
)

// Level Setup
var currentLevel = 0
var replayLevel = false
var levelCap = len(rooms.Dungeon)
var ErrInvalidOutcome = errors.New("invalid match outcome")

type Level struct {
	Number   int
	opponent *beastiary.Beast
	room     *rooms.Room
}

var allLevels []Level

func CreateLevels(wg *sync.WaitGroup) error {
	defer wg.Done()
	wg.Add(1)
	levels := []Level{}

	for i := 0; i < levelCap; i++ {
		levels = append(levels, createLevel(i))
	}
	allLevels = levels
	return nil
}

func createLevel(level int) Level {
	return Level{
		Number:   level,
		opponent: &beastiary.Beastiary[level],
		room:     &rooms.Dungeon[level],
	}
}

func ResetLevels() {
	currentLevel = 0
	replayLevel = false
}

// Main Gameplay Loop
func Levels() {
	for {
		// If you surpass the last level, break, exit event
		if currentLevel >= levelCap {
			events.PrintExitEvent()
			break
		}
		level := allLevels[currentLevel]

		// Only execute enter room event if this is the first time you've entered a room
		if !replayLevel {
			events.PrintEnterRoomEvent(level.Number+1, level.opponent, level.room)
		}

		// Play a round of RPS
		result, err := level.play()
		if err != nil {
			fmt.Println(err, "You cannot escape, Your opponent is waiting...")
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
	playerMove, err := moves.SelectMove()
	if err != nil {
		return -1, err
	}
	moveIndex := moves.MoveIndices[playerMove]

	difficultyFunc, ok := difficultyMoves[Difficulty]
	if ok {
		return difficultyFunc(moveIndex, level), nil
	}
	return -1, ErrInvalidDifficulty
}

// Difficulty Settings
type DifficultyFunc func(moveIndex int, level Level) int

var difficultyMoves = map[string]DifficultyFunc{
	"Easy": func(moveIndex int, level Level) int {
		result, _ := moves.GetMatchResult(moveIndex, level.opponent.Attack)
		return result
	},
	"Medium": func(moveIndex int, level Level) int {
		result, _ := moves.GetMatchResult(moveIndex, moves.GetRandMove())
		return result
	},
	"Hard": func(moveIndex int, level Level) int {
		result, _ := moves.GetMatchResult(moveIndex, moves.GetWeightedMove(level.opponent.Attack))
		return result
	},
}

var Difficulties = []string{"Easy", "Medium", "Hard"}
var Difficulty = Difficulties[0]
var ErrInvalidDifficulty = errors.New("invalid difficulty")

func SetDifficulty(prompter utils.Prompter) {
	d, err := prompter.SelectPrompt("Select Difficulty", Difficulties)
	if err != nil {
		fmt.Printf("Error setting level, defaulting to %s", Difficulty)
	} else {
		Difficulty = d
	}
}
