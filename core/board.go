/*
A note on indexing of the board

0,0 is the left side if you stand on the high side of the board.

For pixels:  along the high side is X, along the left is Y
For squares: lid 0,0 is the left side. 0,3 is the right side.
The board sensors map to the same.

*/

package core
import "sort"
import "fmt"
import "errors"
import "math"
// import "time"

type Board struct {
	strand           *Strand
	sensors          *Sensors
	pixelW           int
	pixelH           int
	squareW          int
	squareH          int
	includeVerticals bool
	poll             chan string
	boardItems		 []Drawer
	shiftx			 int
	shifty			 int
	scalePix		 float32
}


/////////////////////////////////
// CONNECTION FUNCTIONS
/////////////////////////////////

func MakeBoard() (*Board, error) {
	brd := Board{}
	brd.pixelW = 35
	brd.pixelH = 43
	brd.squareW = 5
	brd.squareH = 6
	brd.shiftx  = 0
	brd.shifty  = 0
	brd.strand = &Strand{}
	brd.sensors = &Sensors{}
	brd.sensors.initSensors(brd.squareW, brd.squareH)
	brd.includeVerticals = false
	brd.poll = make(chan string)
	brd.boardItems = make([]Drawer,0,50)
	go brd.pollSensors(brd.poll)
	err := brd.strand.Connect(mapLedColor(brd.pixelW*brd.pixelH) + 35*3 + 42*2 + 100)
	if err != nil {
		return nil, err
	} else {
		return &brd, nil
	}
}

func (brd *Board) DebugSensors(state bool) {
	brd.sensors.DebugSensors(state)
}

func (brd *Board) Free() {
	brd.strand.Free()
	brd.sensors.stopSensors()
}

func (brd *Board) Save() {
	brd.drawItems()
	brd.strand.Save()
}

func (brd *Board) drawItems(){
	for _, item := range brd.boardItems {
		if item != nil {
			if item.remove() {
				item = nil
			} else {
				item.draw(brd)
			}
		}
	}
}

func (brd *Board) ClearItems(){
	brd.boardItems = make([]Drawer,0,50)
}

func (brd *Board) AddItem(d Drawer){
	brd.boardItems = append(brd.boardItems, d)
	sort.Sort(DrawerByZ(brd.boardItems))
	fmt.Println(len(brd.boardItems))
}

/////////////////////////////////
// DRAWING FUNCTIONS
/////////////////////////////////

func (brd *Board) SetVerticalMode(includeVerticals bool) {
	brd.includeVerticals = includeVerticals
}

func (brd *Board) DrawPixel(x, y int, c Color) {
	brd.DrawPixel3(x,y,0,c)
	return
	// if x < 0 || x >= brd.pixelW || y < 0 || y >= brd.pixelH {
	// 	// fmt.Println("Pixel was drawn outside the board's space, at", x, y)
	// 	return
	// }
	// pixelNum := getPixelNum(x, y, brd.squareW, brd.squareH, brd.includeVerticals)
	// brd.setColor(pixelNum, c)
}

// alternative mapping.
// z = 0 is top, with the intermediate vertical if includeverticals=true
// z = 1 is top sides.  x and y is the same as the edge pixels on the top
// z = 2 is bottom sides
// if includeVerticals is true, then the front intermedate must be set with z=0

func (brd *Board) Shift(x,y int) {
	brd.shiftx = x
	brd.shifty = y
}

func (brd *Board) Unshift(){
	brd.shiftx = 0
	brd.shifty = 0
}

func (brd *Board) DrawPixel3(x, y, z int, c Color) {
//	for i:=0; i< 49*30; i++ { brd.setColor(i, Yellow) }
	x += brd.shiftx
	y += brd.shifty
	if x<0 || x>34 || y<0 || y>42 || (y==42 && !brd.includeVerticals) { return }
	// if brd.scalePix!=1.0 {
	// 	c = c.Scale(brd.scalePix)
	// }

	base := 6*5*49
	side := 7
	switch z {
	case 0:  // top surface
		if brd.includeVerticals {
			if y==14 {
				brd.setColor(base+18*side+x, c)
				return
			}
			if y>14 { y-- }
		}
		sCol := x/7
		sRow := y/7
		squares := sRow*5;
		if sRow%2==0 {
			squares += sCol
		} else {
			squares += 4-sCol
		}
		sy := y%7
		sx := x%7
		if sy%2==1{
			brd.setColor(49*squares+sy*7+sx,c)
		} else {
			brd.setColor(49*squares+(sy+1)*7-sx-1,c)
		}
	case 1: // middle level
		switch {
		case x>=0 && x <=34 && y==0:
			brd.setColor(base+30*side-x-1, c)

		case x==0 && y>0 && y<13:
			brd.setColor(base+30*side+y, c)

		case !brd.includeVerticals && x>=0 && x <=34 && y==13:
			brd.setColor(base+18*side+x, c)

		case x==34 && y>0 && y<13:
			brd.setColor(base+25*side-y-1, c)

		}
	case 2: // bottom level
		switch {
		//left side, bottom level
		case x==0 && y>0 && y<13:
			brd.setColor(base+16*side+y, c)
		case !brd.includeVerticals && x==0 && y>13 && y<42:
			brd.setColor(base+30*side+y, c)
		case brd.includeVerticals && x==0 && y>14 && y<43:
			brd.setColor(base+30*side+y-1, c)
		
		case y==0 && x>0 && x<34:
		 	brd.setColor(base+16*side-x-1, c)
		
		case x==34 && y>0 && y<13:
			brd.setColor(base+11*side-y-1, c)
		case !brd.includeVerticals && x==34 && y>13 && y<42:
			brd.setColor(base+11*side-y-1, c)
		case brd.includeVerticals && x==34 && y>14 && y<43:
			brd.setColor(base+11*side-y, c)

		case !brd.includeVerticals && y==41 && x>0 && x<34:
			brd.setColor(base + x, c)
			return
		case brd.includeVerticals && y==42 && x>0 && x<34:
			brd.setColor(base + x, c)
			return
		}
	}
}



func (brd *Board) DrawSidePixel(col, level int, c Color) {
	pixelNum := getSidePixelNum(level, col)
	if pixelNum != 0 {
		brd.setColor(pixelNum, c)
	}
}

func (brd *Board) DrawSquare(col int, row int, c Color) error {
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			brd.DrawPixel(col*7+i, row*7+j, c)
		}
	}
	return nil
}

func (brd *Board) DrawAll(c Color) error {
	brd.DrawAllSides(c)
	for x := 0; x < brd.pixelW; x++ {
		for y := 0; y < brd.pixelH; y++ {
			brd.DrawPixel(x, y, c)
		}
	}
	return nil
}

// writes out a string
// orientation should be orient_0, orient_90, orient_180, or orient_270
//
func (brd *Board) WriteText(text string, x, y, orientation int, c Color) {
	for index, runeValue := range text {
		_ = index
		x, y = brd.DrawRune(runeValue, x, y, orientation, c)
	}
}

// draws one character and returns the next position to draw.
func (brd *Board) DrawRune(ch rune, x, y, orientation int, c Color) (x1 int, y1 int) {
	for i := 0; i < len(letters); i++ {
		if letters[i].r == ch {
			r := letters[i]
			switch {
			case orientation == Orient_0:
				shift := 5 - r.width // adjust for letters less than the full 5 pixels - most are 4 some 3 pixels wide
				for i = 0; i < 7; i++ {
					if (r.data[i] & 16) >= 1 {
						brd.DrawPixel(x+0-shift, y-i, c)
					}
					if (r.data[i] & 8) >= 1 {
						brd.DrawPixel(x+1-shift, y-i, c)
					}
					if (r.data[i] & 4) >= 1 {
						brd.DrawPixel(x+2-shift, y-i, c)
					}
					if (r.data[i] & 2) >= 1 {
						brd.DrawPixel(x+3-shift, y-i, c)
					}
					if (r.data[i] & 1) >= 1 {
						brd.DrawPixel(x+4-shift, y-i, c)
					}
				}
				return x + r.width + 1, y // return the next position
			case orientation == Orient_90:
				y += 5 - r.width // adjust for letters less than the full 5 pixels - most are 4 some 3 pixels wide
				for i = 0; i < 7; i++ {
					if (r.data[i] & 16) >= 1 {
						brd.DrawPixel(x-i, y-0, c)
					}
					if (r.data[i] & 8) >= 1 {
						brd.DrawPixel(x-i, y-1, c)
					}
					if (r.data[i] & 4) >= 1 {
						brd.DrawPixel(x-i, y-2, c)
					}
					if (r.data[i] & 2) >= 1 {
						brd.DrawPixel(x-i, y-3, c)
					}
					if (r.data[i] & 1) >= 1 {
						brd.DrawPixel(x-i, y-4, c)
					}
				}
				return x, y - r.width - 2 // return the next position
			case orientation == Orient_180:
				x += 5 - r.width // adjust for letters less than the full 5 pixels - most are 4 some 3 pixels wide
				for i = 0; i < 7; i++ {
					if (r.data[i] & 16) >= 1 {
						brd.DrawPixel(x-0, y+i, c)
					}
					if (r.data[i] & 8) >= 1 {
						brd.DrawPixel(x-1, y+i, c)
					}
					if (r.data[i] & 4) >= 1 {
						brd.DrawPixel(x-2, y+i, c)
					}
					if (r.data[i] & 2) >= 1 {
						brd.DrawPixel(x-3, y+i, c)
					}
					if (r.data[i] & 1) >= 1 {
						brd.DrawPixel(x-4, y+i, c)
					}
				}
				return x - r.width - 2, y // return the next position
			case orientation == Orient_270:
				y -= 5 - r.width // adjust for letters less than the full 5 pixels - most are 4 some 3 pixels wide
				for i = 0; i < 7; i++ {
					if (r.data[i] & 16) >= 1 {
						brd.DrawPixel(x+i, y+0, c)
					}
					if (r.data[i] & 8) >= 1 {
						brd.DrawPixel(x+i, y+1, c)
					}
					if (r.data[i] & 4) >= 1 {
						brd.DrawPixel(x+i, y+2, c)
					}
					if (r.data[i] & 2) >= 1 {
						brd.DrawPixel(x+i, y+3, c)
					}
					if (r.data[i] & 1) >= 1 {
						brd.DrawPixel(x+i, y+4, c)
					}
				}
				return x, y + r.width + 2 // return the next position
			}
		}
	}
	return x, y // letter not found
}

func ipart(x float64) float64 {
	return math.Floor(x)
}

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}

func fpart(x float64) float64 {
	return x - math.Floor(x)
}

func rfpart(x float64) float64 {
	return 1 - fpart(x)
}

func (brd *Board) DrawLine(x0i, y0i, x1i, y1i int, c Color) {
	x0, y0, x1, y1 := float64(x0i), float64(y0i), float64(x1i), float64(y1i)
	steep := (math.Abs(y1-y0) > math.Abs(x1-x0))

	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := y1 - y0
	gradient := dy / dx

	xend := round(x0)
	yend := y0 + gradient*(xend-x0)
	xgap := rfpart(x0 + 0.5)
	xpxl1 := xend //this will be used in the main loop
	ypxl1 := ipart(yend)
	if steep {
		brd.DrawPixel(int(ypxl1), int(xpxl1), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(ypxl1+1), int(xpxl1), c.WithAlpha(fpart(yend)*xgap))
	} else {
		brd.DrawPixel(int(xpxl1), int(ypxl1), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(xpxl1), int(ypxl1+1), c.WithAlpha(fpart(yend)*xgap))
	}
	intery := yend + gradient // first y-intersection for the main loop

	// handle second endpoint

	xend = round(x1)
	yend = y1 + gradient*(xend-x1)
	xgap = fpart(x1 + 0.5)
	xpxl2 := xend //this will be used in the main loop
	ypxl2 := ipart(yend)
	if steep {
		brd.DrawPixel(int(ypxl2), int(xpxl2), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(ypxl2+1), int(xpxl2), c.WithAlpha(fpart(yend)*xgap))
	} else {
		brd.DrawPixel(int(xpxl2), int(ypxl2), c.WithAlpha(rfpart(yend)*xgap))
		brd.DrawPixel(int(xpxl2), int(ypxl2+1), c.WithAlpha(fpart(yend)*xgap))
	}

	// main loop

	for x := xpxl1 + 1; x <= xpxl2-1; x++ {
		if steep {
			brd.DrawPixel(int(ipart(intery)), int(x), c.WithAlpha(rfpart(intery)))
			brd.DrawPixel(int(ipart(intery)+1), int(x), c.WithAlpha(fpart(intery)))
		} else {
			brd.DrawPixel(int(x), int(ipart(intery)), c.WithAlpha(rfpart(intery)))
			brd.DrawPixel(int(x), int(ipart(intery)+1), c.WithAlpha(fpart(intery)))
		}
		intery = intery + gradient
	}
}

func (brd *Board) DrawRect(x1, y1, x2, y2 int, c Color) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			brd.DrawPixel(x, y, c)
		}
	}
}

func (brd *Board) FillSquare(x, y int, c Color) {
	if brd.includeVerticals && y>1 { 
		brd.DrawRect(x*7, y*7+1, x*7+6, y*7+6+1, c)
	} else {
		brd.DrawRect(x*7, y*7, x*7+6, y*7+6, c)
	} 
}

func (brd *Board) DrawRectOutline(x1, y1, x2, y2 int, c Color) {
	for x := x1; x <= x2; x++ {
		brd.DrawPixel(x, y1, c)
		brd.DrawPixel(x, y2, c)
	}

	for y := y1 + 1; y <= y2-1; y++ {
		brd.DrawPixel(x1, y, c)
		brd.DrawPixel(x2, y, c)
	}
}

func (brd *Board) DrawCircle(x1, y1, r float64, c Color) error {
	return errors.New("Not implemented")
}

func (brd *Board) DrawCircleOutline(x1, y1, r float64, c Color) error {
	return errors.New("Not implemented")
}

func (brd *Board) DrawSmallCircle(x, y float64, c Color) {
	drawx := int(x + 0.5)
	drawy := int(y + 0.5)
	extrax := x - float64(drawx)
	extray := y - float64(drawy)
	brd.DrawPixel(drawx, drawy, c.WithAlpha(1-pointDistance(0, 0, extrax, extray)))
	brd.DrawPixel(drawx+1, drawy, c.WithAlpha(1-pointDistance(1, 0, extrax, extray)))
	brd.DrawPixel(drawx, drawy+1, c.WithAlpha(1-pointDistance(0, 1, extrax, extray)))
	brd.DrawPixel(drawx+1, drawy+1, c.WithAlpha(1-pointDistance(1, 1, extrax, extray)))
	brd.DrawPixel(drawx-1, drawy, c.WithAlpha(1-pointDistance(-1, 0, extrax, extray)))
	brd.DrawPixel(drawx, drawy-1, c.WithAlpha(1-pointDistance(0, -1, extrax, extray)))
	brd.DrawPixel(drawx-1, drawy-1, c.WithAlpha(1-pointDistance(-1, -1, extrax, extray)))
	brd.DrawPixel(drawx-1, drawy+1, c.WithAlpha(1-pointDistance(-1, 1, extrax, extray)))
	brd.DrawPixel(drawx+1, drawy-1, c.WithAlpha(1-pointDistance(1, -1, extrax, extray)))
}

func (brd *Board) DrawSmallArrow(x, y int, c Color, d Direction) {
	drawx, drawy := brd.GetSquare(x, y)
	brd.DrawPixel(drawx+3, drawy+3, c)
	switch d {
	case Up:
		brd.DrawPixel(drawx+3, drawy+4, c)
		brd.DrawPixel(drawx+4, drawy+4, c)
		brd.DrawPixel(drawx+2, drawy+4, c)
	case Down:
		brd.DrawPixel(drawx+3, drawy+2, c)
		brd.DrawPixel(drawx+2, drawy+2, c)
		brd.DrawPixel(drawx+4, drawy+2, c)
	case Left:
		brd.DrawPixel(drawx+4, drawy+3, c)
		brd.DrawPixel(drawx+4, drawy+4, c)
		brd.DrawPixel(drawx+4, drawy+2, c)
	case Right:
		brd.DrawPixel(drawx+2, drawy+3, c)
		brd.DrawPixel(drawx+2, drawy+4, c)
		brd.DrawPixel(drawx+2, drawy+2, c)
	}
}

func pointDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func (board *Board) RefreshSensors() {
	fmt.Println("REFRESH SENSORS")
	select {
	case msg := <-board.poll:
		_ = msg
		board.processSensors()
		go board.pollSensors(board.poll)
	default:
	}
}

func (brd *Board) CheckPressed(col int, row int) bool {
	state := brd.getSensorState(col, row)
	return state == 3
}

func (brd *Board) CheckDown(col, row int) bool {
	state := brd.getSensorState(col, row)
	return state == 2 || state == 3
}

func (brd *Board) CheckAnyDown() bool {
	for x:=0; x<5; x++ {
		for y:=0; y<6; y++ {
			if brd.CheckDown(x,y) { return true }
		}
	}
	return false
}

func (brd *Board) CheckUp(col int, row int) bool {
	state := brd.getSensorState(col, row)
	return state == 0 || state == 1
}

func (brd *Board) CheckReleased(col int, row int) bool {
	state := brd.getSensorState(col, row)
	return state == 1
}

func (brd *Board) GetSquare(col int, row int) (int, int) {
	x := col * 7
	y := row * 7
	if brd.includeVerticals && row >= 2 {
		y++
	}
	return x, y
}

func (brd *Board) DrawAllSides(c Color) {
	for x:=0; x<35; x++ {
		brd.DrawPixel3(x,0,2,c)
		brd.DrawPixel3(x,0,1,c)
		if !brd.includeVerticals {
			brd.DrawPixel3(x,13,1,c)
			brd.DrawPixel3(x,41,2,c)
		} else {
			brd.DrawPixel3(x,42,2,c)
		}
	}
	for y:=0; y<43; y++{
		brd.DrawPixel3(0, y,2,c)
		brd.DrawPixel3(34,y,2,c)
		if y<14 {
			brd.DrawPixel3(0,y,1,c)
			brd.DrawPixel3(34,y,1,c)
		}
	}
}

/////////////////////////////////
// INTERNAL FUNCTIONS
/////////////////////////////////

func (brd *Board) getSensorState(col, row int) int {
	return brd.sensors.getBoardState(col, row) // hack to accomodate the shift
}

func getPixelNum(x, y, sqW, sqH int, includeVerticals bool) int {
	if includeVerticals {
		if y >= 15 {
			y--
		} else if y == 14 {
			return 7*7*sqW*sqH + 35 + 42 + 35 + 14 + x
		}
	}
	col := x / 7
	row := y / 7
	xPixelInSq := x % 7
	yPixelInSq := y % 7

	var boardNum, pixelNum int

	// NOTE: this is hardcoded for 49px/square
	if row%2 == 0 {
		boardNum = row*sqW + col
	} else {
		boardNum = row*sqW + (sqW - 1) - col
	}

	if yPixelInSq%2 == 1 {
		pixelNum = yPixelInSq*7 + xPixelInSq
	} else {
		pixelNum = yPixelInSq*7 + 6 - xPixelInSq
	}

	return boardNum*49 + pixelNum
}

func getSidePixelNum(level, col int) int {
	col += 1
	col += (col - 1) / 5 * 2
	base := 7 * 7 * 5 * 6
	if col < 7*9 && level == 0 {
		return base + col
	}
	if col < 7*18 && level == 0 {
		return base + col
	}
	if col < 7*18 && level == 1 && col > 7*9 {
		return base + col + 7*5 + 7*9
	}
	if col < 7*36 && level == 0 {
		return base + col + 7*5 + 7*9
	}
	return 0 // to non-existant pixel
}

func (brd *Board) setColor(led int, color Color) {
	brd.strand.SetColor(mapLedColor(led), color)
}

func (brd *Board) printBoardState() error {
	for r := 0; r < brd.squareH; r++ {
		for c := brd.squareW - 1; c >= 0; c-- {
			state := brd.sensors.getBoardState(c, r)
			switch {
			case state == up:
				fmt.Printf("-")
			case state == down:
				fmt.Printf("X")
			case state == pressed:
				fmt.Printf("|")
			case state == released:
				fmt.Printf("+")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
	return nil
}

func (brd *Board) pollSensors(poll chan string) {
	fmt.Printf("Enter pollSensors\n")
	brd.sensors.readSensors()
	fmt.Printf("Exit pollSensors\n")
	poll <- "ready"
}

func (brd *Board) processSensors() {
	brd.sensors.processSensors()
}


// removes the extra light at the end.
func mapLedColor(i int) int {
	// board sqW, sqH:
	if i < 50*5*6 {
		return i + i/49
	} else {
		return i + (5*6)*50/49
	}
}
