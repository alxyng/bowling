package bowling

import "testing"

func TestNewGameInputThatDoesNotContainNineFrameBoundariesAndOneBonusBallsBoundaryReturnsANonNilError(t *testing.T) {
	inputs := []string{
		"",
		"|",
		"--|--|--|--|--|--|--|--|--||",
		"--|--|--|--|--|--|--|--|--|--|--||",
	}

	for _, input := range inputs {
		_, err := NewGame(input)
		if err == nil {
			t.Errorf("expected an error for %v but got nil", input)
		}
	}
}

func TestNewGameInputThatContainsNineFrameBoundariesAndOneBonusBallsBoundaryReturnsANilError(t *testing.T) {
	inputs := []string{
		"--|--|--|--|--|--|--|--|--|--||",
		"X|X|X|X|X|X|X|X|X|X||XX",
	}

	for _, input := range inputs {
		_, err := NewGame(input)
		if err != nil {
			t.Errorf("expected a nil error for %v but got %v", input, err)
		}
	}
}

func TestNewGameInputThatContainsInvalidCharactersReturnsANonNilError(t *testing.T) {
	inputs := []string{
		"X|X|X|X|?|X|X|X|X|X||XX",
		"X|X|X|X||X|X|X|X|X||XX",
		"X|X|X|X|---|X|X|X|X|X||XX",
		"X|X|X|X|+|X|X|X|X|X||XX",
		"X|X|X|X|-.|X|X|X|X|X||XX",
		"X|X|X|X|1.|X|X|X|X|X||XX",
	}

	for _, input := range inputs {
		_, err := NewGame(input)
		if err == nil {
			t.Errorf("expected an error for %v but got nil", input)
		}
	}
}

func TestNewGameInputThatContainsValidCharactersReturnsANilError(t *testing.T) {
	inputs := []string{
		"9-|9-|9-|9-|9-|9-|9-|9-|9-|9-||",
		"5/|5/|5/|5/|5/|5/|5/|5/|5/|5/||5",
		"X|7/|9-|X|-8|8/|-6|X|X|X||81",
	}

	for _, input := range inputs {
		_, err := NewGame(input)
		if err != nil {
			t.Errorf("expected a nil error for %v but got %v", input, err)
		}
	}
}

func TestScoreWithAllMissesReturnsZero(t *testing.T) {
	game, err := NewGame("--|--|--|--|--|--|--|--|--|--||")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	s := game.Score()
	expectedScore := 0
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreWithAllOnesReturnsTwenty(t *testing.T) {
	game, err := NewGame("11|11|11|11|11|11|11|11|11|11||")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	s := game.Score()
	expectedScore := 20
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreWithAllMissesAndOnesReturnsTen(t *testing.T) {
	game, err := NewGame("-1|-1|-1|-1|-1|-1|-1|-1|-1|-1||")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	s := game.Score()
	expectedScore := 10
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreWithAllNinesAndMissesReturnsNinety(t *testing.T) {
	game, err := NewGame("9-|9-|9-|9-|9-|9-|9-|9-|9-|9-||")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	s := game.Score()
	expectedScore := 90
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreWithAllFivesAndSparesReturnsOneHundredAndFifty(t *testing.T) {
	game, err := NewGame("5/|5/|5/|5/|5/|5/|5/|5/|5/|5/||5")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	s := game.Score()
	expectedScore := 150
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreWithAllStrikes(t *testing.T) {
	game, err := NewGame("X|X|X|X|X|X|X|X|X|X||XX")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	s := game.Score()
	expectedScore := 300
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreWithVariedFrameScores(t *testing.T) {
	game, err := NewGame("X|7/|9-|X|-8|8/|-6|X|X|X||81")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}

	s := game.Score()
	expectedScore := 167
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func BenchmarkNewGameWithVariedFrameScores(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewGame("X|7/|9-|X|-8|8/|-6|X|X|X||81")
	}
}

var score int

func BenchmarkScoreWithVariedFrameScores(b *testing.B) {
	game, _ := NewGame("X|7/|9-|X|-8|8/|-6|X|X|X||81")
	var s int

	for n := 0; n < b.N; n++ {
		s = game.Score()
	}

	score = s
}
