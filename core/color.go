package core

import "time"
import "math/rand"

type Color struct {
	R, G, B int
	A       float64
}

var red = MakeColor(200, 0, 0)
var blue = MakeColor(0, 200, 0)

func MakeColor(r, g, b int) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: 1,
	}
}

func MakeColorAlpha(r, g, b int, a float64) Color {
	return Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func (c Color) AddAlphaColor(c2 Color) Color {
	c.R = int(float64(c2.R-c.R)*c2.A) + c.R
	c.G = int(float64(c2.G-c.G)*c2.A) + c.G
	c.B = int(float64(c2.B-c.B)*c2.A) + c.B
	return c
}

func (c Color) WithAlpha(a float64) Color {
	c.A = a
	return c
}

func (c Color) Scale(amt float32) Color {
	return MakeColor(
		int(amt*float32(c.R)),
		int(amt*float32(c.G)),
		int(amt*float32(c.B)),
	)
}

func randomColor() Color {
	t := time.Now().UnixNano()
	r := rand.New(rand.NewSource(t))
	g := rand.New(rand.NewSource(t + 1))
	b := rand.New(rand.NewSource(t + 2))
	return MakeColor(r.Intn(255), g.Intn(255), b.Intn(255))
}
