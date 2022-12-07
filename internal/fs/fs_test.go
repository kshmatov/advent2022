package fs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = strings.Split(`
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`, "\n")

func TestFS(t *testing.T) {
	f := BuildTree(data)
	assert.NotNil(t, f.root)
	assert.NotNil(t, f.cur)
	f.Cd("/")
	assert.Equal(t, f.cur.String(), f.root.String())
	f.Cd("a")
	fmt.Printf("%v: %v: %v\n", f.cur, f.cur.files, f.cur.dirs)
	assert.Equal(t, "/a/", f.cur.String())
	assert.Equal(t, "/", f.cur.parent.String())
	f.Cd("e")
	assert.Equal(t, "/a/e/", f.cur.String())
	assert.Equal(t, 584, f.cur.Size(false))
	f.Cd("..")
	assert.Equal(t, "/a/", f.cur.String())
	assert.Equal(t, 94853, f.cur.Size(true))
	assert.Equal(t, 94853-584, f.cur.Size(false))
}

func TestCalcAtMost(t *testing.T) {
	s := CalcAtMost(data, 100000)
	assert.Equal(t, 95437, s)
}
