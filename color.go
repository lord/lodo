package main

import "time"
import "math/rand"

type Color struct {
	R, G, B int
}

func MakeColor(r, g, b int) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

func randomColor() Color {
	t := time.Now().UnixNano()
	r := rand.New(rand.NewSource(t))
	g := rand.New(rand.NewSource(t + 1))
	b := rand.New(rand.NewSource(t + 2))
	return MakeColor(r.Intn(255), g.Intn(255), b.Intn(255))
}
