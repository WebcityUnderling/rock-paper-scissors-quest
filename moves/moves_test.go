package moves_test

import (
	"phptogo/moves"
	"testing"
)

func TestPaperBeatsRock(t *testing.T) {
	paper, rock := moves.MoveIndices["Paper"], moves.MoveIndices["Rock"]
	result, _ := moves.GetMatchResult(paper, rock)
	if result != 1 {
		t.Errorf("GetMatchResult; %s should beat %s", moves.Moves[paper], moves.Moves[rock])
	}
}

func TestRockBeatsScissors(t *testing.T) {
	rock, scissors := moves.MoveIndices["Rock"], moves.MoveIndices["Scissors"]
	result, _ := moves.GetMatchResult(rock, scissors)
	if result != 1 {
		t.Errorf("GetMatchResult; %s should beat %s", moves.Moves[rock], moves.Moves[scissors])
	}
}

func TestScissorsBeatsPaper(t *testing.T) {
	scissors, paper := moves.MoveIndices["Scissors"], moves.MoveIndices["Paper"]
	result, _ := moves.GetMatchResult(scissors, paper)
	if result != 1 {
		t.Errorf("GetMatchResult; %s should beat %s", moves.Moves[scissors], moves.Moves[paper])
	}
}

func TestInvalidMoveResult(t *testing.T) {
	invalid := 99999
	result, err := moves.GetMatchResult(invalid, 0)
	if err == nil {
		t.Errorf("GetMatchResult; Invalid move should return an error")
	}
	if result != 0 {
		t.Errorf("GetMatchResult; Invalid move should return as 0")
	}
}
