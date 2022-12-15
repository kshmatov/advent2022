package beacon

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("x: %v, y: %v", p.X, p.Y)
}

type Points map[Point]struct{}
type Figures map[Ping]Points

type Ping struct {
	sensor Point
	signal Point
}

type Row struct {
	data    map[Point]struct{}
	beacons map[Point]struct{}
}

func NewRow() *Row {
	return &Row{
		data:    map[Point]struct{}{},
		beacons: map[Point]struct{}{},
	}
}

func (b Ping) freeDistance() int {
	return abs(b.sensor.X-b.signal.X) + abs(b.sensor.Y-b.signal.Y)
}

func (b Ping) Borders() Points {
	p := make(Points)
	d := b.freeDistance() + 1
	center := b.sensor
	for y := center.Y - d; y <= center.Y+d; y++ {
		extra := d - b.yRange(y)
		x1 := center.X + extra
		x2 := center.X - extra
		p[Point{X: x1, Y: y}] = struct{}{}
		p[Point{X: x2, Y: y}] = struct{}{}
	}
	return p
}

func (b Ping) yRange(y int) int {
	return abs(y - b.sensor.Y)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func BuildList(s []string) []Ping {
	var b []Ping
	for _, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Printf("Bad string: %v\n", line)
			continue
		}
		sensor := extractPoint(parts[0])
		signal := extractPoint(parts[1])
		b = append(b, Ping{sensor: sensor, signal: signal})
	}
	return b
}

func extractPoint(s string) Point {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	maxIdx := len(parts) - 1
	xStr := parts[maxIdx-1]
	yStr := parts[maxIdx]
	xStr = string(xStr[2 : len(xStr)-1])
	yStr = string(yStr[2:])
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	return Point{X: x, Y: y}
}

func (r *Row) CalcFreeSpaces(b Ping, y int) int {
	if r.data == nil {
		r.data = map[Point]struct{}{}
	}
	if r.beacons == nil {
		r.beacons = map[Point]struct{}{}
	}
	extra := b.freeDistance() - b.yRange(y)
	if extra < 0 {
		return 0
	}
	i := 0
	for x := b.sensor.X - extra; x <= b.sensor.X+extra; x++ {
		r.data[Point{X: x, Y: y}] = struct{}{}
	}
	if b.signal.Y == y {
		r.beacons[b.signal] = struct{}{}
	}
	return i
}

func (r *Row) Len() int {
	return len(r.data) - len(r.beacons)
}

func BuildBorders(out bool, b []Ping) Figures {
	res := Figures{}
	for _, i := range b {
		res[i] = i.Borders()
	}
	return res
}

func FilterBorders(min, max int, f Figures) Point {
	r := map[Point]int{}
	for _, i := range f {
		for p := range i {
			if p.X >= 0 && p.X <= 4000000 && p.Y >= 0 && p.Y <= 4000000 {
				r[p] = r[p] + 1
			}
		}
	}
	// fmt.Printf("%+v\n", r)
	for k := range r {
		if checkInside(k, f) {
			delete(r, k)
		}
	}
	// fmt.Printf("%+v\n", r)
	for k, v := range r {
		if v >= 4 {
			return k
		}
	}
	return Point{}
}

func checkInside(p Point, f Figures) bool {
	for k := range f {
		d := k.freeDistance()
		if p.X < k.sensor.X-d || p.X > k.sensor.X+d ||
			p.Y < k.sensor.Y-d || p.Y > k.sensor.Y+d {
			continue
		}
		extra := k.freeDistance() - k.yRange(p.Y)
		if p.X >= k.sensor.X-extra && p.X <= k.sensor.X+extra {
			// if p.X == 14 && p.Y == 11 {
			// 	// fmt.Printf("%+v - out by %+v\n", p, k)
			// }
			return true
		}
	}
	return false
}
