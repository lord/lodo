package main

import (
	"flag"
	"fmt"
	"github.com/lord/lodo/breakout"
	"github.com/lord/lodo/core"
	"github.com/lord/lodo/rainbow_board"
	"github.com/lord/lodo/test"
)

var gameMode = flag.String(
	"mode",
	"rainbow-board",
	"Selects the game to run. Options are 'test', 'rainbow-board', and 'breakout'.",
)

func main() {
	flag.Parse()

	board, err := core.MakeBoard()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer board.Free()

	switch *gameMode {
	case "rainbow-board":
		rainbowBoard.Run(board)
	case "test":
		test.Run(board)
	case "breakout":
		breakout.Run(board)
	}
}
