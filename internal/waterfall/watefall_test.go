package waterfall

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeBlocks(t *testing.T) {
	// 498,4 -> 498,6 -> 496,6
	p := []Point{
		{x: 498, y: 4},
		{x: 498, y: 6},
		{x: 496, y: 6},
	}
	res := map[Point]fillament{
		{x: 498, y: 4}: stone,
		{x: 498, y: 5}: stone,
		{x: 498, y: 6}: stone,
		{x: 497, y: 6}: stone,
		{x: 496, y: 6}: stone,
	}
	b := makeBlocks(p...)
	assert.Equal(t, 5, len(b))
	assert.Equal(t, res, b)

	// 503,4 -> 502,4 -> 502,9 -> 494,9
	p = []Point{
		{x: 503, y: 4},
		{x: 502, y: 4},
		{x: 502, y: 9},
		{x: 494, y: 9},
	}
	res = map[Point]fillament{
		{x: 503, y: 4}: stone,
		{x: 502, y: 4}: stone,
		{x: 502, y: 5}: stone,
		{x: 502, y: 6}: stone,
		{x: 502, y: 7}: stone,
		{x: 502, y: 8}: stone,
		{x: 502, y: 9}: stone,
		{x: 501, y: 9}: stone,
		{x: 500, y: 9}: stone,
		{x: 499, y: 9}: stone,
		{x: 498, y: 9}: stone,
		{x: 497, y: 9}: stone,
		{x: 496, y: 9}: stone,
		{x: 495, y: 9}: stone,
		{x: 494, y: 9}: stone,
	}
	b = makeBlocks(p...)
	assert.Equal(t, len(res), len(b))
	assert.Equal(t, res, b)
}

func TestFillMap(t *testing.T) {
	data := strings.Split(`498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9`, "\n")
	cnt := Blocks{b: map[Point]fillament{
		{x: 498, y: 4}: stone,
		{x: 498, y: 5}: stone,
		{x: 498, y: 6}: stone,
		{x: 497, y: 6}: stone,
		{x: 496, y: 6}: stone,
		{x: 503, y: 4}: stone,
		{x: 502, y: 4}: stone,
		{x: 502, y: 5}: stone,
		{x: 502, y: 6}: stone,
		{x: 502, y: 7}: stone,
		{x: 502, y: 8}: stone,
		{x: 502, y: 9}: stone,
		{x: 501, y: 9}: stone,
		{x: 500, y: 9}: stone,
		{x: 499, y: 9}: stone,
		{x: 498, y: 9}: stone,
		{x: 497, y: 9}: stone,
		{x: 496, y: 9}: stone,
		{x: 495, y: 9}: stone,
		{x: 494, y: 9}: stone,
	}}

	b := FillMap(data)
	assert.Equal(t, cnt.b, b.b)
}

func TestFall(t *testing.T) {
	data := strings.Split(`498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9`, "\n")
	b := FillMap(data)
	_, cnt := Sand(b)
	assert.Equal(t, 24, cnt)
}

func TestFallFloor(t *testing.T) {
	data := strings.Split(`498,4 -> 498,6 -> 496,6
	503,4 -> 502,4 -> 502,9 -> 494,9`, "\n")
	b := FillMap(data)
	_, cnt := SandFloor(b)
	assert.Equal(t, 93, cnt)
}
