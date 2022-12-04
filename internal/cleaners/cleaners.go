package cleaners

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type sections struct {
	start int
	end   int
}

func included(a, b sections) bool {
	if a.start <= b.start && a.end >= b.end {
		return true
	}
	if a.start >= b.start && a.end <= b.end {
		return true
	}
	return false
}

func overlaped(a, b sections) bool {
	if a.start <= b.end && a.end >= b.start {
		return true
	}
	return false
}

func makeSections(s string) ([]sections, error) {
	l := strings.Split(s, ",")
	if len(l) != 2 {
		return nil, errors.New("must be 2 elves")
	}
	elves := []sections{}
	for i, s := range l {
		margins := strings.Split(s, "-")
		if len(margins) != 2 {
			return nil, fmt.Errorf("elf %v has bad margins <%v>", i, s)
		}
		b, err := strconv.ParseInt(margins[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("elf %v cant parse low margin %v", i, margins[0])
		}
		e, err := strconv.ParseInt(margins[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("elf %v cant parse high margin %v", i, margins[1])
		}
		elves = append(elves, sections{int(b), int(e)})
	}
	return elves, nil
}

func buildAndCheck(s string, f func(sections, sections) bool) (bool, error) {
	pairs, err := makeSections(s)
	if err != nil {
		return false, err
	}
	if len(pairs) != 2 {
		return false, fmt.Errorf("must be 2 sections, has %v", len(pairs))
	}
	if f(pairs[0], pairs[1]) {
		return true, nil
	}
	return false, nil
}

func CheckInclusion(s []string) (int, error) {
	cnt := 0
	for i, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		check, err := buildAndCheck(line, included)
		if err != nil {
			return 0, errors.Wrapf(err, "line %v", i)
		}
		if check {
			cnt++
		}
	}
	return cnt, nil
}

func CheckOverlaption(s []string) (int, error) {
	cnt := 0
	for i, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		check, err := buildAndCheck(line, overlaped)
		if err != nil {
			return 0, errors.Wrapf(err, "line %v", i)
		}
		if check {
			cnt++
		}
	}
	return cnt, nil
}
