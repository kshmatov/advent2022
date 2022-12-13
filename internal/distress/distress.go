package distress

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Signals []Comparable

func (p Signals) Compare() int {
	if len(p) != 2 {
		return -1
	}
	p1 := p[0]
	p2 := p[1]
	return p1.Compare(p2)
}

type Comparable interface {
	Compare(Comparable) int
	String() string
}

type arrArray struct {
	arrs []Comparable
}

func (a *arrArray) addInt(i string) *arrArray {
	if a == nil {
		a = &arrArray{}
	}
	i = strings.TrimSpace(i)

	v, err := strconv.Atoi(i)
	if err != nil {
		panic(err)
	}
	a.arrs = append(a.arrs, single(v))
	return a
}

func (a *arrArray) addArr(b *arrArray) *arrArray {
	if a == nil {
		return b
	}
	a.arrs = append(a.arrs, b)
	return a
}

func (a *arrArray) Compare(c Comparable) int {
	// fmt.Printf("compare <%v> vs <%v>\n", a, c)
	switch t := c.(type) {
	case single:
		return -1 * compareInt2Arr(t, a)
	case *arrArray:
		return a.compareArrs(t)
	}
	return 0
}

func (a *arrArray) compareArrs(c *arrArray) int {
	l1 := len(a.arrs)
	l2 := len(c.arrs)
	minLen := min(l1, l2)
	for i := 0; i < minLen; i++ {
		r := a.arrs[i].Compare(c.arrs[i])
		if r != 0 {
			return r
		}
	}
	if l1 > l2 {
		// fmt.Printf("arr1 : len %v\n", len(a.arrs))
		return 1
	}
	if l1 < l2 {
		// fmt.Printf("arr2 : len %v\n", len(a.arrs))
		return -1
	}
	return 0

}

func (a *arrArray) String() string {
	res := "["
	for _, v := range a.arrs {
		res += v.String() + ","
	}
	res += "]"
	return res
}

type single int

func (s single) Compare(c Comparable) int {
	// fmt.Printf("compare <%v> vs <%v>\n", s, c)
	switch t := c.(type) {
	case single:
		return compareInt(s, t)
	case *arrArray:
		return compareInt2Arr(s, t)
	}
	return 0
}

func compareInt(i1, i2 single) int {
	if i1 > i2 {
		return 1
	}
	if i1 < i2 {
		return -1
	}
	return 0
}

func (s single) String() string {
	return strconv.Itoa(int(s))
}

func compareInt2Arr(i single, a *arrArray) int {
	if len(a.arrs) == 0 {
		// fmt.Printf("single to arr: len %v\n", len(a.arrs))
		return 1
	}
	if r := i.Compare(a.arrs[0]); r == 0 {
		if len(a.arrs) == 1 {
			return 0
		}
		// fmt.Printf("single to arr: len %v\n", len(a.arrs))
		return -1
	} else {
		return r
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func findPair(s string) (int, int) {
	cnt := 0
	start := -1
	for i, c := range s {
		switch c {
		case '[':
			if start < 0 {
				start = i
			}
			cnt++
		case ']':
			cnt--
			if cnt == 0 {
				return start, i
			}
		}
	}
	return -1, -1
}

func buildArr(s string) *arrArray {
	s = strings.TrimSpace(s)
	var out *arrArray

	// [xxx] -> xxx
	start, end := findPair(s)
	s = string(s[start:end])

	pos := 0
	cur := ""
	for pos < len(s) {
		if s[pos] == '[' {
			if cur != "" {
				out = out.addInt(cur)
				cur = ""
			}
			// x,y, [a,s,d],f -> [a,s,d]
			start, end := findPair(string(s[pos:]))
			start += pos
			end += pos
			// [x, y] -> x, y
			// [] - > empty node
			if end > start+1 {
				s1 := string(s[start : end+1])
				if s == s1 {
					panic(fmt.Sprintf("income: <%v>, cut: <%v>", s, s1))
				}
				out = out.addArr(buildArr(s1))
			} else {
				out = out.addArr(&arrArray{})
			}
			pos = end + 2
			continue
		}

		if s[pos] == ',' {
			if cur != "" {
				out = out.addInt(cur)
				cur = ""
			}
		} else {
			cur += string(s[pos])
		}
		pos++
	}
	if cur != "" {
		out = out.addInt(cur)
	}
	return out
}

func BuildPairs(s []string) []Signals {
	var pair Signals
	var res []Signals
	for _, line := range s {
		line = strings.TrimSpace(line)
		if line == "" {
			if len(pair) > 0 {
				res = append(res, pair)
				pair = nil
			}
			continue
		}
		pair = append(pair, buildArr(line))
	}
	if len(pair) > 0 {
		res = append(res, pair)
	}
	return res
}

func Compare(p []Signals) int {
	r := 0
	for i, pair := range p {
		idx := i + 1
		c := pair.Compare()
		// fmt.Printf("%v %v: %v\n%v\n\n", idx, c, pair[0], pair[1])
		if c < 1 {
			r += idx
		}
	}
	return r
}

func BuildList(s []string) Signals {
	var res Signals
	for _, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		res = append(res, buildArr(line))
	}
	return res
}

func (s Signals) Sort() {
	sort.Slice(s, func(i, j int) bool { return s[i].Compare(s[j]) < 0 })
}

func (s Signals) Find(n string) int {
	x := buildArr(n)
	for i, v := range s {
		if x.Compare(v) == 0 {
			return i + 1
		}
	}
	return -1
}
