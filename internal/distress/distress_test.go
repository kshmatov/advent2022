package distress

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = strings.Split(`[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`, "\n")

func TestBuildArr(t *testing.T) {
	a := buildArr("[]")
	assert.Equal(t, 0, len(a.arrs))

	a = buildArr("[1,1,3,1,1]")
	assert.Equalf(t, 5, len(a.arrs), "%v", a)

	a = buildArr("[[1],[2,3,4]]")
	if !assert.Equal(t, 2, len(a.arrs)) {
		return
	}
	assert.Equal(t, 1, len(a.arrs[0].(*arrArray).arrs))
	assert.Equal(t, 3, len(a.arrs[1].(*arrArray).arrs))
	// [ [ [], [9],4 ] , [ [2] ] ]
	// [[9,[]],[2],[[8,9,5,7],[[],[10]],7,[[8,1,6,3,8],[1],[6,8,10],2,[8,5,7,4,6]],[]]]
	a = buildArr("[[[],[9],4],[[2]]]")
	assert.Equal(t, 2, len(a.arrs), a.String())
	v := a.arrs[0].(*arrArray)
	assert.Equal(t, 3, len(v.arrs), v.String())
}

func TestCompare(t *testing.T) {
	a1 := buildArr("[1,1,3,1,1]")
	a2 := buildArr("[1,1,5,1,1]")
	assert.Equal(t, -1, a1.Compare(a2))
	assert.Equal(t, 1, a2.Compare(a1))
	assert.Equal(t, -1, a1.Compare(single(2)))
	assert.Equal(t, 1, a1.Compare(single(0)))
	assert.Equal(t, 1, a1.Compare(single(1)))

	a1 = buildArr("[[1],[2,3,4]]")
	a2 = buildArr("[[1],4]")
	assert.Equal(t, -1, a1.Compare(a2))
	assert.Equal(t, 1, a2.Compare(a1))

	a1 = buildArr("[9]")
	a2 = buildArr("[[8,7,6]]")
	assert.Equal(t, 1, a1.Compare(a2))
	assert.Equal(t, -1, a2.Compare(a1))

	a1 = buildArr("[[4,4],4,4]")
	a2 = buildArr("[[4,4],4,4,4]")
	assert.Equal(t, -1, a1.Compare(a2))
	assert.Equal(t, 1, a2.Compare(a1))

	a1 = buildArr("[7,7,7,7]")
	a2 = buildArr("[7,7,7]")
	assert.Equal(t, 1, a1.Compare(a2))
	assert.Equal(t, -1, a2.Compare(a1))

	a1 = buildArr("[]")
	a2 = buildArr("[3]")
	assert.Equal(t, -1, a1.Compare(a2))
	assert.Equal(t, 1, a2.Compare(a1))

	a1 = buildArr("[[[]]]")
	a2 = buildArr("[[]]")
	assert.Equal(t, 1, a1.Compare(a2))
	assert.Equal(t, -1, a2.Compare(a1))

	a1 = buildArr("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	a2 = buildArr("[1,[2,[3,[4,[5,6,0]]]],8,9]")
	assert.Equal(t, 1, a1.Compare(a2))
	assert.Equal(t, -1, a2.Compare(a1))

	a1 = buildArr("[[],1]")
	a2 = buildArr("[[],2]")
	assert.Equal(t, -1, a1.Compare(a2))
}

func TestPairs(t *testing.T) {
	data := strings.Split(`[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`, "\n")

	pairs := BuildPairs(data)
	assert.Equal(t, 13, Compare(pairs))
}

func TestPacket(t *testing.T) {
	l := BuildList(data)
	i2 := buildArr("[[2]]")
	i6 := buildArr("[[6]]")
	l = append(l, i2)
	l = append(l, i6)

	l.Sort()
	x2 := l.Find("[[2]]")
	assert.Equal(t, 10, x2)
	x6 := l.Find("[[6]]")
	assert.Equal(t, 14, x6)
}
