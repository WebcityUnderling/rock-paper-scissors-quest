package main

import (
	"phptogo/levels"
	"phptogo/utils"
)

var replayResponsesMap = map[string]bool{"Yes": true, "No": false}
var replayResponses = []string{"Yes", "No"}

func main() {
	// Boot screen
	// events.PrintIntroEvent()
	// Gameplay loop
	for {
		levels.SetDifficulty(utils.LivePrompter{})
		levels.Levels()
		replay, _ := utils.Prompter.SelectPrompt(utils.LivePrompter{}, "Play again?", replayResponses)
		if !replayResponsesMap[replay] {
			break
		}
		levels.ResetLevels()
	}
}
