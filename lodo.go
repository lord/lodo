package main

import (
	"fmt"
	//"github.com/lord/lodo/breakout"
        "github.com/lord/lodo/test"
	"github.com/lord/lodo/core"
)

func main() {
	board := core.Board{}

	w := 35
	h := 42
	cols := 5
	rows := 8  
	err := board.Connect(w, h, cols, rows)
	defer board.Free()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {
		test.Run(board)
		board.Save() // draw the board
	}
}
