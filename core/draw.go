package core

// import _ = "errors"
// import _ = "math"
import "time"
import "fmt"

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