package main

import (
	"fmt"
	"github.com/lord/lodo/core"
	"github.com/lord/lodo/rainbow_board"
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

	rainbowBoard.Run(&board)

	// strand := core.Strand{}
	// err := strand.Connect(2000)
	// defer strand.Free()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// rainbow.Run(&strand)
}
