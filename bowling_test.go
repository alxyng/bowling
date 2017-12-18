package bowling

import "testing"

func TestScoreGameStringMustContainNineFrameBoundariesAndOneBonusBallBoundary(t *testing.T) {
	games := []struct {
		game  string
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
		valid := Score(game.game) >= 0
		if valid != game.valid {
			t.Errorf("incorrect validity for %v, got %v, want %v", game.game, valid, game.valid)
		}
	}
}

func TestScoreGameStringMustContainOnlyValidCharacters(t *testing.T) {
	games := []struct {
		game  string
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
		valid := Score(game.game) >= 0
		if valid != game.valid {
			t.Errorf("incorrect validity for %v, got %v, want %v", game.game, valid, game.valid)
		}
	}
}

func TestScoreGameWithAllMissesReturnsZero(t *testing.T) {
	s := Score("--|--|--|--|--|--|--|--|--|--||")

	expectedScore := 0
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreGameWithAllOnesReturnsTwenty(t *testing.T) {
	s := Score("11|11|11|11|11|11|11|11|11|11||")

	expectedScore := 20
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreGameWithAllMissesAndOnesReturnsTen(t *testing.T) {
	s := Score("-1|-1|-1|-1|-1|-1|-1|-1|-1|-1||")

	expectedScore := 10
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreGameWithAllNinesAndMissesReturnsNinety(t *testing.T) {
	s := Score("9-|9-|9-|9-|9-|9-|9-|9-|9-|9-||")

	expectedScore := 90
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreGameWithAllFivesAndSparesReturnsOneHundredAndFifty(t *testing.T) {
	s := Score("5/|5/|5/|5/|5/|5/|5/|5/|5/|5/||5")

	expectedScore := 150
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}

func TestScoreWithAllStrikes(t *testing.T) {
	s := Score("X|X|X|X|X|X|X|X|X|X||XX")

	expectedScore := 300
	if s != expectedScore {
		t.Errorf("incorrect score, got %v, want %v", s, expectedScore)
	}
}
