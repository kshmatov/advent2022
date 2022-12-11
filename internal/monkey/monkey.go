package monkey

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	EOI = errors.New("end of list")
)

type operation func() worryLevel

type worryLevel uint64

type monkey struct {
	items   []worryLevel
	op      operation
	test    int
	ifTrue  int
	ifFalse int
	cnt     int
}

func (m *monkey) getHead() worryLevel {
	return m.items[0]
}

type Monkeys struct {
	lower int
	data  []*monkey
	modul int
}

func getNumFromMonkey(s string) (int, error) {
	numStr := string(s[:len(s)-1])
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, errors.Wrapf(err, "parse num from  '%v' <%v>", s, numStr)
	}
	return int(num), nil
}

func getItemsFromMonkey(s string) ([]worryLevel, error) {
	var res []worryLevel
	is := strings.Split(s, ",")
	for _, item := range is {
		item = strings.TrimSpace(item)
		wl, err := strconv.Atoi(item)
		if err != nil {
			return nil, errors.Wrapf(err, "parse item from <%v> '%v' <%v>", item, s, is)
		}
		res = append(res, worryLevel(wl))
	}
	return res, nil
}

func parseMonkey(s []string) (int, *monkey, error) {
	if len(s) != 6 {
		return 0, nil, errors.New("not enaought data")
	}

	m := monkey{}

	parts := strings.Split(s[0], " ")
	num, err := getNumFromMonkey(parts[1])
	if err != nil {
		return 0, nil, err
	}

	parts = strings.Split(s[1], ":")
	m.items, err = getItemsFromMonkey(parts[1])
	if err != nil {
		return 0, nil, err
	}

	parts = strings.Split(s[2], "=")
	ops := strings.Split(strings.TrimSpace(parts[1]), " ")
	var op1, op2 operation
	if ops[0] == "old" {
		op1 = m.getHead
	} else {
		i, err := strconv.Atoi(ops[0])
		if err != nil {
			return 0, nil, err
		}
		op1 = func() worryLevel { return worryLevel(i) }
	}

	if ops[2] == "old" {
		op2 = m.getHead
	} else {
		i, err := strconv.Atoi(ops[2])
		if err != nil {
			return 0, nil, err
		}
		op2 = func() worryLevel { return worryLevel(i) }
	}

	switch ops[1] {
	case "*":
		m.op = func() worryLevel { return op1() * op2() }
	case "+":
		m.op = func() worryLevel { return op1() + op2() }
	default:
		return 0, nil, fmt.Errorf("unknown operator: %v", ops[1])
	}

	parts = strings.Split(s[3], " ")
	i, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, nil, fmt.Errorf("cant parse int from <%v>", parts)
	}
	m.test = i

	parts = strings.Split(s[4], " ")
	i, err = strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, nil, fmt.Errorf("cant parse int from <%v>", parts)
	}
	m.ifTrue = i

	parts = strings.Split(s[5], " ")
	i, err = strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return 0, nil, fmt.Errorf("cant parse int from <%v>", parts)
	}
	m.ifFalse = i

	return num, &m, nil
}

func (m *monkey) getItem(lowerWorry int) (worryLevel, int, error) {
	if len(m.items) == 0 {
		return 0, 0, EOI
	}
	defer func() {
		m.items = m.items[1:]
		if len(m.items) == 0 {
			m.items = nil
		}
	}()
	m.cnt++
	newLevel := m.op()
	if lowerWorry > 0 {
		newLevel = worryLevel(math.Trunc(float64(newLevel) / float64(lowerWorry)))
	}
	test := newLevel % worryLevel(m.test)
	if test == 0 {
		return newLevel, m.ifTrue, nil
	}
	return newLevel, m.ifFalse, nil
}

func (m *monkey) add(i worryLevel) {
	m.items = append(m.items, i)
}

func BuildMokeyList(s []string, size int, lower int) (*Monkeys, error) {
	var monkeyDesc []string

	mks := Monkeys{
		data:  make([]*monkey, size),
		lower: lower,
	}
	mod := 1
	for _, v := range s {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			if monkeyDesc != nil {
				i, m, err := parseMonkey(monkeyDesc)
				if err != nil {
					return nil, err
				}
				if len(mks.data) <= i {
					return nil, fmt.Errorf("list too short: %v, idx %v given", len(mks.data), i)
				}
				mks.data[i] = m
				mod *= m.test
				monkeyDesc = nil
			}
			continue
		}
		monkeyDesc = append(monkeyDesc, v)
	}
	if monkeyDesc != nil {
		i, m, err := parseMonkey(monkeyDesc)
		if err != nil {
			return nil, err
		}
		mks.data[i] = m
		mod *= m.test
	}
	mks.modul = mod
	return &mks, nil
}

func (m *Monkeys) round() {
	for _, item := range m.data {
		for l, to, err := item.getItem(m.lower); err != EOI; l, to, err = item.getItem(m.lower) {
			rest := l % worryLevel(m.modul)
			m.data[to].add(rest)
		}
	}
}

func (m *Monkeys) Run(i int) {
	for c := 0; c < i; c++ {
		m.round()
	}
}

func (m *Monkeys) sort() {
	sort.Slice(m.data, func(i, j int) bool { return m.data[i].cnt > m.data[j].cnt })
}

func (m *Monkeys) GetTopN(i int) []int {
	m.sort()
	var res []int
	if i > len(m.data) {
		i = len(m.data)
	}
	for x := 0; x < i; x++ {
		res = append(res, m.data[x].cnt)
	}
	return res
}
