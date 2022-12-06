package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/rucksack"
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

	score, err := rucksack.CheckRuckSack(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", score)

	badge, err := rucksack.GetBadges(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("badges: %v\n", badge)
}
