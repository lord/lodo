package core

import (
	"os/exec"
//    "log"
    "bytes"
    "strings"
)

func PlayWave() {
        cmd := exec.Command("aplay", "-D","default:CARD=Set", "/root/Pong.wav")
        cmd.Stdin = strings.NewReader("some input")
        var out bytes.Buffer
        cmd.Stdout = &out
        go cmd.Run()
        // if err != nil {
        //          log.Fatal(err)
        // }
}
