package core

// #cgo LDFLAGS:  -lpruio -L"/usr/local/lib/freebasic/linux/" -lfb -lpthread -lprussdrv -ltermcap -lsupc++
// #include "sensors.h"
import "C"

//import "errors"

const sensorCount = 48

type Sensors struct {
	raw   [80]C.int
	last  [40]int
	net   [40]int
	pruio *C.struct_PruIo
	rows  int
	cols  int
}

const (
	up       = 0
	released = 1
	down     = 2
	pressed  = 3
)

var sensors Sensors

func (sensors *Sensors) initSensors(rows, cols int) error {
	sensors.rows = rows
	sensors.cols = cols
	sensors.pruio = C.pruio_new(0, 0x98, 0, 1)
	C.initSensors(sensors.pruio)
	for i := 0; i < sensors.rows*sensors.cols; i++ {
		sensors.last[i] = up
		sensors.net[i] = up
	}
	return nil
}

func (sensors *Sensors) stopSensors() error {
	C.stopSensors(sensors.pruio)
	return nil
}

// I think this is backwards, actually x,y, not y,x as the args would suggest?
func (sensors *Sensors) getBoardState(x int, y int) int {
	return sensors.net[y+x*8]
}

func (sensors *Sensors) readSensors() error {
	C.readSensors(sensors.pruio, &sensors.raw[0])
	return nil
}

func (sensors *Sensors) processSensors() error {
	var thd = C.int(20000)
	for i := 0; i < sensors.rows*sensors.cols; i++ {
		sensors.last[i] = sensors.net[i]
		sensors.net[i] = up
	}

	for bank := 0; bank < 5; bank++ {
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
		// fmt.Printf("|| %.5d %.5d %.5d %.5d || %.5d %.5d %.5d %.5d ||%.5d %.5d %.5d %.5d || %.5d %.5d %.5d %.5d ||  \n",
		//             sensors.raw[0+bank*16],sensors.raw[1+bank*16],sensors.raw[2+bank*16],sensors.raw[3+bank*16],
		//             sensors.raw[4+bank*16],sensors.raw[5+bank*16],sensors.raw[6+bank*16],sensors.raw[7+bank*16],
		//             sensors.raw[8+bank*16],sensors.raw[9+bank*16],sensors.raw[10+bank*16],sensors.raw[11+bank*16],
		//             sensors.raw[12+bank*16],sensors.raw[13+bank*16],sensors.raw[14+bank*16],sensors.raw[15+bank*16])
	}
	// fmt.Printf("\n")
	for i := 0; i < sensors.rows*sensors.cols; i++ {
		if sensors.net[i] == down && (sensors.last[i] == up || sensors.last[i] == released) {
			sensors.net[i] = pressed
		}
		if sensors.net[i] == up && (sensors.last[i] == down || sensors.last[i] == pressed) {
			sensors.net[i] = released
		}
	}
	return nil
}
