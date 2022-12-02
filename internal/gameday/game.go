package gameday

import (
	"fmt"
	"strings"
)

type input string
type output string

const (
	inRock    = input("A")
	inPapper  = input("B")
	inScissor = input("C")

	outRock    = output("X")
	outPapper  = output("Y")
	outScissor = output("Z")

	shouldWin   = "Z"
	shouldDraw  = "Y"
	shouldLoose = "X"

	win   = 6
	draw  = 3
	loose = 0
)

type rules struct {
	wins   input
	looses input
	eq     input
	score  int
}

type reverseRules struct {
	wins  output
	loose output
	eq    output
}

var combinatoions = map[output]rules{
	outRock:    {wins: inScissor, looses: inPapper, eq: inRock, score: 1},
	outPapper:  {wins: inRock, looses: inScissor, eq: inPapper, score: 2},
	outScissor: {wins: inPapper, looses: inRock, eq: inScissor, score: 3},
}

var reverseCombination = map[input]reverseRules{
	inRock:    {wins: outPapper, loose: outScissor, eq: outRock},
	inPapper:  {wins: outScissor, loose: outRock, eq: outPapper},
	inScissor: {wins: outRock, loose: outPapper, eq: outScissor},
}

func Calculate(s []string) (int, error) {
	score := 0
	for num, line := range s {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		round := strings.Split(line, " ")
		if len(round) != 2 {
			return 0, fmt.Errorf("%v: bad line <%v>", num, line)
		}

		lScore, err := calcRound(round[0], round[1])
		if err != nil {
			return 0, fmt.Errorf("%v: %v", num, err)
		}
		score += lScore
	}
	return score, nil
}

func calcRound(first, second string) (int, error) {
	comb, ok := combinatoions[output(second)]
	if !ok {
		return 0, fmt.Errorf("unknown second value: <%v>", second)
	}
	switch input(first) {
	case comb.wins:
		return win + comb.score, nil
	case comb.looses:
		return comb.score, nil
	case comb.eq:
		return draw + comb.score, nil
	default:
		return 0, fmt.Errorf("unknown first value: <%v>", first)
	}
}

func ReverseCalculate(s []string) (int, error) {
	score := 0
	for num, line := range s {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		round := strings.Split(line, " ")
		if len(round) != 2 {
			return 0, fmt.Errorf("%v: bad line <%v>", num, line)
		}

		lScore, err := reverseCalcRound(round[0], round[1])
		if err != nil {
			return 0, fmt.Errorf("%v: %v", num, err)
		}
		score += lScore
	}
	return score, nil
}

func reverseCalcRound(first, result string) (int, error) {
	comb, ok := reverseCombination[input(first)]
	if !ok {
		return 0, fmt.Errorf("unknown first value: <%v>", first)
	}
	var second output
	score := 0
	switch result {
	case shouldWin:
		second = comb.wins
		score = win
	case shouldLoose:
		second = comb.loose
		score = loose
	case shouldDraw:
		second = comb.eq
		score = draw
	default:
		return 0, fmt.Errorf("unknown result value: <%v>", result)
	}
	return combinatoions[second].score + score, nil
}
