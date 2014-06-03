package main

// #cgo LDFLAGS: -lm
// #include "tclled.h"
import "C"
import "errors"

type Strand struct {
	device   C.int
	buffer   *_Ctype_tcl_buffer
	ledCount int
}

func (s *Strand) Connect(ledCount int) error {
	s.ledCount = ledCount
	s.device = C.open_device()

	if s.device <= 0 {
		return errors.New("Device init failed")
	}

	C.set_gamma(2.2, 2.2, 2.2)
	spiStatus := C.spi_init(s.device)
	if spiStatus != 0 {
		return errors.New("SPI init failed")
	}

	s.buffer = &C.tcl_buffer{}
	tclStatus := C.tcl_init(s.buffer, C.int(s.ledCount))
	if tclStatus != 0 {
		return errors.New("TCL init failed")
	}

	for i := 0; i < ledCount; i++ {
		s.SetColor(i, 0, 0, 0)
	}
	s.Save()

	return nil
}

func (s *Strand) Free() error {
	C.tcl_free(s.buffer)
	C.close_device(s.device)

	return nil
}

func (s *Strand) SetColor(ledNumber int, r int, g int, b int) {
	C.write_gamma_color_to_buffer(s.buffer, C.int(ledNumber), C.uint8_t(r), C.uint8_t(g), C.uint8_t(b))
}

func (s *Strand) Save() {
	C.send_buffer(s.device, s.buffer)
}
