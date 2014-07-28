package main

import "fmt"

// import "time"

//import "math"

const rows int = 5
const cols int = 4

func main() {
	// start := time.Now() // starting time in ns
	board := Board{}

	w := 20
	h := 25
	err := board.Connect(w, h, cols, rows)
	defer board.Free()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {
		board.DrawAll(MakeColor(0, 0, 0))
		err = board.DrawRectOutline(9, 9, 15, 15, MakeColor(255, 0, 0))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		board.DrawLine(10, 0, 10, 19, MakeColor(200, 200, 200))
		board.Save()
	}

	// poll := make(chan string)
	// // go board.pollSensors(poll)
	// for {
	// 	select {
	// 	case msg := <-poll:
	// 		_ = msg
	// 		board.processSensors()
	// 		go board.pollSensors(poll)
	// 	default:
	// 	}
	// 	//		_ = board.printBoardState()

	// 	// process board
	// 	ns := time.Since(start)
	// 	process_dance(board, ns)

	// 	board.Save() // draw the board
	// }
}
