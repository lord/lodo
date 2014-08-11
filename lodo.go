package main

import (
	"fmt"
	"github.com/lord/lodo/breakout"
	"github.com/lord/lodo/core"
)

func main() {
	board := core.Board{}

	w := 35
	h := 42
	cols := 5
	rows := 6
	err := board.Connect(w, h, cols, rows)
	defer board.Free()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	breakout.Run(&board)

	// strand := core.Strand{}
	// err := strand.Connect(2000)
	// defer strand.Free()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// rainbow.Run(&strand)
}
