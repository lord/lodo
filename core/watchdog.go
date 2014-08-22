package core

import "os"
import "fmt"

var f *os.File
var watchdog = true
// starts the watchdog service - will restart if not written to every 60 seconds
func StartDog() {
	watchdog = true
	var err error
	f, err = os.Create("/dev/watchdog")
	if err != nil {
		fmt.Println("Watchdog Error:", err)
		os.Exit(1)
	}
	_, err = f.WriteString("X")
	if err != nil {
		fmt.Println("StartDog Write Error:", err)
		os.Exit(1)
	}
	// err = f.Sync()
	// if err != nil {
	// 	fmt.Println("StartDog sync Error:", err)
	// 	os.Exit(1)
	// }
}

func StopDog() {
	if watchdog {
		watchdog = false
		f.Close()
	}
}

func PetDog(){
	// fmt.Printf("PetDog\n")
	if watchdog {
		// fmt.Printf("Write X\n")
		var err error
		_, err = f.WriteString("X")
		if err != nil {
			fmt.Println("PetDog Write Error:", err)
			os.Exit(1)
		}
		// err = f.Sync()
		// if err != nil {
		// 	fmt.Println("PetDog Sync Error:", err)
		// 	os.Exit(1)
		// }
	}
}