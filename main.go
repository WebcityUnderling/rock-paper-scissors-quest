package main

import (
	"rpsq/events"
	"rpsq/levels"
	"rpsq/utils"
	"sync"
)

var replayResponsesMap = map[string]bool{"Yes": true, "No": false}
var replayResponses = []string{"Yes", "No"}

func main() {
	events.SetEventTimeout()
	wg := sync.WaitGroup{}

	go events.PrintIntroEvent(&wg)
	go levels.CreateLevels(&wg) // Boot screen

	wg.Wait()

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
