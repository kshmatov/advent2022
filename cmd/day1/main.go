package main

import (
	"fmt"
	"os"

	caloriesqueue "github.com/kshmatov/advent2022/internal/caloriesQueue"
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
	q, err := caloriesqueue.NewEflQueue(vals)
	if err != nil {
		panic(err)
	}
	// 1 day 1 task
	maxIDx, err := q.HasMax()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v elf has <%v>\n", maxIDx, q[maxIDx])
	// 1 day 2 task
	sum, _ := q.SummTopN(3)
	fmt.Printf("3 top sum: %v\n", sum)
}
