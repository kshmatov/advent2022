package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/cleaners"
	"github.com/kshmatov/advent2022/internal/lib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: $ day1 <filename>")
		return
	}
	fn := os.Args[1]
	vals, err := lib.ReadStringFile(fn)
	if err != nil {
		panic(err)
	}

	score, err := cleaners.CheckInclusion(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", score)

	score, err = cleaners.CheckOverlaption(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Overlaped: %v\n", score)

}
