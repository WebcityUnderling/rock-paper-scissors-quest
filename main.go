package main

import (
	"phptogo/events"
	"phptogo/levels"
	"phptogo/utils"
)

var initialised = false
var replayResponsesMap = map[string]bool{"Yes": true, "No": false}
var replayResponses = []string{"Yes", "No"}

func main() {
	// Boot screen
	events.PrintIntroEvent()
	// Gameplay loop
	for {
		levels.SetDifficulty()
		levels.Levels()
		replay, _ := utils.SelectPrompt("Play Again?", replayResponses)
		if !replayResponsesMap[replay] {
			break
		}
		levels.ResetLevels()
	}
}
