package bowling

import (
	"fmt"
	"strconv"
	"strings"
)

// Game represents a game of Ten-Pin Bowling. It is an
// implementation of the "Bowling Game" problem from
// cyber-dojo.org. The problem has the following description:
//
// Write a program to score a game of Ten-Pin Bowling.
//
// Input: string (described below) representing a bowling game
// Ouput: integer score
//
// The scoring rules:
//
// Each game, or "line" of bowling, includes ten turns,
// or "frames" for the bowler.
//
// In each frame, the bowler gets up to two tries to
// knock down all ten pins.
//
// If the first ball in a frame knocks down all ten pins,
// this is called a "strike". The frame is over. The score
// for the frame is ten plus the total of the pins knocked
// down in the next two balls.
//
// If the second ball in a frame knocks down all ten pins,
// this is called a "spare". The frame is over. The score
// for the frame is ten plus the number of pins knocked
// down in the next ball.
//
// If, after both balls, there is still at least one of the
// ten pins standing the score for that frame is simply
// the total number of pins knocked down in those two balls.
//
// If you get a spare in the last (10th) frame you get one
// more bonus ball. If you get a strike in the last (10th)
// frame you get two more bonus balls.
// These bonus throws are taken as part of the same turn.
// If a bonus ball knocks down all the pins, the process
// does not repeat. The bonus balls are only used to
// calculate the score of the final frame.
//
// The game score is the total of all frame scores.
//
// Examples:
//
// X indicates a strike
// / indicates a spare
// - indicates a miss
// | indicates a frame boundary
// The characters after the || indicate bonus balls
//
// X|X|X|X|X|X|X|X|X|X||XX
// Ten strikes on the first ball of all ten frames.
// Two bonus balls, both strikes.
// Score for each frame == 10 + score for next two
// balls == 10 + 10 + 10 == 30
// Total score == 10 frames x 30 == 300
//
// 9-|9-|9-|9-|9-|9-|9-|9-|9-|9-||
// Nine pins hit on the first ball of all ten frames.
// Second ball of each frame misses last remaining pin.
// No bonus balls.
// Score for each frame == 9
// Total score == 10 frames x 9 == 90
//
// 5/|5/|5/|5/|5/|5/|5/|5/|5/|5/||5
// Five pins on the first ball of all ten frames.
// Second ball of each frame hits all five remaining
// pins, a spare.
// One bonus ball, hits five pins.
// Score for each frame == 10 + score for next one
// ball == 10 + 5 == 15
// Total score == 10 frames x 15 == 150
//
// X|7/|9-|X|-8|8/|-6|X|X|X||81
// Total score == 167
type Game struct {
	frames     []frame
	bonusBall1 int64
	bonusBall2 int64
}

// NewGame creates a game that represents a game of Ten-Pin Bowling
func NewGame(input string) (Game, error) {
	parts := strings.Split(input, "||")
	if len(parts) != 2 {
		return Game{}, fmt.Errorf("bad input string: %v", input)
	}

	framesInput := parts[0]
	bonusBallsInput := parts[1]

	frameParts := strings.Split(framesInput, "|")
	if len(frameParts) != 10 {
		return Game{}, fmt.Errorf("bad input string: %v", input)
	}

	var frames []frame
	for i := 0; i < 10; i++ {
		f, err := newFrame(frameParts[i])
		if err != nil {
			return Game{}, fmt.Errorf("error creating frame %v: %v", i, err)
		}
		frames = append(frames, f)
	}

	var bonusBall1 int64
	var bonusBall2 int64
	if len(bonusBallsInput) > 0 {
		ball, err := parseBonusBall(bonusBallsInput[:1])
		if err != nil {
			return Game{}, fmt.Errorf("error parsing bonus ball: %v", err)
		}
		bonusBall1 = ball
	}
	if len(bonusBallsInput) > 1 {
		ball, err := parseBonusBall(bonusBallsInput[1:2])
		if err != nil {
			return Game{}, fmt.Errorf("error parsing bonus ball: %v", err)
		}
		bonusBall2 = ball
	}

	return Game{
		frames:     frames,
		bonusBall1: bonusBall1,
		bonusBall2: bonusBall2,
	}, nil
}

// Score scores a game of Ten-Pin bowling
func (g Game) Score() int {
	var score int64

	for i := 0; i < 10; i++ {
		score += g.frames[i].first + g.frames[i].second

		if g.frames[i].strike {
			if i < 9 {
				if g.frames[i+1].strike && i < 8 {
					score += g.frames[i+1].first + g.frames[i+2].first
				} else if g.frames[i+1].strike && i < 9 {
					score += g.frames[i+1].first + g.bonusBall1
				} else {
					score += g.frames[i+1].first + g.frames[i+1].second
				}
			} else if i == 9 {
				score += g.bonusBall1 + g.bonusBall2
			}
		} else if g.frames[i].spare {
			if i < 9 {
				score += g.frames[i+1].first
			} else if i == 9 {
				score += g.bonusBall1
			}
		}
	}

	return int(score)
}

func parseBonusBall(input string) (int64, error) {
	if input == "X" {
		return 10, nil
	}

	if input == "-" {
		return 0, nil
	}

	bonusBall, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("bad input string: %v", input)
	}

	return bonusBall, nil
}

type frame struct {
	strike bool
	spare  bool
	first  int64
	second int64
}

func newFrame(src string) (frame, error) {
	if len(src) < 1 || len(src) > 2 {
		return frame{}, fmt.Errorf("bad src string: %v", src)
	}

	if src == "X" {
		return frame{
			strike: true,
			spare:  false,
			first:  10,
			second: 0,
		}, nil
	}

	if src[:1] == "-" {
		if src[1:2] == "-" {
			return frame{
				strike: false,
				spare:  false,
				first:  0,
				second: 0,
			}, nil
		}

		j, err := strconv.ParseInt(src[1:2], 10, 64)
		if err != nil {
			return frame{}, fmt.Errorf("bad src string: %v", src)
		}

		return frame{
			strike: false,
			spare:  false,
			first:  0,
			second: j,
		}, nil
	}

	i, err := strconv.ParseInt(src[:1], 10, 64)
	if err != nil {
		return frame{}, fmt.Errorf("bad src string: %v", src)
	}

	if src[1:2] == "-" {
		return frame{
			strike: false,
			spare:  false,
			first:  i,
			second: 0,
		}, nil
	}

	if src[1:2] == "/" {
		return frame{
			strike: false,
			spare:  true,
			first:  i,
			second: 10 - i,
		}, nil
	}

	j, err := strconv.ParseInt(src[1:2], 10, 64)
	if err != nil {
		return frame{}, fmt.Errorf("bad src string: %v", src)
	}

	return frame{
		strike: false,
		spare:  false,
		first:  i,
		second: j,
	}, nil
}
