package levels_test

import (
	"fmt"
	"rpsq/levels"
	"sync"
	"testing"
)

type MockPrompter struct {
	Response string
	Err      error
}

func (mp MockPrompter) SelectPrompt(prompt string, options []string) (string, error) {
	return mp.Response, mp.Err
}
func TestInvalidDifficultyResultsInInitialDifficulty(t *testing.T) {
	mp := MockPrompter{Response: "invalid", Err: levels.ErrInvalidDifficulty}
	levels.SetDifficulty(mp)
	if levels.Difficulty != levels.Difficulties[0] {
		fmt.Printf("\nSetDifficulty: invalid selection should result in %s, got %s", levels.Difficulties[0], levels.Difficulty)
	}
}

func TestAllLevelsCanBeCreated(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	levels.CreateLevels(&wg)
	wg.Wait()
	if levels.LevelCap != len(levels.AllLevels) {
		t.Errorf("CreateLevels; Level amount mismatch, should have created %d levels, created %d", levels.LevelCap, len(levels.AllLevels))
	}
}
