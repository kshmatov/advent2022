package rope

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type history struct {
	d map[string]bool
}

func newHistory() *history {
	return &history{
		d: make(map[string]bool),
	}
}

func (h *history) add(x, y int) {
	h.d[fmt.Sprintf("%v_%v", x, y)] = true
}

func (h *history) length() int {
	return len(h.d)
}

type knot struct {
	x int
	y int
}

type Rope struct {
	head knot
	tail knot
	path *history
}

func NewRope() *Rope {
	h := newHistory()
	h.add(0, 0)
	return &Rope{
		head: knot{0, 0},
		tail: knot{0, 0},
		path: h,
	}
}

func (r *Rope) move(d string) error {
	defer func() {
		r.path.add(r.tail.x, r.tail.y)
	}()

	switch d {
	case "U":
		r.head.y++
	case "D":
		r.head.y--
	case "L":
		r.head.x--
	case "R":
		r.head.x++
	default:
		return errors.New("unknown direction")
	}
	if abs(r.head.x-r.tail.x) > 1 {
		r.tail.y = r.head.y
		if r.head.x > r.tail.x {
			r.tail.x = r.head.x - 1
		} else {
			r.tail.x = r.head.x + 1
		}
	}
	if abs(r.head.y-r.tail.y) > 1 {
		r.tail.x = r.head.x
		if r.head.y > r.tail.y {
			r.tail.y = r.head.y - 1
		} else {
			r.tail.y = r.head.y + 1
		}
	}
	return nil
}

func (r *Rope) Travel(commsnds []Command) error {
	for _, c := range commsnds {
		for i := 0; i < c.steps; i++ {
			err := r.move(c.direction)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Rope) Points() int {
	return r.path.length()
}

type LongRope struct {
	knots []knot
	path  *history
}

func NewLongRope(i int) *LongRope {
	h := newHistory()
	h.add(0, 0)
	return &LongRope{
		knots: make([]knot, i),
		path:  h,
	}
}

func (l *LongRope) move(d string) error {
	last := len(l.knots) - 1
	defer func() {
		l.path.add(l.knots[last].x, l.knots[last].y)
	}()

	switch d {
	case "U":
		l.knots[0].y++
	case "D":
		l.knots[0].y--
	case "L":
		l.knots[0].x--
	case "R":
		l.knots[0].x++
	default:
		return errors.New("unknown direction")
	}

	for i := 1; i < len(l.knots); i++ {
		h := abs(l.knots[i-1].x - l.knots[i].x)
		v := abs(l.knots[i-1].y - l.knots[i].y)

		k := h - 1
		if (h > 1) || (h == 1 && v == 2) {
			if l.knots[i-1].x > l.knots[i].x {
				l.knots[i].x = l.knots[i-1].x - k
			} else {
				l.knots[i].x = l.knots[i-1].x + k
			}
		}

		k = v - 1
		if (v > 1) || (v == 1 && h == 2) {
			if l.knots[i-1].y > l.knots[i].y {
				l.knots[i].y = l.knots[i-1].y - k
			} else {
				l.knots[i].y = l.knots[i-1].y + k
			}
		}
	}
	return nil
}

func (l *LongRope) Travel(cmds []Command) error {
	for _, c := range cmds {
		for i := 0; i < c.steps; i++ {
			err := l.move(c.direction)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *LongRope) Points() int {
	return l.path.length()
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type Command struct {
	direction string
	steps     int
}

func BuildPath(s []string) ([]Command, error) {
	var res []Command
	for _, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		i, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, Command{direction: parts[0], steps: int(i)})
	}
	return res, nil
}
