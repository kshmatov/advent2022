package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/treehouse"
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

	f, err := treehouse.NewForest(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Visible: %v\n", treehouse.Count(f))
	fmt.Printf("Inner: %v\n", treehouse.InnerCount(f))
}
