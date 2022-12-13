package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/distress"
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

	pairs := distress.BuildPairs(vals)
	res := distress.Compare(pairs)
	fmt.Printf("Score: %v\n", res)

	vals = append(vals, "[[2]]")
	vals = append(vals, "[[6]]")
	l := distress.BuildList(vals)
	l.Sort()
	x2 := l.Find("[[2]]")
	x6 := l.Find("[[6]]")
	fmt.Printf("Marks: %v, %v -> %v\n", x2, x6, x2*x6)
}
