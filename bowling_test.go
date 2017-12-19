package bowling

import "testing"

func TestScoreGameStringMustContainNineFrameBoundariesAndOneBonusBallBoundary(t *testing.T) {
	games := []struct {
		input string
		valid bool
	}{
		{"", false},
		{"|", false},
		{"--|--|--|--|--|--|--|--|--||", false},
		{"--|--|--|--|--|--|--|--|--|--|--||", false},
		{"--|--|--|--|--|--|--|--|--|--||", true},
		{"X|X|X|X|X|X|X|X|X|X||XX", true},
	}

	for _, game := range games {
		_, err := NewGame(game.input)
		valid := err == nil
		if valid != game.valid {
			t.Errorf("incorrect validity for %v, got %v, want %v", game.input, valid, game.valid)
		}
	}
}

func TestScoreGameStringMustContainOnlyValidCharacters(t *testing.T) {
	games := []struct {
		input string
		valid bool
	}{
		{"X|X|X|X|X|X|X|X|X|X||XX", true},
		{"9-|9-|9-|9-|9-|9-|9-|9-|9-|9-||", true},
		{"5/|5/|5/|5/|5/|5/|5/|5/|5/|5/||5", true},
		{"X|7/|9-|X|-8|8/|-6|X|X|X||81", true},
		{"X|X|X|X|X|X|X|?|X|X||XX", false},
		{"X|X|X|+|--|X|X|X|X|X||XX", false},
	}

	for _, game := range games {
		_, err := NewGame(game.input)
		valid := err == nil
		if valid != game.valid {
			t.Errorf("incorrect validity for %v, got %v, want %v", game.input, valid, game.valid)
		}
	}
}

func TestScoreGameWithAllMissesReturnsZero(t *testing.T) {
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

func TestScoreGameWithAllOnesReturnsTwenty(t *testing.T) {
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

func TestScoreGameWithAllMissesAndOnesReturnsTen(t *testing.T) {
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

func TestScoreGameWithAllNinesAndMissesReturnsNinety(t *testing.T) {
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

func TestScoreGameWithAllFivesAndSparesReturnsOneHundredAndFifty(t *testing.T) {
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
