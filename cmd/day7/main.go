package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/kshmatov/advent2022/internal/fs"
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

	f := fs.BuildTree(vals)
	f.Cd("/")
	sum := 0
	f.Walk(func(d *fs.Dir) {
		s := d.Size(true)
		if s < 100000 {
			// fmt.Printf("%v\t%v\n", d, s)
			sum += s
		}
	})

	fmt.Printf("Score: %v\n", sum)

	f.Cd("/")
	total := 70000000
	need := 30000000

	var arr []int
	f.Walk(func(d *fs.Dir) {
		arr = append(arr, d.Size(true))
	})

	sort.Ints(arr)

	free := total - f.Cur().Size(true)
	fmt.Printf("free: %v, need more: %v\n", free, need-free)
	for _, s := range arr {
		if free+s >= need {
			fmt.Printf("Delete %v\n", s)
			break
		}
	}
}
