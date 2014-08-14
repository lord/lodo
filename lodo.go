package main

import (
	"fmt"
	"github.com/lord/lodo/core"
	"github.com/lord/lodo/rainbow_board"
	// "github.com/lord/lodo/test"
)

func main() {
	board, err := core.MakeBoard()
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
	rainbowBoard.Run(board)
}
