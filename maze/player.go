package maze

import "github.com/lord/lodo/core"

type Player struct {
	x, y         int
	drawX, drawY int
}

var playerColor = core.MakeColor(0, 10, 0)

func MakePlayer(x, y int) *Player {
	return &Player{
		x:     x,
		y:     y,
		drawX: x * 7,
		drawY: y * 7,
	}
}

func (player *Player) Step(game *Game) {
	if game.CheckPressed(player.x+1, player.y) && game.CheckMove(player.x, player.y, Right) {
		player.x += 1
	} else if game.CheckPressed(player.x, player.y+1) && game.CheckMove(player.x, player.y, Down) {
		player.y += 1
	} else if game.CheckPressed(player.x-1, player.y) && game.CheckMove(player.x, player.y, Left) {
		player.x -= 1
	} else if game.CheckPressed(player.x, player.y-1) && game.CheckMove(player.x, player.y, Up) {
		player.y -= 1
	}

	targetX, targetY := game.board.GetSquare(player.x, player.y)
	if targetX > player.drawX {
		player.drawX++
	}
	if targetX < player.drawX {
		player.drawX--
	}
	if targetY > player.drawY {
		player.drawY++
	}
	if targetY < player.drawY {
		player.drawY--
	}
}

func (player *Player) Draw(board *core.Board) {
	board.DrawPixel(player.drawX+3, player.drawY+3, playerColor)
	board.DrawPixel(player.drawX+4, player.drawY+3, playerColor)
	board.DrawPixel(player.drawX+3, player.drawY+4, playerColor)
	board.DrawPixel(player.drawX+3, player.drawY+2, playerColor)
	board.DrawPixel(player.drawX+2, player.drawY+3, playerColor)
}
