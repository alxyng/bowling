package bowling

import (
	"regexp"
	"strconv"
)

type frame struct {
	strike bool
	spare  bool
	score  int64
}

func newFrame(src string) frame {
	// fmt.Println(src)

	if src == "X" {
		return frame{
			strike: true,
			spare:  false,
			score:  10,
		}
	}

	if src[:1] == "-" {
		if src[1:2] == "-" {
			return frame{
				strike: false,
				spare:  false,
				score:  0,
			}
		}

		i, err := strconv.ParseInt(src[1:2], 10, 64)
		if err != nil {
			panic(err)
		}

		return frame{
			strike: false,
			spare:  false,
			score:  i,
		}
	}

	i, err := strconv.ParseInt(src[:1], 10, 64)
	if err != nil {
		panic(err)
	}

	if src[1:2] == "-" {
		return frame{
			strike: false,
			spare:  false,
			score:  i,
		}
	}

	if src[1:2] == "/" {
		return frame{
			strike: false,
			spare:  true,
			score:  10,
		}
	}

	j, err := strconv.ParseInt(src[1:2], 10, 64)
	if err != nil {
		panic(err)
	}

	return frame{
		strike: false,
		spare:  false,
		score:  i + j,
	}
}

func score(game string) int {
	r := regexp.MustCompile("^([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|([X/0-9\\-]+)\\|\\|([X/0-9\\-]*)$")
	matches := r.FindAllStringSubmatch(game, -1)
	if matches == nil {
		return -1
	}

	frames := []frame{}

	results := matches[0]
	for i := 1; i < 11; i++ {
		frames = append(frames, newFrame(results[i]))
	}

	var score int64
	for _, frame := range frames {
		score += frame.score
	}

	return int(score)
}
