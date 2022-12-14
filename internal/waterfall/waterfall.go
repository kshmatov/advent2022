package waterfall

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type fillament byte

const (
	void = fillament(iota)
	stone
	sand
)

type Point struct {
	x int
	y int
}

func (p Point) isLower(x Point) bool {
	return p.y > x.y
}

type Blocks struct {
	b           map[Point]fillament
	lowestPoint Point
}

var (
	ErrStoped = errors.New("stopped")
	ErrVoid   = errors.New("void")
)

func fillToByte(f fillament) byte {
	switch f {
	case sand:
		return '*'
	case stone:
		return '#'
	default:
		return '.'
	}
}

func (b Blocks) Print() {
	p := make([]Point, 0, len(b.b))
	lp := Point{500, 0}
	rp := Point{500, 0}
	for k := range b.b {
		p = append(p, k)
		if k.x < lp.x {
			lp.x = k.x
		}
		if k.y > lp.y {
			lp.y = k.y
		}
		if k.x > rp.x {
			rp.x = k.x
		}
	}

	widht := rp.x - lp.x + 1

	sort.Slice(p, func(x, y int) bool { return p[x].y < p[y].y })
	curLevel := 0
	row := getRow(widht)
	for _, x := range p {
		if x.y != curLevel {
			fmt.Println(string(row))
			row = getRow(widht)
			curLevel = x.y
		}
		row[x.x-lp.x] = fillToByte(b.b[x])
	}
	fmt.Println(string(row))
}

func getRow(w int) []byte {
	b := make([]byte, w)
	for i := range b {
		b[i] = '.'
	}
	return b
}

func (b Blocks) Path(p Point) (Point, error) {
	np := Point{x: p.x, y: p.y + 1}
	x, ok := b.b[np]
	if !ok || x == void {
		if np.isLower(b.lowestPoint) {
			return np, ErrVoid
		}
		return np, nil
	}

	np.x -= 1
	x, ok = b.b[np]
	if !ok || x == void {
		return np, nil
	}
	np.x += 2
	x, ok = b.b[np]
	if !ok || x == void {
		return np, nil
	}
	return p, ErrStoped
}

func (b Blocks) PathToFloor(p Point) (Point, error) {
	x, ok := b.b[p]
	if ok && x != void {
		return p, ErrVoid
	}
	np := Point{x: p.x, y: p.y + 1}
	if np.y == b.lowestPoint.y+2 {
		b.b[Point{x: p.x, y: b.lowestPoint.y + 2}] = stone
		return p, ErrStoped
	}

	x, ok = b.b[np]
	if !ok || x == void {
		return np, nil
	}

	np.x -= 1
	x, ok = b.b[np]
	if !ok || x == void {
		return np, nil
	}
	np.x += 2
	x, ok = b.b[np]
	if !ok || x == void {
		return np, nil
	}
	return p, ErrStoped
}

func FillMap(s []string) Blocks {
	m := Blocks{b: make(map[Point]fillament)}
	for _, line := range s {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		coords := strings.Split(line, "->")
		var pack []Point
		for _, x := range coords {
			p := makePoint(x)
			pack = append(pack, p)
		}
		for k, v := range makeBlocks(pack...) {
			m.b[k] = v
			if k.y > m.lowestPoint.y {
				m.lowestPoint = k
			}
		}
	}
	return m
}

func minmax(p1, p2 Point) (Point, Point) {
	if p1.x > p2.x {
		return p2, p1
	}
	if p1.y > p2.y {
		return p2, p1
	}
	return p1, p2
}

func makePoint(s string) Point {
	xs := strings.Split(strings.TrimSpace(s), ",")
	p := Point{}
	p.x, _ = strconv.Atoi(xs[0])
	p.y, _ = strconv.Atoi(xs[1])
	return p
}

func makeBlocks(p ...Point) map[Point]fillament {
	b := make(map[Point]fillament)
	start := p[0]
	for _, end := range p[1:] {
		p1, p2 := minmax(start, end)
		if p1.x < p2.x {
			for i := p1.x; i <= p2.x; i++ {
				b[Point{x: i, y: p1.y}] = stone
			}
		} else {
			for i := p1.y; i <= p2.y; i++ {
				b[Point{x: p1.x, y: i}] = stone
			}
		}
		start = end
	}
	return b
}

func Sand(b Blocks) (Blocks, int) {
	cnt := 0
	for {
		p, err := grain(b)
		if err == ErrStoped {
			b.b[p] = sand
			cnt++
		}
		if err == ErrVoid {
			return b, cnt
		}
	}
}

func grain(b Blocks) (Point, error) {
	p := Point{x: 500, y: 0}
	var err error
	for {
		p, err = b.Path(p)
		if err != nil {
			return p, err
		}
	}
}

func SandFloor(b Blocks) (Blocks, int) {
	cnt := 0
	for {
		p, err := grainFloor(b)
		if err == ErrStoped {
			b.b[p] = sand
			cnt++
		}
		if err == ErrVoid {
			return b, cnt
		}
	}
}

func grainFloor(b Blocks) (Point, error) {
	p := Point{x: 500, y: 0}
	var err error
	for {
		p, err = b.PathToFloor(p)
		if err != nil {
			return p, err
		}
	}
}
