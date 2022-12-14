package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/waterfall"
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

	b := waterfall.FillMap(vals)
	b, cnt := waterfall.Sand(b)
	b.Print()
	fmt.Printf("Score: %v\n", cnt)

	// data := strings.Split(`498,4 -> 498,6 -> 496,6
	// 503,4 -> 502,4 -> 502,9 -> 494,9`, "\n")
	b = waterfall.FillMap(vals)
	b, cnt = waterfall.SandFloor(b)
	b.Print()
	fmt.Printf("Score: %v\n", cnt)
}
