package main

import (
	"fmt"
	"github.com/lord/lodo/core"
	// "github.com/lord/lodo/rainbow_board"
	"github.com/lord/lodo/test"
)

func main() {
	w := 35
	h := 42
	cols := 5
	rows := 6
	board, err := core.MakeBoard(w, h, cols, rows)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer board.Free()

	// strand := core.Strand{}
	// err := strand.Connect(2000)
	// defer strand.Free()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	//RunServer(&board)
	test.Run(board)
}
