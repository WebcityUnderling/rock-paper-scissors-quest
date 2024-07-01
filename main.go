package main

import (
	"rpsq/levels"
	"rpsq/utils"
)

var replayResponsesMap = map[string]bool{"Yes": true, "No": false}
var replayResponses = []string{"Yes", "No"}

func main() {
	events.SetEventTimeout()
	// Boot screen
	events.PrintIntroEvent()
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
