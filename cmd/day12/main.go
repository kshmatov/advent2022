package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/climber"
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

	i := climber.Travel(vals)
	fmt.Printf("Stat: %v\n", i)
	i = climber.FindBest(vals)
	fmt.Printf("Lowest: %v\n", i)
}
