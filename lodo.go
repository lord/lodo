package main

import "fmt"

import "time"

//import "math"

const rows int = 5
const cols int = 4

func main() {
	start := time.Now() // starting time in ns
	board := Board{}

	w := 35
	h := 28
	err := board.Connect(w, h, cols, rows)
	defer board.Free()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// for {
	// 	for y := 0; y < h; y++ {
	// 		for x := 0; x < w; x++ {
	// 			board.DrawAll(MakeColor(0, 0, 0))
	// 			board.DrawPixel(x, y, MakeColor(20, 20, 20))
	// 			board.Save()
	// 			time.Sleep(50 * time.Millisecond)
	// 		}
	// 	}
	// err = board.DrawRectOutline(9, 9, 15, 15, MakeColor(20, 0, 0))
	// board.DrawLine(0, 0, 10, 19, MakeColor(20, 20, 20))
	// }

	poll := make(chan string)
	go board.pollSensors(poll)
	for {
		select {
		case msg := <-poll:
			_ = msg
			board.processSensors()
			go board.pollSensors(poll)
			fmt.Println("foo")
		default:
		}
		_ = board.printBoardState()

		// process board
		ns := time.Since(start)
		process_dance(board, ns)
		board.Save() // draw the board
	}
}
