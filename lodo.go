package main

import (
	"fmt"
	"github.com/lord/lodo/core"
	"github.com/lord/lodo/rainbow"
)

func main() {
	// board := core.Board{}

	// w := 35
	// h := 42
	// cols := 5
	// rows := 8
	// err := board.Connect(w, h, cols, rows)
	// defer board.Free()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	strand := core.Strand{}

	w := 35
	h := 42
	err := strand.Connect(w * h)
	defer strand.Free()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	rainbow.Run(&strand)
}
