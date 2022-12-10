package videosignal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type CPU struct {
	debug      bool
	register   int
	cycle      int
	mem        map[int]int
	controller *controller
	video      *Screen
}

func (c *CPU) addx(i int) {
	c.cycle += 2
	if v := c.controller.check(2); v > 0 {
		c.mem[v] = c.register
	}
	if c.video != nil {
		c.video.tick(c.register)
		c.video.tick(c.register)
	}
	c.register += i
}

func (c *CPU) noop() {
	c.cycle += 1
	if v := c.controller.check(1); v > 0 {
		c.mem[v] = c.register
	}
	if c.video != nil {
		c.video.tick(c.register)
	}
}

func (c *CPU) SetDebug(debug bool) *CPU {
	c.debug = debug
	return c
}

func (c *CPU) SetScreen(s *Screen) {
	c.video = s
}

func (c *CPU) ControlSum() int {
	cs := 0
	for k, v := range c.mem {
		cs += k * v
	}
	return cs
}

func (c *CPU) Exec(cmd string, arg ...string) error {
	switch cmd {
	case "noop":
		c.noop()
	case "addx":
		if len(arg) != 1 {
			return fmt.Errorf("addx expected 1 argument, %v given", len(arg))
		}
		a, err := strconv.ParseInt(arg[0], 10, 64)
		if err != nil {
			return fmt.Errorf("addx expect number as argument, got %#v", arg[0])
		}
		c.addx(int(a))
	}
	if c.debug {
		fmt.Printf("%v %v\treg %4v cyc %4v\tmem: <%v>\n", cmd, arg, c.register, c.cycle, c.mem)
	}
	return nil
}

func (c *CPU) Mem() map[int]int {
	return c.mem
}

func NewCPU(start, step int) *CPU {
	return &CPU{
		register:   1,
		cycle:      0,
		controller: newController(start, step),
		mem:        map[int]int{},
	}
}

type controller struct {
	inc   int
	cur   int
	cycle int
}

func (c *controller) check(cycle int) int {
	c.cycle += cycle
	ret := 0
	if c.cycle >= c.cur {
		ret = c.cur
		c.cur += c.inc
	}
	return ret
}

func newController(start, step int) *controller {
	return &controller{
		inc:   step,
		cur:   start,
		cycle: 0,
	}
}

func Exec(s []string, c *CPU) error {
	for i, line := range s {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		cmd := parts[0]
		err := c.Exec(cmd, parts[1:]...)
		if err != nil {
			return errors.Wrapf(err, "on line: %v", i)
		}
	}
	return nil
}

type Screen struct {
	width  int
	heigth int
	x      int
	y      int
}

func (s *Screen) tick(i int) {
	if s.x >= s.width {
		fmt.Print("\n")
		s.x = 0
		s.y++
	}
	if s.y >= s.heigth {
		s.y = 0
		fmt.Println("- new screen - new screen - new screen -")
	}

	if i <= s.x+1 && i >= s.x-1 {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
	s.x++
}

func NewScreen(h, w int) *Screen {
	return &Screen{
		width:  w,
		heigth: h,
		x:      0,
		y:      0,
	}
}
