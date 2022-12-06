package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/gameday"
	"github.com/kshmatov/advent2022/internal/lib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: $ day1 <filename>")
		return
	}
	fn := os.Args[1]
	vals, err := lib.ReadStringsFile(fn)
	if err != nil {
		panic(err)
	}

	score, err := gameday.Calculate(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", score)

	score, err = gameday.ReverseCalculate(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Reverse score: %v\n", score)
}
