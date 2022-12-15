package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kshmatov/advent2022/internal/beacon"
	"github.com/kshmatov/advent2022/internal/lib"
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

	b := beacon.BuildList(vals)
	y := 2000000

	t := time.Now()
	m := beacon.NewRow()
	for _, i := range b {
		m.CalcFreeSpaces(i, y)
	}

	fmt.Printf("Score: %v (%v)\n", m.Len(), time.Since(t))
	t = time.Now()
	p := beacon.BuildList(vals)
	bb := beacon.BuildBorders(true, p)
	res := beacon.FilterBorders(0, 20, bb)
	fmt.Printf("Score (%v): %v . %v", res, res.X*4000000+res.Y, time.Since(t))
}
