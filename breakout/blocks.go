package breakout

import "github.com/lord/lodo/core"


type block struct {
	// x: coord r/l
	// y: coor back / forth
	// x/y at edge of paddle, top of square
	// w: width of paddle
	x, y            float64
	color           core.Color
	btype			int
	show			bool
}
const blockCount = 45
var blocks []block

func initBlocks() {
	blocks = make([]block, blockCount)
	for br := 0; br<9; br++ { //block row
		for c := 0; c<5; c++ {
			blocks[br*5+c].show = true
			switch {
			case br==0:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 0.0
				blocks[br*5+c].color=core.Green
			case br==1:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 2.0
				blocks[br*5+c].color = core.Blue
			case br==2:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 4.0
				blocks[br*5+c].color=core.Red
			case br==3:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 7.0
				blocks[br*5+c].color=core.Gray
			case br==4:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 9.0
				blocks[br*5+c].color=core.Gray
			case br==5:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 11.0
				blocks[br*5+c].color=core.Blue10
			case br==6:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 16.0
				blocks[br*5+c].color=core.Purple
			case br==7:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 18.0
				blocks[br*5+c].color=core.Yellow
			case br==8:
				blocks[br*5+c].x = float64(c)*7.0
				blocks[br*5+c].y = 20.0
				blocks[br*5+c].color=core.Green
			}
		}
	}
}

func (blk *block) Draw (brd *core.Board) {
	if (blk.show) {
		x := int(blk.x + 0.5)
		y := int(blk.y + 0.5)
		brd.DrawRect(x, y, x+6, y+1, blk.color)
	}
}

func (blk *block) hit(b *ball) bool {
	if blk.show && (b.y >= blk.y && b.y <= blk.y + 1) && (b.x >= blk.x && b.x <= blk.x + 7) {
		s := core.MakeSound(core.Bounce1)
		b.hits++
		s.Play()
		return true
	}	
	return false
}
