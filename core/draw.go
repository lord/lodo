package core

// import _ = "errors"
// import _ = "math"
import "time"
import "fmt"
import "sort"

// constants
const Anime_departright = 1
const Anime_departleft  = 2
const Anime_arriveright = 3
const Anime_arriveleft  = 4


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

//
// ShapeController
//
type ShapeController interface {
	canDelete() bool
	color() Color
}

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

type CRect struct {
	x,y,width,depth,z int
	sc ShapeController
}

func NewCRect(x,y,width,depth,z int, sc ShapeController) CRect {
	return CRect{x,y,width,depth,z,sc}
}

func (r CRect) zOrder() int {
	return r.z
}

func (r CRect) remove() bool {
	//if Time.Now()
	return r.sc.canDelete()
}

func (r CRect) draw(brd *Board) {
	c := r.sc.color()
	brd.DrawRect(r.x, r.y, r.x+r.width-1, r.y+r.depth-1, c)

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

// Rect is a one pixel wide item with the upper corner at x y
type   BlinkyRect struct {
	x,y,level,width,depth,z,rateMS int
	c []Color
	start time.Time
	canDelete bool
}


// Create a new rectangle.  Period is blink rate.  0 is no blink.
func NewBlinkyRect(x,y,level,w,h,z,rateMS int, c []Color) BlinkyRect {
	return BlinkyRect{x,y,level,w,h,z,rateMS,c, time.Now(), false}
}

func (r BlinkyRect) draw(b *Board) {
	x:=r.x
	y:=r.y
	pattern := int(time.Since(r.start)/time.Millisecond) % (4*r.rateMS) / r.rateMS
	dir := 1 // 1= right, 2 = up, 3 = left, 4 = down
	incX := 1; incY := 0
	for i:=0; i<2*(r.width+r.depth-2); i++ {
		b.DrawPixel3(x,y,r.level, r.c[(i+pattern)%len(r.c)])
		switch {
			case dir == 1 && x>=r.x+r.width-1:			dir = 2; incX= 0; incY=1
			case dir == 2 && y>=r.y+r.depth-1:			dir = 3; incX=-1; incY=0
			case dir == 3 && x<=r.x:					dir = 4; incX= 0; incY=-1
			case dir == 4 && y<=r.y:					return
		}
		x+= incX; y+= incY
	}
	// if r.periodMS == 0 {
	// 	brd.DrawRectOutline(r.x, r.y, r.x+r.width-1, r.y+r.depth-1, r.c)
	// } else {
	// 	ms := int(time.Since(r.start)/time.Millisecond) % r.periodMS
	// 	scale := 1.0 - float32(ms) * 2 / float32(r.periodMS)
	// 	if scale < 0 { scale = -scale }
	// 	brd.DrawRectOutline(r.x, r.y, r.x+r.width-1, r.y+r.depth-1, r.c.Scale(scale))
	// }
}

func (r BlinkyRect) zOrder() int {
	return r.z
}

func (r BlinkyRect) remove() bool {
	return r.canDelete
}




/////////////////////
// Pallette
///////////////////
type coords struct {
	x, y int
} 

type anime struct {
	animeStart time.Time
	animeMode int
	animeDurMS int
	visible bool
}

type Pallette struct {
	coord *coords
	z int
	canDelete bool
	palletteItems []Drawer
	a *anime
}

func NewPallette(x, y, z int) Pallette {
	return Pallette{&coords{x,y},z,false, make([]Drawer,0,0), &anime{time.Now(), 0,0,true}}
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
	// fmt.Printf("X: %v isVis: %v\n", p.coord.x, p.a.visible)
	if p.a.animeMode != 0 {
		ms := int(time.Since(p.a.animeStart).Nanoseconds()/1000/1000)
		p.a.visible = true 
		switch p.a.animeMode {
			case Anime_departright:
				if ms>p.a.animeDurMS { 
					p.a.visible = false 
					p.a.animeMode = 0
				} else {
					p.coord.x = int(35.0*float32(ms)/float32(p.a.animeDurMS))
				}
			case Anime_departleft:
				if ms>p.a.animeDurMS { 
					p.a.visible = false 
					p.a.animeMode = 0
				} else {
					p.coord.x = -int(35.0*float32(ms)/float32(p.a.animeDurMS))
				}
			case Anime_arriveright:
				if ms>p.a.animeDurMS { 
					p.a.animeMode = 0
					p.coord.x = 0
				} else {
					p.coord.x = 35-int(35.0*float32(ms)/float32(p.a.animeDurMS))
				}
			case Anime_arriveleft:
				if ms>p.a.animeDurMS { 
					p.a.animeMode = 0
					p.coord.x = 0
				} else {
					p.coord.x = -35+int(35.0*float32(ms)/float32(p.a.animeDurMS))
				}
			default:
		}
	}

	if !p.a.visible { return }
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

func (p *Pallette) Visible(vis bool) {
	p.a.visible = vis
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

func (p Pallette) BeginAnime(animeType, durMS int) {
	p.a.animeStart = time.Now()
	p.a.animeDurMS = durMS
	p.a.animeMode = animeType
	p.a.visible = true
//	oOn := time.Now().Add(time.Duration(delayMS)*time.Millisecond)
}

/////////////////////
// Flashing Arrow
///////////////////

type SArrow struct {
	x,y,orient,z int
	sc ShapeController	
}

func NewSArrow(x,y,orient,z int, sc ShapeController) SArrow {
	return SArrow{x,y,orient,z,sc}
}

func (a SArrow) draw(brd *Board) {
	c := a.sc.color()	
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
		brd.DrawRect(a.x  ,a.y+3,a.x  ,a.y+3,c)
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
		brd.DrawRect(a.x  ,a.y+3,a.x  ,a.y+3,c)
		return
	}
}

func (a SArrow) zOrder() int {
	return a.z
}

func (a SArrow) remove() bool {
	return a.sc.canDelete()
}




///////////////////////////////////
//
// Controllers

type SolidControl struct {
	c Color
}

func NewSolidControl(c Color) SolidControl {
	return SolidControl{c}
}

func (f SolidControl) canDelete() bool{
	return false
}

func (f SolidControl) color() Color{
	return f.c
}

type SolidShortControl struct {
	start time.Time
	durMS int
	c Color
}

func NewSolidShortControl(dur int, c Color) SolidShortControl {
	return SolidShortControl{time.Now(), dur, c}
}

func (f SolidShortControl) canDelete() bool{
	msSince := int(time.Since(f.start).Nanoseconds()/1000/1000)
	return msSince > f.durMS
}

func (f SolidShortControl) color() Color{
	return f.c
}

type FlashControl struct {
	start time.Time 
	fullDurationMS int
	decayDurationMS int
	c Color
}

func NewFlashControl(full, decay int, c Color) FlashControl {
	return FlashControl{ time.Now(), full, decay, c}
}

func (f FlashControl) canDelete() bool{
	msSince := int(time.Since(f.start).Nanoseconds()/1000/1000)
	return msSince > f.fullDurationMS+f.decayDurationMS
}

func (f FlashControl) color() Color{
	msSince := int(time.Since(f.start).Nanoseconds()/1000/1000)
	switch {
	case msSince < f.fullDurationMS:
		return f.c
	case msSince < f.fullDurationMS+f.decayDurationMS:
		scale := 1.0 - (float32(msSince)-float32(f.fullDurationMS)) / float32(f.decayDurationMS)
		return f.c.Scale(scale)
	default:
		return Black
	}
}

