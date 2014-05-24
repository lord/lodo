package main

// #cgo LDFLAGS: -lm
// #include "tclled.h"
import "C"
import "fmt"
import "time"

func setColor(buffer *_Ctype_tcl_buffer, led_number int, r int, g int, b int) {
	C.write_gamma_color_to_buffer(buffer, C.int(led_number), C.uint8_t(r), C.uint8_t(g), C.uint8_t(b))
}

func sendBuffer(device C.int, buffer *_Ctype_tcl_buffer) {
	C.send_buffer(device, buffer)
}

func main() {
	device := C.open_device()
	fmt.Print("Device status: ")
	fmt.Println(device)

	if device <= 0 {
		fmt.Println("Device init failed.")
		return
	}

	C.set_gamma(2.2, 2.2, 2.2)

	spi_status := C.spi_init(device)
	fmt.Print("SPI status: ")
	fmt.Println(spi_status)

	if spi_status != 0 {
		fmt.Println("SPI init failed.")
		return
	}

	buffer := &C.tcl_buffer{}
	tcl_status := C.tcl_init(buffer, 30)
	fmt.Print("TCL status: ")
	fmt.Println(tcl_status)

	if tcl_status != 0 {
		fmt.Println("TCL init failed.")
		return
	}

	color := 0
	goingDown := false
	for true {
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

		for i := 0; i < 30; i++ {
			setColor(buffer, i, color, color, color)
		}
		sendBuffer(device, buffer)
		time.Sleep(10 * time.Millisecond)
	}
}
