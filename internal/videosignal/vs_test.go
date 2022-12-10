package videosignal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ctlData = strings.Split(`addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`, "\n")

func TestController(t *testing.T) {
	c := newController(2, 5)

	v := c.check(1)
	assert.Equal(t, 0, v)
	assert.Equal(t, 2, c.cur)

	v = c.check(2)
	assert.Equal(t, 2, v)
	assert.Equal(t, 7, c.cur)

	v = c.check(1)
	assert.Equal(t, 0, v)
	assert.Equal(t, 7, c.cur)

	v = c.check(9)
	assert.Equal(t, 7, v)
	assert.Equal(t, 12, c.cur)
}

func TestAdd(t *testing.T) {
	c := NewCPU(1, 2)
	c.noop()
	assert.Equal(t, 1, c.register)
	assert.Equal(t, 1, c.cycle)
	c.addx(3)
	assert.Equal(t, 4, c.register)
	assert.Equal(t, 3, c.cycle)
	c.addx(-5)
	assert.Equal(t, -1, c.register)
	assert.Equal(t, 5, c.cycle)

	assert.Equal(t, 1, c.mem[1])
	assert.Equal(t, 1, c.mem[3])
	assert.Equal(t, 4, c.mem[5])
}

func TestExec(t *testing.T) {
	expected := map[int]int{
		20:  21,
		60:  19,
		100: 18,
		140: 21,
		180: 16,
		220: 18,
	}
	c := NewCPU(20, 40)
	err := Exec(ctlData, c)
	assert.Nil(t, err)
	m := c.Mem()
	for k, v := range expected {
		assert.Equalf(t, v, m[k], "%v expected %v, got %v", k, v, m[k])
	}
	assert.Equal(t, 13140, c.ControlSum())
}

func TestShortExec(t *testing.T) {
	s := strings.Split(`addx 15
	addx -11
	addx 6
	addx -3
	addx 5
	addx -1
	addx -8
	addx 13
	addx 4
	noop
	addx -1`, "\n")
	c := NewCPU(20, 40).SetDebug(true)
	err := Exec(s, c)
	assert.Nil(t, err)
	assert.Equal(t, 21, c.mem[20])
}
