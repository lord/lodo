package maze

import "github.com/lord/lodo/core"

type Wash struct {
	goingUp  bool
	goingIn  bool
	borderY  int
	callback func()
}

func MakeWash(goingUp, goingIn bool, callback func()) *Wash {
	var borderY int
	if goingUp {
		borderY = 7*6 - 1
	} else {
		borderY = 0
	}
	return &Wash{
		goingUp:  goingUp,
		goingIn:  goingIn,
		borderY:  borderY,
		callback: callback,
	}
}

var washColor = core.MakeColor(0, 10, 0)
var washSpeed = 3

func (wash *Wash) Step(game *Game) {
	if wash.goingUp {
		wash.borderY -= washSpeed
	} else {
		wash.borderY += washSpeed
	}
	if wash.borderY < 0 || wash.borderY >= 7*6 {
		game.DeleteObject(wash)
		if wash.callback != nil {
			wash.callback()
		}
	}
}

func (wash *Wash) Draw(board *core.Board) {
	if wash.goingIn == wash.goingUp {
		board.DrawRect(0, wash.borderY, 7*5, 6*7, washColor)
	} else {
		board.DrawRect(0, 0, 7*5, wash.borderY, washColor)
	}
}
