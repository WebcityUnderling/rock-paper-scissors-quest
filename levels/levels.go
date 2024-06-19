package levels

import (
	"fmt"
	"phptogo/beastiary"
	"phptogo/events"
	"phptogo/moves"
	"phptogo/rooms"
)

type Level struct {
	Number   int
	opponent *beastiary.Beast
	room     *rooms.Room
}

var currentLevel = 0
var invalidMove = false

func Levels() {
	for {
		if currentLevel == 0 {
			events.IntroEvent()
			currentLevel++
		} else {
			if currentLevel > 10 {
				events.ExitEvent()
				break
			}

			level := createLevel(currentLevel)
			if !invalidMove {
				fmt.Println("\n\nLevel", level.Number, level.room.Name)
				fmt.Println(level.room.Enter)
				fmt.Println(level.opponent.EntryMessage)
			}
			result := level.play()

			switch result {
			case "Win!":
				events.TriumphEvent()
				fmt.Println("You defeated", level.opponent.Name)
				fmt.Println(level.opponent.DefeatMessage)
				fmt.Println(level.room.Leave)
				currentLevel++
			case "Lose!":
				events.DeathEvent(level.Number)
				fmt.Println("\n")
				fmt.Println(level.opponent.WinMessage)
				fmt.Println(level.room.Defeat)
				return

			case "error":
				invalidMove = true
			default:
				events.TieEvent()
				invalidMove = true
				fmt.Println("Go Again!")
			}
		}
	}
}

func (level Level) play() string {
	invalidMove = false
	var playerMove string
	fmt.Println("\n\nRock, Paper, Scissors.... SHOOT!")
	fmt.Scan(&playerMove)

	moveIndex, err := moves.GetMoveIndex(playerMove)
	if err != nil {
		fmt.Println(err.Error())
		return "error"
	}
	return moves.MatchResult(moveIndex, moves.GetRandMove())
}

func createLevel(level int) Level {
	beast := beastiary.GetBeastForLevel(level)
	return Level{
		Number:   currentLevel,
		opponent: beast,
		room:     &rooms.Dungeon[level-1],
	}
}
