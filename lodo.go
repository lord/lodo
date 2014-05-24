package main

// #cgo LDFLAGS: -lm
// #include "tclled.h"
import "C"
import "fmt"
import "time"
import "errors"

// import "time"

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

	return nil
}

func (s *Strand) SetColor(ledNumber int, r int, g int, b int) {
	C.write_gamma_color_to_buffer(s.buffer, C.int(ledNumber), C.uint8_t(r), C.uint8_t(g), C.uint8_t(b))
}

func (s *Strand) Save() {
	C.send_buffer(s.device, s.buffer)
}

func main() {
	strand := Strand{}
	ledCount := 30
	err := strand.Connect(ledCount)

	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
	}

	color := 0
	goingDown := false
	for true {
	for j := 0; j < 1000; j++ {
		if goingDown {
			color--
		} else {
			color++
		}
		if color < 0 {
			goingDown = false
			color = 0
		} else if color > 255 {
			goingDown = true
			color = 255
		}

		for i := 0; i < ledCount; i++ {
			strand.SetColor(i, color, color, color)
		}
		strand.Save()
		time.Sleep(10 * time.Millisecond)
	}
}
