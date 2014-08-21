package maze

import "github.com/lord/lodo/core"

type Player struct {
	x, y         int
	drawX, drawY int
}

var playerColor = core.MakeColor(0, 10, 0)
var playerArrowColor = core.MakeColor(0, 2, 0)
var playerGateArrowColor = core.MakeColor(2, 2, 0)

func MakePlayer(x, y int) *Player {
	return &Player{
		x:     x,
		y:     y,
		drawX: x * 7,
		drawY: y * 7,
	}
}

const speed int = 2

var keyCounterColor = core.MakeColor(4, 4, 0)

func (player *Player) Step(game *Game) {
	if game.CheckPressed(player.x+1, player.y) && game.CheckMove(player.x, player.y, Right) {
		player.x += 1
	} else if game.CheckPressed(player.x, player.y+1) && game.CheckMove(player.x, player.y, Down) {
		player.y += 1
	} else if game.CheckPressed(player.x-1, player.y) && game.CheckMove(player.x, player.y, Left) {
		player.x -= 1
	} else if game.CheckPressed(player.x, player.y-1) && game.CheckMove(player.x, player.y, Up) {
		player.y -= 1
	} else if game.CheckPressed(player.x, player.y) {
		// do bomb stuff
	} else if gate := game.FindGate(player.x, player.y, Right); game.keys > 0 && game.CheckPressed(player.x+1, player.y) && gate != nil {
		game.DeleteObject(gate)
		game.keys--
		player.x += 1
	} else if gate := game.FindGate(player.x, player.y, Down); game.keys > 0 && game.CheckPressed(player.x, player.y+1) && gate != nil {
		game.DeleteObject(gate)
		game.keys--
		player.y += 1
	} else if gate := game.FindGate(player.x, player.y, Left); game.keys > 0 && game.CheckPressed(player.x-1, player.y) && gate != nil {
		game.DeleteObject(gate)
		game.keys--
		player.x -= 1
	} else if gate := game.FindGate(player.x, player.y, Up); game.keys > 0 && game.CheckPressed(player.x, player.y-1) && gate != nil {
		game.DeleteObject(gate)
		game.keys--
		player.y -= 1
	} else if game.CheckPressed(player.x, player.y) {
		// do bomb stuff
	} else {
		for x := 0; x <= 4; x++ {
			for y := 0; y <= 5; y++ {
				if game.CheckPressed(x, y) {
					game.objects = append(game.objects, MakeBurst(core.MakeColor(25, 0, 0), 20))
				}
			}
		}
	}

	if game.CheckMove(player.x, player.y, Up) {
		game.board.DrawSmallArrow(player.x, player.y-1, playerArrowColor, Up.ToCoreDirection())
	}
	if game.CheckMove(player.x, player.y, Down) {
		game.board.DrawSmallArrow(player.x, player.y+1, playerArrowColor, Down.ToCoreDirection())
	}
	if game.CheckMove(player.x, player.y, Left) {
		game.board.DrawSmallArrow(player.x-1, player.y, playerArrowColor, Left.ToCoreDirection())
	}
	if game.CheckMove(player.x, player.y, Right) {
		game.board.DrawSmallArrow(player.x+1, player.y, playerArrowColor, Right.ToCoreDirection())
	}

	if game.keys > 0 && game.FindGate(player.x, player.y, Up) != nil {
		game.board.DrawSmallArrow(player.x, player.y-1, playerGateArrowColor, Up.ToCoreDirection())
	}
	if game.keys > 0 && game.FindGate(player.x, player.y, Down) != nil {
		game.board.DrawSmallArrow(player.x, player.y+1, playerGateArrowColor, Down.ToCoreDirection())
	}
	if game.keys > 0 && game.FindGate(player.x, player.y, Left) != nil {
		game.board.DrawSmallArrow(player.x-1, player.y, playerGateArrowColor, Left.ToCoreDirection())
	}
	if game.keys > 0 && game.FindGate(player.x, player.y, Right) != nil {
		game.board.DrawSmallArrow(player.x+1, player.y, playerGateArrowColor, Right.ToCoreDirection())
	}

	// game.objects = append(game.objects, MakeFootprint(player.x, player.y, Left))

	targetX, targetY := game.board.GetSquare(player.x, player.y)
	if targetX > player.drawX {
		player.drawX += speed
		if targetX < player.drawX {
			player.drawX = targetX
		}
	}
	if targetX < player.drawX {
		player.drawX -= speed
		if targetX > player.drawX {
			player.drawX = targetX
		}
	}
	if targetY > player.drawY {
		player.drawY += speed
		if targetY < player.drawY {
			player.drawY = targetY
		}
	}
	if targetY < player.drawY {
		player.drawY -= speed
		if targetY > player.drawY {
			player.drawY = targetY
		}
	}
	for i := 0; i < game.keys; i++ {
		game.board.DrawPixel(player.drawX+1+i, player.drawY+1, keyCounterColor)
	}
}

func (player *Player) Draw(board *core.Board) {
	board.DrawPixel(player.drawX+3, player.drawY+3, playerColor)
	board.DrawPixel(player.drawX+4, player.drawY+3, playerColor)
	board.DrawPixel(player.drawX+3, player.drawY+4, playerColor)
	board.DrawPixel(player.drawX+3, player.drawY+2, playerColor)
	board.DrawPixel(player.drawX+2, player.drawY+3, playerColor)
}
