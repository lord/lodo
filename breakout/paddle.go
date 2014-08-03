package breakout

type paddle struct {
	x, y, w, h float64
}

func MakePaddle(x, y, w, h float64) paddle {
	return paddle{
		x: x,
		y: y,
		w: w,
		h: h,
	}
}
