package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/signal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: $ day6 <filename>")
		return
	}
	fn := os.Args[1]
	val, err := lib.ReadStringFile(fn)
	if err != nil {
		panic(err)
	}

	start, err := signal.Start(val, 4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", start+4)
	start, err = signal.Start(val, 14)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", start+14)

}
