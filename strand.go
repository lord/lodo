package main

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

func (s *Strand) Connect(ledCount int) error {
	s.ledCount = ledCount
	s.device = C.open_device()
	s.buffer = make([]Color, ledCount)

	if s.device <= 0 {
		return errors.New("Device init failed")
	}

	C.set_gamma(2.2, 2.2, 2.2)
	spiStatus := C.spi_init(s.device)
	if spiStatus != 0 {
		return errors.New("SPI init failed")
	}

	s.cbuf = &C.struct__tcl_buffer{}
	tclStatus := C.tcl_init(s.cbuf, C.int(s.ledCount))
	if tclStatus != 0 {
		return errors.New("TCL init failed")
	}

	for i := 0; i < ledCount; i++ {
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
	if c.A == 1 {
		s.buffer[ledNumber] = c
	} else {
		s.buffer[ledNumber] = s.buffer[ledNumber].AddAlphaColor(c)
	}
}

func (s *Strand) Save() {
	for i, c := range s.buffer {
		C.write_gamma_color_to_buffer(s.cbuf, C.int(i), C.uint8_t(c.R), C.uint8_t(c.G), C.uint8_t(c.B))
	}
	C.send_buffer(s.device, s.cbuf)
}
