package main

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
