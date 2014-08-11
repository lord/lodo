package core

// #cgo LDFLAGS: -lm
// #include "tclled.h"
import "C"
import "errors"

type Strand struct {
	device   C.int
	cbuf     *C.struct__tcl_buffer
	buffer   []Color
	ledCount int
}

func (s *Strand) Length() int {
	return s.ledCount
}

func (s *Strand) Connect(ledCount int) error {
	s.ledCount = ledCount + 100
	s.device = C.open_device()
	s.buffer = make([]Color, s.ledCount)

	if s.device <= 0 {
		return errors.New("Device init failed")
	}

	spiStatus := C.spi_init(s.device)
	if spiStatus != 0 {
		return errors.New("SPI init failed")
	}

	s.cbuf = &C.struct__tcl_buffer{}
	tclStatus := C.tcl_init(s.cbuf, C.int(s.ledCount))
	if tclStatus != 0 {
		return errors.New("TCL init failed")
	}

	for i := 0; i < s.ledCount; i++ {
		s.SetColor(i, MakeColor(0, 0, 0))
	}
	s.Save()

	return nil
}

func (s *Strand) Free() error {
	C.tcl_free(s.cbuf)
	C.close_device(s.device)

	return nil
}

func (s *Strand) GetColor(ledNumber int) Color {
	return s.buffer[ledNumber]
}

func (s *Strand) SetColor(ledNumber int, c Color) {
	var color Color
	if c.A == 1 {
		color = c
	} else {
		color = s.buffer[ledNumber].AddAlphaColor(c)
	}
	s.buffer[ledNumber] = color
	// These colors are rotated so that they actually set the sensors correctly
	C.write_color_to_buffer(s.cbuf, C.int(ledNumber), C.uint8_t(color.G), C.uint8_t(color.R), C.uint8_t(color.B))
}

func (s *Strand) Save() {
	C.send_buffer(s.device, s.cbuf)
}
