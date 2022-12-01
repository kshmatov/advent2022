package caloriesqueue

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ElfCalories int

type ElfQueue []ElfCalories

type queueRes []int

func (r queueRes) add(i int) bool {
	found := false
	for pos, val := range r {
		if i > val {
			r[pos] = i
			i = val
			found = true
		}
	}
	return found
}

func NewEflQueue(s []string) (ElfQueue, error) {
	if len(s) == 0 {
		return nil, errors.New("empty data")
	}
	var res ElfQueue
	var cur ElfCalories
	for line, item := range s {
		item := strings.TrimSpace(item)
		if item == "" {
			res = append(res, cur)
			cur = 0
			continue
		}
		iVal, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("line: %-3v is not integer value: <%v>", line, item)
		}
		cur += ElfCalories(iVal)
	}
	if cur > 0 {
		res = append(res, cur)
	}
	return res, nil
}

func (e ElfQueue) HasMax() (int, error) {
	if e == nil {
		return 0, errors.New("queue is not created")
	}
	maxVal := ElfCalories(-1)
	maxIdx := -1
	for i, v := range e {
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}
	return maxIdx, nil
}

func (e ElfQueue) SummTopN(n uint) (int, error) {
	if n == 1 {
		i, err := e.HasMax()
		return int(e[i]), err
	}
	qr := make(queueRes, n)
	for _, v := range e {
		qr.add(int(v))
	}
	sum := 0
	for _, i := range qr {
		sum += i
	}
	return sum, nil
}
