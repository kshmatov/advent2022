package rope

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data = strings.Split(`
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`, "\n")
	longData = strings.Split(`
	R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`, "\n")
)

func TestMove(t *testing.T) {
	r := NewRope()
	_ = r.move("U")
	assert.Equal(t, 0, r.tail.x)
	assert.Equal(t, 0, r.tail.y)
	_ = r.move("U")
	assert.Equal(t, 0, r.tail.x)
	assert.Equal(t, 1, r.tail.y)
	_ = r.move("D")
	assert.Equal(t, 0, r.tail.x)
	assert.Equal(t, 1, r.tail.y)
	_ = r.move("D")
	assert.Equal(t, 0, r.tail.x)
	assert.Equal(t, 1, r.tail.y)
	_ = r.move("D")
	assert.Equal(t, 0, r.tail.x)
	assert.Equal(t, 0, r.tail.y)
	assert.Equal(t, 2, r.path.length())
	_ = r.move("U")

	_ = r.move("R")
	assert.Equal(t, 0, r.tail.x)
	assert.Equal(t, 0, r.tail.y)
	_ = r.move("R")
	assert.Equal(t, 1, r.tail.x)
	assert.Equal(t, 0, r.tail.y)
	_ = r.move("L")
	assert.Equal(t, 1, r.tail.x)
	assert.Equal(t, 0, r.tail.y)
	_ = r.move("L")
	assert.Equal(t, 1, r.tail.x)
	assert.Equal(t, 0, r.tail.y)
	_ = r.move("L")
	assert.Equal(t, 0, r.tail.x)
	assert.Equal(t, 0, r.tail.y)
	assert.Equal(t, 3, r.path.length())
	_ = r.move("R")
	assert.Equal(t, 0, r.head.x)
	assert.Equal(t, 0, r.head.y)
	_ = r.move("U")
	_ = r.move("R")
	_ = r.move("U")
	assert.Equal(t, 1, r.tail.x)
	assert.Equal(t, 1, r.tail.y)
	assert.Equal(t, 4, r.path.length())
}

func TestTravel(t *testing.T) {
	r := NewRope()
	coms, err := BuildPath(data)
	assert.Nil(t, err)
	assert.Equal(t, 8, len(coms))
	err = r.Travel(coms)
	assert.Nil(t, err)
	assert.Equal(t, 13, r.path.length())
}

func TestLongRope(t *testing.T) {
	l := NewLongRope(10)
	p, err := BuildPath(longData)
	assert.Nil(t, err)
	err = l.Travel(p)
	assert.Nil(t, err)
	assert.Equal(t, 36, l.Points())
}

func TestLongRopeDiag2(t *testing.T) {
	l := NewLongRope(2)
	l.knots[0].x = 1
	l.knots[0].y = 2
	_ = l.move("R")
	assert.Equalf(t, 1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	l = NewLongRope(2)
	l.knots[0].x = -1
	l.knots[0].y = 2
	_ = l.move("L")
	assert.Equalf(t, -1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)

	l = NewLongRope(2)
	l.knots[0].x = 1
	l.knots[0].y = -2
	_ = l.move("R")
	assert.Equalf(t, 1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, -1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)

	l = NewLongRope(2)
	l.knots[0].x = -1
	l.knots[0].y = -2
	_ = l.move("L")
	assert.Equalf(t, -1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, -1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
}

func TestLongRopeHorse2(t *testing.T) {
	l := NewLongRope(2)
	l.knots[0].x = 1
	l.knots[0].y = 1
	_ = l.move("R")
	assert.Equalf(t, 1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	l = NewLongRope(2)
	l.knots[0].x = -1
	l.knots[0].y = 1
	_ = l.move("L")
	assert.Equalf(t, -1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)

	l = NewLongRope(2)
	l.knots[0].x = 1
	l.knots[0].y = -1
	_ = l.move("R")
	assert.Equalf(t, 1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, -1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)

	l = NewLongRope(2)
	l.knots[0].x = -1
	l.knots[0].y = -1
	_ = l.move("L")
	assert.Equalf(t, -1, l.knots[1].x, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, -1, l.knots[1].y, "Expected %v, %v, got %v, %v", 1, 1, l.knots[1].x, l.knots[1].y)
}

func TestLongRopeLineMove(t *testing.T) {
	l := NewLongRope(2)
	_ = l.move("R")
	assert.Equalf(t, 0, l.knots[1].x, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 0, l.knots[1].y, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
	l = NewLongRope(2)
	_ = l.move("L")
	assert.Equalf(t, 0, l.knots[1].x, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 0, l.knots[1].y, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
	l = NewLongRope(2)
	_ = l.move("U")
	assert.Equalf(t, 0, l.knots[1].x, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 0, l.knots[1].y, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
	l = NewLongRope(2)
	_ = l.move("D")
	assert.Equalf(t, 0, l.knots[1].x, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
	assert.Equalf(t, 0, l.knots[1].y, "Expected %v, %v, got %v, %v", 0, 0, l.knots[1].x, l.knots[1].y)
}

func TestLongRopeMove(t *testing.T) {
	l := NewLongRope(10)
	_ = l.move("R")
	_ = l.move("R")
	_ = l.move("R")
	_ = l.move("R")
	assert.Equal(t, 3, l.knots[1].x)
	assert.Equal(t, 2, l.knots[2].x)
	assert.Equal(t, 1, l.knots[3].x)
	_ = l.move("U")
	assert.Equal(t, 3, l.knots[1].x)
	assert.Equal(t, 2, l.knots[2].x)
	assert.Equal(t, 1, l.knots[3].x)
	_ = l.move("U")
	assert.Equalf(t, 4, l.knots[1].x, "0: x %v, y %v , 1: x %v, y %v", l.knots[0].x, l.knots[0].y, l.knots[1].x, l.knots[1].y)
	assert.Equal(t, 1, l.knots[1].y)
	assert.Equal(t, 3, l.knots[2].x)
	assert.Equal(t, 1, l.knots[2].y)
	assert.Equal(t, 2, l.knots[3].x)
	assert.Equal(t, 1, l.knots[3].y)

	_ = l.move("U")
	assert.Equal(t, 4, l.knots[1].x)
	assert.Equal(t, 2, l.knots[1].y)
	assert.Equal(t, 3, l.knots[2].x)
	assert.Equal(t, 1, l.knots[2].y)
	assert.Equal(t, 2, l.knots[3].x)
	assert.Equal(t, 1, l.knots[3].y)

	_ = l.move("U")
	assert.Equal(t, 4, l.knots[1].x)
	assert.Equal(t, 3, l.knots[1].y)
	assert.Equal(t, 4, l.knots[2].x)
	assert.Equal(t, 2, l.knots[2].y)
	assert.Equal(t, 3, l.knots[3].x)
	assert.Equal(t, 2, l.knots[3].y)
}
