package signal

import (
	"github.com/pkg/errors"
)

func Start(s string, l int) (int, error) {
	for i := range s[:len(s)-l] {
		if checkSubstr(s[i : i+l]) {
			return i, nil
		}
	}
	return 0, errors.New("not found")
}

func checkSubstr(s string) bool {
	control := map[rune]struct{}{}
	for _, a := range s {
		control[a] = struct{}{}
	}
	return len(control) == len(s)
}
