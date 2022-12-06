package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	data := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}
	for k, v := range data {
		i, err := Start(k, 4)
		i += 4
		assert.Nilf(t, err, "%v expected nil errror", k)
		assert.Equalf(t, v, i, "%v expected %v got %v", k, v, i)
	}
}

func TestMessage(t *testing.T) {
	data := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}
	for k, v := range data {
		i, err := Start(k, 14)
		i += 14
		assert.Nilf(t, err, "%v expected nil errror", k)
		assert.Equalf(t, v, i, "%v expected %v got %v", k, v, i)
	}
}

func TestSbdstr(t *testing.T) {
	data := map[string]bool{
		"asdfa":       false,
		"asdfg":       true,
		"asas":        false,
		"aaaaaaaaaaa": false,
		"alskdjfhg":   true,
	}
	for k, v := range data {
		r := checkSubstr(k)
		assert.Equal(t, v, r, k)
	}
}
