package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/stock"
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

	score, err := stock.StockOps(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", score)

	ordered, err := stock.Stock9001(vals)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ordered Score: %v\n", ordered)

}
