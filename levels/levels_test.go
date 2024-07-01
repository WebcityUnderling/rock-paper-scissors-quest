package levels_test

import (
	"fmt"
	"phptogo/levels"
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
