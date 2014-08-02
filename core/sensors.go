package core

// #cgo LDFLAGS:  -lpruio -L"/usr/local/lib/freebasic/linux/" -lfb -lpthread -lprussdrv -ltermcap -lsupc++
// #include "sensors.h"
import "C"
import "fmt"

//import "errors"

const sensorCount = 48

type Sensors struct {
	raw   [48]C.int
	last  [24]int
	net   [24]int
	pruio *C.struct_PruIo
}

const (
	up       = 0
	released = 1
	down     = 2
	pressed  = 3
)

var sensors Sensors

//var sensorMap = []int { 5,7,1,3,1,7,5,3,8,6,4,2,4,2,8,6,13,15,9, 11,9, 15,13,11,16,14,12,10,12,10,16,14,21,23,17,19,17,23,21,19,24,22,20,18,20,18,24,22}

func (sensors *Sensors) initSensors() error {
	sensors.pruio = C.pruio_new(0, 0x98, 0, 1)
	C.initSensors(sensors.pruio)
	for i := 0; i < rows*cols; i++ {
		sensors.last[i] = up
		sensors.net[i] = up
	}
	return nil
}

func (sensors *Sensors) stopSensors() error {
	C.stopSensors(sensors.pruio)
	return nil
}

func (sensors *Sensors) getBoardState(row int, col int) int {
	return sensors.net[row*4+col]
}

func (sensors *Sensors) readSensors() error {
	C.readSensors(sensors.pruio, &sensors.raw[0])
	return nil
}

func (sensors *Sensors) processSensors() error {
	var thd = C.int(16000)
	for i := 0; i < rows*cols; i++ {
		sensors.last[i] = sensors.net[i]
		sensors.net[i] = up
	}
	for bank := 0; bank < 3; bank++ {
		if sensors.raw[2+bank*16] > thd || sensors.raw[4+bank*16] > thd {
			sensors.net[0+bank*8] = down
		}
		if sensors.raw[11+bank*16] > thd || sensors.raw[13+bank*16] > thd {
			sensors.net[1+bank*8] = down
		}
		if sensors.raw[3+bank*16] > thd || sensors.raw[7+bank*16] > thd {
			sensors.net[2+bank*8] = down
		}
		if sensors.raw[10+bank*16] > thd || sensors.raw[12+bank*16] > thd {
			sensors.net[3+bank*8] = down
		}
		if sensors.raw[0+bank*16] > thd || sensors.raw[6+bank*16] > thd {
			sensors.net[4+bank*8] = down
		}
		if sensors.raw[9+bank*16] > thd || sensors.raw[15+bank*16] > thd {
			sensors.net[5+bank*8] = down
		}
		if sensors.raw[1+bank*16] > thd || sensors.raw[5+bank*16] > thd {
			sensors.net[6+bank*8] = down
		}
		if sensors.raw[8+bank*16] > thd || sensors.raw[14+bank*16] > thd {
			sensors.net[7+bank*8] = down
		}
	}
	for i := 0; i < rows*cols; i++ {
		if sensors.net[i] == down && (sensors.last[i] == up || sensors.last[i] == released) {
			sensors.net[i] = pressed
		}
		if sensors.net[i] == up && (sensors.last[i] == down || sensors.last[i] == pressed) {
			sensors.net[i] = released
		}
	}
	return nil
}

func (sensors *Sensors) printSensors() error {
	for bank := 0; bank < 3; bank++ {
		max := 0
		for ch := 0; ch < 16; ch++ {
			if max < int(sensors.raw[bank*16+ch]) {
				max = int(sensors.raw[bank*16+ch])
			}
			if sensors.raw[bank*16+ch] > 10000 {
				fmt.Printf("X ")
			} else {
				fmt.Printf("- ")
			}
			//fmt.Printf("%.5d ", int(sensors.raw[bank*16+ch]))
		}
		fmt.Printf(" : %d\n", max)
	}
	fmt.Printf("\n\n")
	return nil
}
