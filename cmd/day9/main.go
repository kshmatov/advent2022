package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/rope"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: $ day6 <filename>")
		return
	}
	fn := os.Args[1]
	vals, err := lib.ReadStringsFile(fn)
	if err != nil {
		panic(err)
	}

	cmds, err := rope.BuildPath(vals)
	if err != nil {
		panic(err)
	}
	r := rope.NewRope()
	err = r.Travel(cmds)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Visited: %v\n", r.Points())

	l := rope.NewLongRope(10)
	err = l.Travel(cmds)
	if err != nil {
		panic(err)
	}
	fmt.Printf("LongRope Visited: %v\n", l.Points())
}
