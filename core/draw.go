package core

// import _ = "errors"
// import _ = "math"
import "time"
import "fmt"
import "sort"

//
// Graphics Section
//
type Drawer interface {
	draw(*Board) 
	zOrder() int
	remove() bool
}

type DrawerByZ []Drawer 

func (a DrawerByZ) Len() int           { return len(a) }
func (a DrawerByZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DrawerByZ) Less(i, j int) bool { return a[i].zOrder() < a[j].zOrder() }

// Rect is a one pixel wide item with the upper corner at x y
type   Rect struct {
	x,y,width,depth,z int
	c Color
	start time.Time
	periodMS int 
	canDelete bool
}


// Create a new rectangle.  Period is blink rate.  0 is no blink.
func NewDrawRect(x,y,w,h,z int, c Color, periodMS int) Rect {
	return Rect{x,y,w,h,z,c, time.Now(), periodMS, false}
}

func (r Rect) draw(brd *Board) {
	if r.periodMS == 0 {
		brd.DrawRectOutline(r.x, r.y, r.x+r.width-1, r.y+r.depth-1, r.c)
	} else {
		ms := int(time.Since(r.start)/time.Millisecond) % r.periodMS
		scale := 1.0 - float32(ms) * 2 / float32(r.periodMS)
		if scale < 0 { scale = -scale }
		brd.DrawRectOutline(r.x, r.y, r.x+r.width-1, r.y+r.depth-1, r.c.Scale(scale))
	}
}

func (r Rect) zOrder() int {
	return r.z
}

func (r Rect) remove() bool {
	return r.canDelete
}

type RectFlash struct {
	x,y,width,depth,z int
	c Color
	start time.Time
	flashDurationMS int 
	canDelete bool
}

// Create a new filled rectangle that starts strong, and decays.
func NewDrawRectFlash(x,y,w,h,z int, c Color, flashDurationMS int) RectFlash {
	return RectFlash{x,y,w,h,z,c, time.Now(), flashDurationMS, false}
}

func (r RectFlash) draw(brd *Board) {
	msSince := int(time.Since(r.start).Nanoseconds()/1000/1000)
	switch {
	case msSince < r.flashDurationMS/3:
		brd.DrawRect(r.x, r.y, r.x+r.width-1, r.y+r.depth-1, r.c)
	case msSince < r.flashDurationMS:
		scale := 1.0 - (float32(msSince)-float32(r.flashDurationMS)/3) / (2.0/3.0*float32(r.flashDurationMS))
		fmt.Println(scale)
		brd.DrawRect(r.x, r.y, r.x+r.width-1, r.y+r.depth-1, r.c.Scale(scale))
	default:
		r.canDelete = true
	}
}

func (r RectFlash) zOrder() int {
	return r.z
}

func (r RectFlash) remove() bool {
	//if Time.Now()
	return r.canDelete
}


/////////////////////
// sText
/////////////////////
type Stext struct {
	x,y,z int
	c Color
	orient int
	text string
	canDelete bool
}

func NewStext(x,y,z int, c Color, orient int, text string) Stext {
	return Stext{x,y,z,c, orient, text, false}
}


func (t Stext) draw(brd *Board) {
	brd.WriteText(t.text,t.x,t.y,t.orient, t.c)
}

func (t Stext) zOrder() int {
	return t.z
}

func (t Stext) remove() bool {
	//if Time.Now()
	return t.canDelete
}

/////////////////////
// Pallette
///////////////////
type coords struct {
	x, y int
} 

type Pallette struct {
	coord *coords
	z int
	canDelete bool
	palletteItems []Drawer
}

func NewPallette(x, y, z int) Pallette {
	return Pallette{&coords{x,y},z,false,make([]Drawer,0,0)}
}

func (p Pallette) draw(brd *Board) {
	brd.Shift(p.coord.x, p.coord.y)
	defer brd.Unshift()
	p.drawItems(brd)
}

func (p Pallette) zOrder() int {
	return p.z
}

func (p *Pallette) drawItems(brd *Board){
	for _, item := range p.palletteItems {
		if item != nil {
			if item.remove() {
				item = nil
			} else {
				item.draw(brd)
			}
		}
	}
}

func (p *Pallette) ClearItems(){
	p.palletteItems = make([]Drawer,0,50)
}

func (p *Pallette) AddItem(d Drawer){
	fmt.Print("Before ")
	fmt.Println(len(p.palletteItems))
	p.palletteItems = append(p.palletteItems, d)
	fmt.Print("After ")
	fmt.Println(len(p.palletteItems))
	sort.Sort(DrawerByZ(p.palletteItems))
	fmt.Println(len(p.palletteItems))
}

func (p Pallette) remove() bool {
	//if Time.Now()
	return p.canDelete
}

func (p Pallette) Shift(x,y int) {
	p.coord.x = x
	p.coord.y = y
}

/////////////////////
// Flashing Arrow
///////////////////

type SArrow struct {
	x,y,orient,z int
	c Color
	start time.Time
	periodMS int 
	canDelete bool
}

func NewSArrow(x,y,orient,z int, c Color, flashDurationMS int) SArrow {
	return SArrow{x,y,orient,z,c, time.Now(), flashDurationMS, false}
}

func (a SArrow) draw(brd *Board) {
	scale := float32(1.0)
	if a.periodMS != 0 {
		ms := int(time.Since(a.start)/time.Millisecond) % a.periodMS
		scale = 1.0 - float32(ms) * 2 / float32(a.periodMS)
		if scale < 0 { scale = -scale }
	}
	c := a.c.Scale(scale)
	switch a.orient {
		case 0: 
		brd.DrawRect(a.x-3,a.y,a.x-3,a.y,c)
		brd.DrawRect(a.x-2,a.y-1,a.x-2,a.y,c)
		brd.DrawRect(a.x-1,a.y-2,a.x+1,a.y+3,c)
		brd.DrawRect(a.x,a.y-3,a.x,a.y-3,c)
		brd.DrawRect(a.x+2,a.y-1,a.x+2,a.y,c)
		brd.DrawRect(a.x+3,a.y,a.x+3,a.y,c)
		return
		case 90: 
		brd.DrawRect(a.x  ,a.y-3,a.x,  a.y-3,c)
		brd.DrawRect(a.x  ,a.y-2,a.x+1,a.y-2,c)
		brd.DrawRect(a.x-3,a.y-1,a.x+2,a.y+1,c)
		brd.DrawRect(a.x+3,a.y,  a.x+3,a.y,  c)
		brd.DrawRect(a.x,  a.y+2,a.x+1,a.y+2,c)
		brd.DrawRect(a.x  ,a.y+3,a.x  ,a.y+3,  c)
		return
		case 180: 
		brd.DrawRect(a.x-3,a.y,a.x-3,a.y,c)
		brd.DrawRect(a.x-2,a.y,a.x-2,a.y+1,c)
		brd.DrawRect(a.x-1,a.y-3,a.x+1,a.y+2,c)
		brd.DrawRect(a.x,a.y+3,a.x,a.y+3,c)
		brd.DrawRect(a.x+2,a.y,a.x+2,a.y+1,c)
		brd.DrawRect(a.x+3,a.y,a.x+3,a.y,c)

		return
		case 270: 
		brd.DrawRect(a.x  ,a.y-3,a.x,  a.y-3,c)
		brd.DrawRect(a.x-1,a.y-2,a.x  ,a.y-2,c)
		brd.DrawRect(a.x-2,a.y-1,a.x+3,a.y+1,c)
		brd.DrawRect(a.x-3,a.y,  a.x-3,a.y,  c)
		brd.DrawRect(a.x-1,a.y+2,a.x  ,a.y+2,c)
		brd.DrawRect(a.x  ,a.y+3,a.x  ,a.y+3,  c)
		return
	}
}



func (a SArrow) zOrder() int {
	return a.z
}

func (a SArrow) remove() bool {
	//if Time.Now()
	return a.canDelete
}