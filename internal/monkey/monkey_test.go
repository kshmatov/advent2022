package monkey

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data = strings.Split(`Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`, "\n")
	shortData = strings.Split(`Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
	If true: throw to monkey 2
	If false: throw to monkey 3`, "\n")
)

func TestParseMonkey(t *testing.T) {
	i, m, err := parseMonkey(shortData)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, 0, i)
	assert.Equal(t, worryLevel(1501), m.op())
	assert.Equal(t, uint(23), m.test)
	assert.Equal(t, 2, m.ifTrue)
	assert.Equal(t, 3, m.ifFalse)
}

func TestBuildShort(t *testing.T) {
	m, err := BuildMokeyList(shortData, 2, 3)
	if !assert.Nil(t, err) {
		return
	}
	if !assert.NotNil(t, m.data[0]) {
		return
	}
	assert.Equal(t, worryLevel(1501), m.data[0].op())
	assert.Equal(t, uint(23), m.data[0].test)
	assert.Equal(t, 2, m.data[0].ifTrue)
	assert.Equal(t, 3, m.data[0].ifFalse)
}

func TestBuildFull(t *testing.T) {
	res := []struct {
		items   int
		op      worryLevel
		test    uint
		ifTrue  int
		ifFalse int
	}{
		{
			items:   2,
			op:      1501,
			test:    23,
			ifTrue:  2,
			ifFalse: 3,
		},
		{
			items:   4,
			op:      60,
			test:    19,
			ifTrue:  2,
			ifFalse: 0,
		},
		{
			items:   3,
			op:      6241,
			test:    13,
			ifTrue:  1,
			ifFalse: 3,
		},
		{
			items:   1,
			op:      77,
			test:    17,
			ifTrue:  0,
			ifFalse: 1,
		},
	}
	mks, err := BuildMokeyList(data, 4, 3)
	if !assert.Nil(t, err) {
		return
	}
	for i, m := range mks.data {
		if !assert.NotNil(t, m) {
			return
		}
		assert.Equalf(t, res[i].items, len(m.items), "monkei %v", i)
		assert.Equalf(t, res[i].op, m.op(), "monkei %v", i)
		assert.Equalf(t, res[i].test, m.test, "monkei %v", i)
		assert.Equalf(t, res[i].ifTrue, m.ifTrue, "monkei %v", i)
		assert.Equalf(t, res[i].ifFalse, m.ifFalse, "monkei %v", i)
	}
}

func TestMonkeyGet(t *testing.T) {
	_, m, err := parseMonkey(shortData)
	if !assert.Nil(t, err) {
		return
	}
	l, to, err := m.getItem(3)
	assert.Nil(t, err)
	assert.Equal(t, worryLevel(500), l)
	assert.Equal(t, 3, to)
	assert.Equal(t, 1, m.cnt)

	l, to, err = m.getItem(3)
	assert.Nil(t, err)
	assert.Equal(t, worryLevel(620), l)
	assert.Equal(t, 3, to)
	assert.Equal(t, 2, m.cnt)

	_, _, err = m.getItem(3)
	assert.ErrorIs(t, EOI, err)
	assert.Equal(t, 2, m.cnt)
	assert.Nil(t, m.items)
}

func TestRound(t *testing.T) {
	m, err := BuildMokeyList(data, 4, 3)
	if !assert.Nil(t, err) {
		return
	}
	m.round()
	assert.Equal(t, []worryLevel{20, 23, 27, 26}, m.data[0].items)
	assert.Equal(t, []worryLevel{2080, 25, 167, 207, 401, 1046}, m.data[1].items)
	assert.Nil(t, m.data[2].items)
	assert.Nil(t, m.data[3].items)
}

func TestRun(t *testing.T) {
	m, err := BuildMokeyList(data, 4, 3)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(20)
	assert.Equal(t, []worryLevel{10, 12, 14, 26, 34}, m.data[0].items)
	assert.Equal(t, []worryLevel{245, 93, 53, 199, 115}, m.data[1].items)
	assert.Nil(t, m.data[2].items)
	assert.Nil(t, m.data[3].items)

	assert.Equal(t, 101, m.data[0].cnt)
	assert.Equal(t, 95, m.data[1].cnt)
	assert.Equal(t, 7, m.data[2].cnt)
	assert.Equal(t, 105, m.data[3].cnt)

	res := m.GetTopN(2)
	assert.Equal(t, []int{105, 101}, res)
}

func TestBIGRun(t *testing.T) {
	m, err := BuildMokeyList(data, 4, 0)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(1)
	assert.Equal(t, 2, m.data[0].cnt)
	assert.Equal(t, 4, m.data[1].cnt)
	assert.Equal(t, 3, m.data[2].cnt)
	assert.Equal(t, 6, m.data[3].cnt)

	m, err = BuildMokeyList(data, 4, 0)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(20)
	assert.Equal(t, 99, m.data[0].cnt)
	assert.Equal(t, 97, m.data[1].cnt)
	assert.Equal(t, 8, m.data[2].cnt)
	assert.Equal(t, 103, m.data[3].cnt)

	m, err = BuildMokeyList(data, 4, 0)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(1000)
	assert.Equal(t, 5204, m.data[0].cnt)
	assert.Equal(t, 4792, m.data[1].cnt)
	assert.Equal(t, 199, m.data[2].cnt)
	assert.Equal(t, 5192, m.data[3].cnt)

	m, err = BuildMokeyList(data, 4, 0)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(2000)
	assert.Equal(t, 10419, m.data[0].cnt)
	assert.Equal(t, 9577, m.data[1].cnt)
	assert.Equal(t, 392, m.data[2].cnt)
	assert.Equal(t, 10391, m.data[3].cnt)

	m, err = BuildMokeyList(data, 4, 0)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(3000)
	assert.Equal(t, 15638, m.data[0].cnt)
	assert.Equal(t, 14358, m.data[1].cnt)
	assert.Equal(t, 587, m.data[2].cnt)
	assert.Equal(t, 15593, m.data[3].cnt)

	m, err = BuildMokeyList(data, 4, 0)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(5000)
	assert.Equal(t, 26075, m.data[0].cnt)
	assert.Equal(t, 23921, m.data[1].cnt)
	assert.Equal(t, 974, m.data[2].cnt)
	assert.Equal(t, 26000, m.data[3].cnt)

	m, err = BuildMokeyList(data, 4, 0)
	if !assert.Nil(t, err) {
		return
	}
	m.Run(10000)
	assert.Equal(t, 52166, m.data[0].cnt)
	assert.Equal(t, 47830, m.data[1].cnt)
	assert.Equal(t, 1938, m.data[2].cnt)
	assert.Equal(t, 52013, m.data[3].cnt)
	res := m.GetTopN(2)
	assert.Equal(t, []int{52166, 52013}, res)
}
