package beacon

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testData = strings.Split(`Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	Sensor at x=9, y=16: closest beacon is at x=10, y=16
	Sensor at x=13, y=2: closest beacon is at x=15, y=3
	Sensor at x=12, y=14: closest beacon is at x=10, y=16
	Sensor at x=10, y=20: closest beacon is at x=10, y=16
	Sensor at x=14, y=17: closest beacon is at x=10, y=16
	Sensor at x=8, y=7: closest beacon is at x=2, y=10
	Sensor at x=2, y=0: closest beacon is at x=2, y=10
	Sensor at x=0, y=11: closest beacon is at x=2, y=10
	Sensor at x=20, y=14: closest beacon is at x=25, y=17
	Sensor at x=17, y=20: closest beacon is at x=21, y=22
	Sensor at x=16, y=7: closest beacon is at x=15, y=3
	Sensor at x=14, y=3: closest beacon is at x=15, y=3
	Sensor at x=20, y=1: closest beacon is at x=15, y=3`, "\n")
)

func TestRange(t *testing.T) {
	b := Ping{
		sensor: Point{X: 8, Y: 7},
		signal: Point{X: 2, Y: 10},
	}
	assert.Equal(t, 9, b.freeDistance())

	b = Ping{
		sensor: Point{X: 2, Y: 10},
		signal: Point{X: 8, Y: 7},
	}
	assert.Equal(t, 9, b.freeDistance())
}

func TestExtractPoint(t *testing.T) {
	p := extractPoint("Sensor at x=2, y=18")
	assert.Equal(t, Point{X: 2, Y: 18}, p)
	p = extractPoint("closest beacon is at x=-2, y=15")
	assert.Equal(t, Point{X: -2, Y: 15}, p)
}

func TestBuildList(t *testing.T) {
	b := BuildList([]string{"Sensor at x=2, y=18: closest beacon is at x=-2, y=15"})
	if !assert.Equal(t, 1, len(b)) {
		t.Fatal("bad list")
	}
	assert.Equal(t, Ping{sensor: Point{X: 2, Y: 18}, signal: Point{X: -2, Y: 15}}, b[0])
}

func TestYRange(t *testing.T) {
	b := Ping{sensor: Point{X: 2, Y: 18}, signal: Point{X: -2, Y: 15}}
	y := 10
	assert.Equal(t, 8, b.yRange(y))
	b.sensor.Y = -2
	assert.Equal(t, 12, b.yRange(y))
	y = -1
	assert.Equal(t, 1, b.yRange(y))
	b.sensor.Y = 5
	assert.Equal(t, 6, b.yRange(y))
}

func TestFreeSpace(t *testing.T) {
	b := Ping{sensor: Point{X: 8, Y: 7}, signal: Point{X: 2, Y: 10}}

	m := NewRow()
	m.CalcFreeSpaces(b, -3)
	assert.Equalf(t, 0, m.Len(), "%+v", m)

	m = NewRow()
	m.CalcFreeSpaces(b, -2)
	assert.Equalf(t, 1, m.Len(), "%+v", m)

	m = NewRow()
	m.CalcFreeSpaces(b, -1)
	assert.Equalf(t, 3, m.Len(), "%+v", m)

	m = NewRow()
	m.CalcFreeSpaces(b, 7)
	assert.Equalf(t, 19, m.Len(), "%+v", m)

	m = NewRow()
	m.CalcFreeSpaces(b, 10)
	assert.Equal(t, 12, m.Len(), "%+v", m)
}

func TestTask1(t *testing.T) {
	b := BuildList(testData)
	r := Row{}
	for _, i := range b {
		r.CalcFreeSpaces(i, 10)
	}
	assert.Equal(t, 26, r.Len())
}

func TestBorders(t *testing.T) {
	p := Ping{sensor: Point{X: 8, Y: 7}, signal: Point{X: 2, Y: 10}}
	b := p.Borders()
	assert.Equalf(t, 40, len(b), "%+v", b)
	_, ok := b[Point{X: 14, Y: 11}]
	assert.Truef(t, ok, "%+v", b)
}

func TestTask2(t *testing.T) {
	p := BuildList(testData)
	b := BuildBorders(true, p)
	res := FilterBorders(0, 20, b)
	assert.Equal(t, Point{X: 14, Y: 11}, res)
}
