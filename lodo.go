package main

// #cgo LDFLAGS: -lm
// #include "tclled.h"
import "C"
import "fmt"

func main() {
	device := C.open_device()
	fmt.Print("Device status: ")
	fmt.Println(device)

	if device <= 0 {
		fmt.Println("Device init failed.")
		return 1
	}
	fmt.Print("Device status: ")
	fmt.Println(device)

}
