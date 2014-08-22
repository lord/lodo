package core

import (
	"fmt"
	"os/exec"
	//    "log"
	"bytes"
	"strings"
)

const BreakoutMusic = "/root/gopath/src/github.com/lord/lodo/media/Breakout_music.mp3"
const RippleMusic = "/root/gopath/src/github.com/lord/lodo/media/Ripple_Music.mp3"

const A1deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/A1_deep.wav"
const A2deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/A2_deep.wav"
const C1deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/C1_deep.wav"
const C2deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/C2_deep.wav"
const D1deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/D1_deep.wav"
const D2deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/D2_deep.wav"
const E1deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/E1_deep.wav"
const E2deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/E2_deep.wav"
const G1deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/G1_deep.wav"
const G2deep = "/root/gopath/src/github.com/lord/lodo/media/ripple/G2_deep.wav"
const A1ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_A1.wav"
const A2ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_A2.wav"
const C1ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_C1.wav"
const C2ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_C2.wav"
const D1ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_D1.wav"
const D2ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_D2.wav"
const E1ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_E1.wav"
const E2ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_E2.wav"
const G1ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_G1.wav"
const G2ripple = "/root/gopath/src/github.com/lord/lodo/media/ripple/ripple_G2.wav"

const Selectgame = "/root/gopath/src/github.com/lord/lodo/media/voice/select_game.wav"
const Bounce1 = "/root/gopath/src/github.com/lord/lodo/media/wav/bounce_1.wav"
const Bounce5 = "/root/gopath/src/github.com/lord/lodo/media/wav/bounce_5.wav"
const Pong = "/root/gopath/src/github.com/lord/lodo/media/wav/Pong.wav"
const Glass = "/root/gopath/src/github.com/lord/lodo/media/wav/glass_breaking_1.wav"

const wav = 1
const mpg = 2

type Sound struct{
 	cmd 		*exec.Cmd
 	volume   	int
 	filetype 	int
} 

func MakeSound(soundfile string) Sound {
	s := Sound{ volume: 40}
	if strings.Contains(soundfile, ".wav") {
		s.cmd = exec.Command("aplay", "-D", "default:CARD=Set", soundfile)	
		s.filetype = wav
	} else {
		s.cmd = exec.Command("mpg321", "-a", "default:CARD=Set", "-g", fmt.Sprintf("%d",s.volume), soundfile)	
		s.filetype = mpg
	}
	return s
}

func PlaySound(soundfile string) {
	s := MakeSound(soundfile)
	s.Play()
}

// plays sound, waits for completion
func (s *Sound) Run() {
	s.cmd.Run()
}

// play the sound, plays indefinitely
func (s *Sound) Play(){
	s.cmd.Start()
}

func (s *Sound) Stop(){
	s.cmd.Process.Kill()
}

func PlayWave() {
	cmd := exec.Command("aplay", "-D", "default:CARD=Set", "/root/Pong.wav")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	go cmd.Run()
}

