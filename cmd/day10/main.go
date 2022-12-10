package main

import (
	"fmt"
	"os"

	"github.com/kshmatov/advent2022/internal/lib"
	"github.com/kshmatov/advent2022/internal/videosignal"
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

	c := videosignal.NewCPU(20, 40)
	err = videosignal.Exec(vals, c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ControlSun: %v\n", c.ControlSum())

	outC := videosignal.NewCPU(0, 40)
	vid := videosignal.NewScreen(6, 40)
	outC.SetScreen(vid)
	videosignal.Exec(vals, outC)
}
