package rucksack

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	data = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	vals = "pLPvts"
)

var expectConverted = []int{16, 38, 42, 22, 20, 19}
var expectHalves = [][]string{
	{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
	{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
	{"PmmdzqPrV", "vPwwTWBwg"},
	{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
	{"ttgJtRGJ", "QctTZtZT"},
	{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
}

func TestConverter(t *testing.T) {
	// 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s)
	vals := "pLPvts"
	expect := []int{16, 38, 42, 22, 20, 19}
	for i, r := range vals {
		assert.Equalf(t, expect[i], converter(r), "with %v", string(r))
	}
}

func TestSplit(t *testing.T) {
	vals := strings.Split(data, "\n")
	for i, v := range vals {
		a, b := split(v)
		assert.Equal(t, expectHalves[i][0], a, "%v but %v", expectHalves[i][0], a)
		assert.Equal(t, expectHalves[i][1], b, "%v but %v", expectHalves[i][1], b)
	}
}

func TestCheckItems(t *testing.T) {
	vals := strings.Split(data, "\n")
	for i, v := range vals {
		r, err := checkItems(v)
		assert.Nil(t, err)
		assert.Equal(t, expectConverted[i], r, "with %v", v)
	}
}

func TestBackpack(t *testing.T) {
	_, err := newBackpack([]string{}, 1)
	assert.Error(t, err)
	_, err = newBackpack([]string{"a", "b"}, 3)
	assert.Error(t, err)
	b, err := newBackpack([]string{"a", "b", "c"}, 3)
	assert.Nil(t, err)
	assert.Equal(t, -3, b.pos)
	assert.Equal(t, 3, b.step)
	assert.Equal(t, 3, len(b.data))
	assert.True(t, b.next())
	assert.Equal(t, 0, b.pos)
	r := b.get()
	assert.Equal(t, "a", r[0])
	assert.Equal(t, "b", r[1])
	assert.Equal(t, "c", r[2])

	assert.False(t, b.next())
	assert.Equal(t, 3, b.pos)
	r = b.get()
	assert.Nil(t, r)
	b, err = newBackpack([]string{"a", "b", "c", "d"}, 2)
	assert.Nil(t, err)
	b.next()
	b.next()
	r = b.get()
	assert.Equal(t, "c", r[0])
	assert.Equal(t, "d", r[1])

}
