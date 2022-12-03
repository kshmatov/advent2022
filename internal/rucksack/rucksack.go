package rucksack

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	base      = 96
	upperBase = 64 - 26
)

func converter(a rune) int {
	s := byte(a)
	if s >= base {
		return int(s - 96)
	}
	return int(s - upperBase)
}

func split(s string) (string, string) {
	mid := len(s) / 2
	first := s[:mid]
	second := string(s[mid:])
	return first, second
}

func checkItems(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("empty string")
	}
	if len(s)%2 != 0 {
		return 0, errors.New("odd len")
	}

	result := 0
	first, second := split(s)
	exists := ""
	for _, r := range first {
		if strings.ContainsRune(second, r) && !strings.ContainsRune(exists, r) {
			exists += string(r)
			result += converter(r)
		}
	}

	return result, nil
}

func CheckRuckSack(s []string) (int, error) {
	result := 0
	for num, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		r, err := checkItems(line)
		if err != nil {
			return 0, errors.Wrapf(err, "%v: %v", num, line)
		}
		result += r
	}
	return result, nil
}

// second part

type backpacks struct {
	data []string
	pos  int
	step int
}

func newBackpack(data []string, step int) (*backpacks, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data")
	}
	if len(data)%step != 0 {
		return nil, fmt.Errorf("must be grouped by step (%v by %v)", len(data), step)
	}
	return &backpacks{data: data, pos: step * -1, step: step}, nil
}

func (b *backpacks) next() bool {
	b.pos += b.step
	if b.pos > len(b.data) {
		b.pos = len(b.data)
	}

	return b.pos < len(b.data)
}

func (b *backpacks) get() []string {
	if b.pos >= len(b.data) {
		return nil
	}
	return b.data[b.pos : b.pos+b.step]
}

func getBadge(s ...string) rune {
	control := s[0]
	for _, r := range control {
		ok := true
		for _, pack := range s[1:] {
			if !strings.ContainsRune(pack, r) {
				ok = false
			}
		}
		if ok {
			return r
		}
	}
	return 0
}

func GetBadges(s []string) (int, error) {
	prepared := []string{}
	for _, line := range s {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		prepared = append(prepared, line)
	}
	bp, err := newBackpack(prepared, 3)
	if err != nil {
		return 0, err
	}
	result := 0
	for bp.next() {
		group := bp.get()
		r := getBadge(group...)
		if r == 0 {
			return 0, fmt.Errorf("group <%v> doesnt have badge", group)
		}
		result += converter(r)
	}
	return result, nil
}
