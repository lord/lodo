
package main

import "fmt"
import "os"
import "strconv"
import "time"

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readPin(pin int) int {
    ainName := fmt.Sprintf("/sys/devices/ocp.3/helper.15/AIN%d", pin)
    ainInput, err := os.Open(ainName)
    check(err)
	text := make([]byte, 10)
	count, err := ainInput.Read(text)
	check(err)
	s := string(text[0:count-1])
	val,err := strconv.Atoi(s)
	check(err)
    err = ainInput.Close();
    check(err)
    return val
}

var values = []int{0,0,0,0,0,0,0}

func ReadAIO() []int {
	for p:=0; p<7; p++ {
		values[p] = readPin(p)
	}
    return values
}

func setGPIO(pin int, value int) {
	pinName := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
    gpioOutput, err := os.Create(pinName)
    check(err)
    s := fmt.Sprintf("%d",value)
	count, err := gpioOutput.Write( []byte(s))
	check(err)
	count++
    err = gpioOutput.Close();
    check(err)
    return
}

func setChannel (channel int) {
	setGPIO(23, channel&1)
	setGPIO(47, channel&2)
	setGPIO(27, channel&4)
	setGPIO(22, channel&8)
}

func readSensor(bank int) []int {
	values := make([]int,16,16)
	for i:=0; i<16; i++{
		setChannel(i)
		values[i] = readPin(bank)
	}
	return values
}

func main() {
	setChannel(0);
    for i:=0; i<10000000; i++ {
    	values := readSensor(0)
    	fmt.Printf("%.4d\n", values)
		time.Sleep(1)
	}
}
