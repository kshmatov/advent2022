package climber

import (
	"fmt"
	"strings"
)

type row []byte

type point struct {
	x int
	y int
}

func (p point) idx() string {
	return fmt.Sprintf("%v_%v", p.x, p.y)
}

type mountainMap struct {
	data    []row
	start   point
	end     point
	visited map[string]struct{}
}

func (m *mountainMap) isEnd(p point) bool {
	return p.x == m.end.x && p.y == m.end.y
}

func (m *mountainMap) isLowest(p point) bool {
	return m.getHeight(p) == 'a'
}

func (m *mountainMap) getHeight(p point) byte {
	if p.y >= len(m.data) || p.y < 0 {
		return 0
	}
	row := m.data[p.y]
	if p.x >= len(row) || p.x < 0 {
		return 0
	}
	return row[p.x]
}

func (m *mountainMap) checkMove(c byte, p point) bool {
	n := m.getHeight(p)
	if n == 0 {
		return false
	}
	if n == 'S' {
		n = 'a'
	}
	if c == 'E' {
		c = 'z'
	}
	return c-1 <= n
}

func (m *mountainMap) getWays(p point) []point {
	c := m.getHeight(p)
	top := point{x: p.x, y: p.y - 1}
	rigth := point{x: p.x + 1, y: p.y}
	bottom := point{x: p.x, y: p.y + 1}
	left := point{x: p.x - 1, y: p.y}

	var res []point
	if m.checkMove(c, top) {
		res = append(res, top)
	}
	if m.checkMove(c, rigth) {
		res = append(res, rigth)
	}
	if m.checkMove(c, bottom) {
		res = append(res, bottom)
	}
	if m.checkMove(c, left) {
		res = append(res, left)
	}
	return res
}

type track struct {
	len      int
	visited  map[string]struct{}
	position point
}

func (t *track) duplicate() *track {
	m := map[string]struct{}{}
	for k := range t.visited {
		m[k] = struct{}{}
	}
	return &track{
		len:      t.len,
		visited:  m,
		position: point{x: t.position.x, y: t.position.y},
	}
}

func (t *track) move(m *mountainMap) []*track {
	var res []*track
	enabled := m.getWays(t.position)
	for _, way := range enabled {
		idx := way.idx()
		if _, ok := m.visited[idx]; ok {
			continue
		}

		m.visited[idx] = struct{}{}
		new := t.duplicate()
		new.position = way
		new.len++
		new.visited[idx] = struct{}{}
		res = append(res, new)
	}
	return res
}

func travel(m *mountainMap, checker func(point) bool) int {
	tList := []*track{{
		position: m.start,
		visited:  map[string]struct{}{m.start.idx(): {}},
		len:      0,
	}}
	for {
		var new []*track
		for _, tr := range tList {
			l := tr.move(m)
			for _, item := range l {
				if checker(item.position) {
					return item.len
				}
				new = append(new, item)
			}
		}
		tList = new
	}
}

func buildMap(s []string) *mountainMap {
	m := mountainMap{
		visited: make(map[string]struct{}),
	}
	for y, line := range s {
		x := strings.IndexByte(line, 'S')
		if x >= 0 {
			m.end = point{x: x, y: y}
		}
		x = strings.IndexByte(line, 'E')
		if x >= 0 {
			m.start = point{x: x, y: y}
		}
		m.data = append(m.data, []byte(line))
	}
	return &m
}

func Travel(s []string) int {
	m := buildMap(s)
	return travel(m, m.isEnd)
}

func FindBest(s []string) int {
	m := buildMap(s)
	return travel(m, m.isLowest)
}
