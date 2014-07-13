
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

func readPin(bank int) int {
//    time.Sleep(100 * time.Millisecond)
    ainName := fmt.Sprintf("/sys/devices/ocp.3/helper.15/AIN%d", bank)
    ainInput, err := os.OpenFile(ainName, os.O_RDONLY,0444)
    check(err)
	text := make([]byte, 10)
	count, err := ainInput.Read(text)
	time.Sleep(50 * time.Millisecond)
	count, err = ainInput.Read(text)
	if (err != nil) {
		count, err = ainInput.Read(text)
		check(err)
	}
	s := string(text[0:count-1])
	val,err := strconv.Atoi(s)
	check(err)
    err = ainInput.Close();
    check(err)
//    fmt.Printf("%d ", val)
//	time.Sleep(10 * time.Millisecond)
	return val
}

func readGPIO(pin int) int {
    gpioName := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
    gpioInput, err := os.Open(gpioName)
    check(err)
	text := make([]byte, 10)
	count, err := gpioInput.Read(text)
	if (err != nil) {
		count, err = gpioInput.Read(text)
		check(err)
	}
	_ = count
	s := string(text[0:count-1])
	val,err := strconv.Atoi(s)
	check(err)
    err = gpioInput.Close();
    check(err)
//    fmt.Printf("%d ", val)
    return val	
}

func setGPIO(pin int, value int) {
	pinName := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
    gpioOutput, err := os.Create(pinName)
    check(err)
    s := fmt.Sprintf("%d",value)
	count, err := gpioOutput.Write( []byte(s))
	check(err)
	_ = count
//	fmt.Printf("wrote: %d", count)
    err = gpioOutput.Close();
    check(err)
    return
}

func setChannel (channel int) {
//	fmt.Printf("Chan:%d",channel)
	setGPIO(23, channel&1)
	setGPIO(47, channel&2/2)
	setGPIO(27, channel&4/4)
	setGPIO(22, channel&8/8)
 //   fmt.Printf("%d -> %d %d %d %d ", channel, channel&1,channel&2/2,channel&4/4,channel&8/8)
 //   val23 := readGPIO(23)
 //   val47 := readGPIO(47)
 //   val27 := readGPIO(27)
 //   val22 := readGPIO(22)
//    fmt.Printf("(%d %d %d %d)\n ", val23,val47,val27,val22)
}

func readSensor(bank int) []int {
	values := make([]int,16,16)
	for i:=0; i<16; i++{
		setChannel(i)
		pValue := readPin(bank)
		pValue = readPin(bank)
		pValue = readPin(bank)
		pValue = readPin(bank)
		pValue = readPin(bank)
		pValue = readPin(bank)		
		pValue = readPin(bank)		
		pValue = readPin(bank)		
		values[i] = pValue
//   		fmt.Printf(" Value: %d\n", pValue)		
	}
	return values
}

func printSensor(bank int, values []int){
	fmt.Printf("%d ", bank)
	for i:=0; i<16; i++ {
		if values[i] > 300 {
			fmt.Printf("X ")
		} else {
			fmt.Printf("- ")
		}
	}
	fmt.Printf("\n")
}

func main() {
    for i:=0; i<10000000; i++ {
    	values0 := readSensor(0)
    	values1 := readSensor(1)
    	printSensor(0,values0)
    	printSensor(1,values1)
    	fmt.Printf("\n")
//		time.Sleep(1000 * time.Millisecond)
	}
}
