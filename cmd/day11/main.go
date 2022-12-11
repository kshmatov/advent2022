package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/monkey"
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

	m, err := monkey.BuildMokeyList(vals, 8, 3)
	if err != nil {
		panic(err)
	}
	m.Run(20)
	cnt := m.GetTopN(2)
	fmt.Printf("Stat: %v\n", cnt[0]*cnt[1])

	m, err = monkey.BuildMokeyList(vals, 8, 0)
	if err != nil {
		panic(err)
	}
	m.Run(10000)
	cnt = m.GetTopN(2)
	fmt.Printf("Stat: %v\n", cnt[0]*cnt[1])

}
