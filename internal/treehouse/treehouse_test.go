package treehouse

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = strings.Split(`30373
25512
65332
33549
35390`, "\n")

func TestForest(t *testing.T) {
	f, err := NewForest(data)
	assert.Nil(t, err)
	assert.Equal(t, 5, f.cols())
	assert.Equal(t, 5, f.rows())
	assert.Equal(t, 16, f.perimeter())
}

func TestCheck(t *testing.T) {
	f, _ := NewForest(data)
	row := f.getRow(1)
	assert.Equal(t, []int{2, 5, 5, 1, 2}, row)
	res := check(row)
	assert.Equal(t, []int{0, 1}, res)
	row = reverse(row)
	assert.Equal(t, []int{2, 1, 5, 5, 2}, row)
	res = check(row)
	assert.Equal(t, []int{0, 2}, res)

	col := f.getCol(3)
	assert.Equal(t, []int{7, 1, 3, 4, 9}, col)
	res = check(col)
	assert.Equal(t, []int{0, 4}, res)
	col = reverse(col)
	assert.Equal(t, []int{9, 4, 3, 1, 7}, col)
	res = check(col)
	assert.Equal(t, []int{0}, res)

}

func TestVisibility(t *testing.T) {
	f, _ := NewForest(data)
	r := checkOuterVisibility(f)
	fmt.Printf("%v\n", r)
	i := Count(f)
	assert.Equal(t, 21, i)
}

func TestSplit(t *testing.T) {
	arr := []int{2, 5, 5, 1, 2}
	b, a := split(arr, 0)
	assert.Nil(t, b)
	assert.Equal(t, []int{5, 5, 1, 2}, a)

	b, a = split(arr, 1)
	assert.Equal(t, []int{2}, b)
	assert.Equal(t, []int{5, 1, 2}, a)

	b, a = split(arr, 2)
	assert.Equal(t, []int{5, 2}, b)
	assert.Equal(t, []int{1, 2}, a)

	b, a = split(arr, 3)
	assert.Equal(t, []int{5, 5, 2}, b)
	assert.Equal(t, []int{2}, a)

	b, a = split(arr, 4)
	assert.Equal(t, []int{1, 5, 5, 2}, b)
	assert.Nil(t, a)

}

func TestCheckHeight(t *testing.T) {
	i := checkHeight([]int{5, 2}, 5)
	assert.Equal(t, 1, i)
	i = checkHeight([]int{1, 2}, 5)
	assert.Equal(t, 2, i)
}

func TestCheckInner(t *testing.T) {
	i := chekInner([]int{2, 1, 5, 5, 2}, 2)
	assert.Equal(t, 2, i)
	i = chekInner([]int{3, 5, 3, 5, 3}, 2)
	assert.Equal(t, 1, i)
}

func TestCheckInnerVisibilityRow(t *testing.T) {
	a := []int{2, 1, 5, 5, 2}
	res := map[string]int{"0": 1, "1": 1, "2": 1, "3": 1, "4": 1}
	res = checkInnerVisibleRow(a, res, func(i int) string { return fmt.Sprint(i) })
	assert.Equal(t, 0, res["0"])
	assert.Equal(t, 1, res["1"])
	assert.Equal(t, 2, res["2"])
	assert.Equal(t, 1, res["3"])
	assert.Equal(t, 0, res["4"])
}

func TestInnerVisibility(t *testing.T) {
	f, _ := NewForest(data)
	res := checkInnerVisibility(f)
	assert.Equal(t, 4, res["1_2"])
	assert.Equal(t, 8, res["3_2"])
}
