package climber

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data = strings.Split(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`, "\n")
)

func TestBuildMap(t *testing.T) {
	m := buildMap(data)
	assert.Equal(t, point{0, 0}, m.end)
	assert.Equal(t, point{y: 2, x: 5}, m.start)
	assert.True(t, m.isEnd(point{0, 0}))
	assert.True(t, m.checkMove('E', point{y: 2, x: 4}))
	assert.False(t, m.checkMove('E', point{y: 2, x: 6}))
	points := m.getWays(point{x: 5, y: 2})
	if !assert.Equal(t, 1, len(points)) {
		return
	}
	assert.Equal(t, point{y: 2, x: 4}, points[0])
}

func TestTrack(t *testing.T) {
	m := buildMap(data)
	tr := &track{
		position: m.start,
		visited:  map[string]struct{}{},
	}
	res := tr.move(m)
	if !assert.Equal(t, 1, len(res)) {
		return
	}
	r := res[0]
	assert.Equal(t, point{y: 2, x: 4}, r.position)
	assert.Equal(t, 1, r.len)
}

func TestTravel(t *testing.T) {
	m := buildMap(data)
	i := travel(m, m.isEnd)
	assert.Equal(t, 31, i)
}

func TestLowest(t *testing.T) {
	m := buildMap(data)
	i := travel(m, m.isLowest)
	assert.Equal(t, 29, i)
}
