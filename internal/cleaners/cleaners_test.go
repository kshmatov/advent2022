package cleaners

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = strings.Split(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`, "\n")

var trueData = []string{
	"1-9,1-1",
	"1-9,9-9",
	"1-1,1-9",
	"9-9,1-9",
	"5-5,5-5",
	"1-9,1-9",
}

var overlapres = []bool{false, false, true, true, true, true}

var res = []bool{false, false, false, true, true, false}

func TestCheck(t *testing.T) {
	for i, item := range data {
		pairs, err := makeSections(item)
		assert.Nil(t, err)
		r := included(pairs[0], pairs[1])
		assert.Equalf(t, res[i], r, "%v expected %v got %v (%v)", item, res[i], r, pairs)
	}
	for _, item := range trueData {
		pairs, err := makeSections(item)
		assert.Nil(t, err)
		r := included(pairs[0], pairs[1])
		assert.Truef(t, r, "%v expected true got %v (%v)", item, r, pairs)
	}
	for i, item := range data {
		pairs, err := makeSections(item)
		assert.Nil(t, err)
		r := overlaped(pairs[0], pairs[1])
		assert.Equalf(t, overlapres[i], r, "%v expected %v got %v (%v)", item, overlapres[i], r, pairs)
	}
}
