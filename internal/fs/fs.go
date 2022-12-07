package fs

import (
	"fmt"
	"strconv"
	"strings"
)

type command int

const (
	unknown = command(iota)
	ls
	cd
)

type Dir struct {
	name   string
	parent *Dir
	dirs   map[string]*Dir
	files  map[string]int
}

type Fs struct {
	root *Dir
	cur  *Dir
}

func newDir(name string, parent *Dir) *Dir {
	return &Dir{
		name:   name,
		parent: parent,
		dirs:   map[string]*Dir{},
		files:  map[string]int{},
	}
}

func (d *Dir) cd(path string) *Dir {
	if path == ".." {
		return d.parent
	}
	return d.addDir(path)
}

func (d *Dir) sizeDirs() map[string]int {
	res := map[string]int{}
	for n, d := range d.dirs {
		res[n] = d.Size(true)
	}
	return res
}

func (d *Dir) Size(children bool) int {
	i := 0
	for _, f := range d.files {
		i += f
	}
	if children {
		for _, s := range d.sizeDirs() {
			i += s
		}
	}
	return i
}

func (d *Dir) walk(f func(d *Dir)) {
	f(d)
	for _, c := range d.dirs {
		c.walk(f)
	}
}

func (d *Dir) addDir(name string) *Dir {
	if n, ok := d.dirs[name]; ok {
		return n
	}
	n := newDir(name, d)
	d.dirs[name] = n
	return n
}

func (d *Dir) addFile(name string, size int) *Dir {
	if _, ok := d.files[name]; ok {
		return d
	}
	d.files[name] = size
	return d
}

func (d *Dir) String() string {
	p := ""
	if d.parent != nil {
		p = d.parent.String()
	}
	p += d.name
	if p != "/" {
		p += "/"
	}
	return p
}

func newFS() *Fs {
	return &Fs{}
}

func (f *Fs) Cd(name string) {
	if f.root == nil {
		f.root = newDir(name, nil)
		f.cur = f.root
		return
	}
	if name == "/" {
		f.cur = f.root
		return
	}
	f.cur = f.cur.cd(name)
}

func (f *Fs) ls(s []string) {
	for _, item := range s {
		parts := strings.Split(item, " ")
		if parts[0] == "dir" {
			continue
		}
		size, _ := strconv.ParseInt(parts[0], 10, 64)
		f.cur.addFile(parts[1], int(size))
	}
}

func (f *Fs) Walk(a func(d *Dir)) {
	f.cur.walk(a)
}

func (f *Fs) Cur() *Dir {
	return f.cur
}

func BuildTree(s []string) *Fs {
	f := newFS()
	var lsBuf []string
	for _, line := range s {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line[0] == '$' {
			if len(lsBuf) > 0 {
				f.ls(lsBuf)
			}
			lsBuf = nil
			cmd, arg := getCommand(line)
			switch cmd {
			case cd:
				f.Cd(arg)
			case ls:
				lsBuf = nil
			}
		} else {
			lsBuf = append(lsBuf, line)
		}
	}
	if len(lsBuf) > 0 {
		f.ls(lsBuf)
	}
	return f
}

func getCommand(s string) (command, string) {
	parts := strings.Split(s, " ")
	switch parts[1] {
	case "cd":
		return cd, parts[2]
	case "ls":
		return ls, ""
	default:
		return unknown, ""
	}
}

func CalcAtMost(s []string, max int) int {
	l := map[string]int{}

	f := BuildTree(s)
	cnt := func(d *Dir) {
		s := d.Size(true)
		if s < max {
			l[d.String()] = s
		}
	}

	f.root.walk(cnt)

	sum := 0
	for i, c := range l {
		fmt.Printf("%v: %v\n", i, c)
		sum += c
	}
	return sum
}
