package treehouse

import (
	"fmt"
	"strconv"
	"strings"
)

type row []int
type forest []row

func NewForest(s []string) (forest, error) {
	var f forest
	for _, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		var r row
		for _, n := range line {
			h, err := strconv.ParseInt(string(n), 10, 64)
			if err != nil {
				return nil, err
			}
			r = append(r, int(h))
		}
		f = append(f, r)
	}
	return f, nil
}

func (f forest) perimeter() int {
	return len(f)*2 + len(f[0])*2 - 4
}

func (f forest) rows() int {
	return len(f)
}

func (f forest) cols() int {
	return len(f[0])
}

func (f forest) getRow(i int) []int {
	if i >= len(f) {
		return nil
	}
	return f[i]
}

func (f forest) getCol(i int) []int {
	if i >= len(f[0]) {
		return nil
	}
	res := []int{}
	for _, l := range f {
		res = append(res, l[i])
	}
	return res
}

func check(row []int) []int {
	var res []int
	maxHeight := -1
	for i, h := range row {
		if h > maxHeight {
			res = append(res, i)
			maxHeight = h
		}
	}
	return res
}

func getIdx(rowIdx, comIdx int) string {
	return fmt.Sprintf("%d_%d", rowIdx, comIdx)
}

func reverse(row []int) []int {
	var r []int
	for i := len(row) - 1; i >= 0; i-- {
		r = append(r, row[i])
	}
	return r
}

func checkOutherVisibleRow(row []int, checked map[string]bool, f func(int) string) map[string]bool {
	r := check(row)
	for _, v := range r {
		checked[f(v)] = true
	}
	r = check(reverse(row))
	for _, v := range r {
		v = len(row) - v - 1
		checked[f(v)] = true
	}
	return checked
}

func checkOuterVisibility(f forest) map[string]bool {
	checked := map[string]bool{}
	for rowI := 0; rowI < f.rows(); rowI++ {
		checked = checkOutherVisibleRow(f.getRow(rowI), checked, func(i int) string { return getIdx(rowI, i) })
	}

	for colI := 0; colI < f.cols(); colI++ {
		col := f.getCol(colI)
		checked = checkOutherVisibleRow(col, checked, func(i int) string { return getIdx(i, colI) })
	}
	return checked
}

func checkHeight(row []int, h int) int {
	if len(row) == 0 {
		return 0
	}
	for i, ch := range row {
		if ch >= h {
			return i + 1
		}
	}
	return len(row)
}

func split(row []int, i int) ([]int, []int) {
	if i == 0 {
		return nil, row[1:]
	}
	if i == len(row)-1 {
		return reverse(row[:len(row)-1]), nil
	}
	return reverse(row[:i]), row[i+1:]
}

func chekInner(row []int, pos int) int {
	before, after := split(row, pos)
	return checkHeight(before, row[pos]) * checkHeight(after, row[pos])
}

func checkInnerVisibleRow(row []int, checked map[string]int, fIdx func(int) string) map[string]int {
	for i := range row {
		r := chekInner(row, i)
		checked[fIdx(i)] *= r
	}
	return checked
}

func checkInnerVisibility(f forest) map[string]int {
	checked := map[string]int{}
	for i := range f {
		for y := range f[i] {
			checked[getIdx(i, y)] = 1
		}
	}

	for rowI := 0; rowI < f.rows(); rowI++ {
		checked = checkInnerVisibleRow(f.getRow(rowI), checked, func(i int) string { return getIdx(rowI, i) })
	}

	for colI := 0; colI < f.cols(); colI++ {
		col := f.getCol(colI)
		checked = checkInnerVisibleRow(col, checked, func(i int) string { return getIdx(i, colI) })
	}
	return checked
}

func Count(f forest) int {
	r := checkOuterVisibility(f)
	return len(r)
}

func InnerCount(f forest) int {
	r := checkInnerVisibility(f)
	max := 0
	for _, v := range r {
		if v > max {
			max = v
		}
	}
	return max
}
