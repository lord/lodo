package core

import (
	"fmt"
	"os/exec"
	//"bytes"
	"strings"
)

const BreakoutMusic = "/root/gopath/src/github.com/james/lodo/media/Breakout_music.mp3"
const RippleMusic = "/root/gopath/src/github.com/james/lodo/media/Ripple_Music.mp3"
const SitarMusic = "/root/gopath/src/github.com/james/lodo/media/Sitar.mp3"

const A1deep = "/root/gopath/src/github.com/james/lodo/media/ripple/A1_deep.wav"
const A2deep = "/root/gopath/src/github.com/james/lodo/media/ripple/A2_deep.wav"
const C1deep = "/root/gopath/src/github.com/james/lodo/media/ripple/C1_deep.wav"
const C2deep = "/root/gopath/src/github.com/james/lodo/media/ripple/C2_deep.wav"
const D1deep = "/root/gopath/src/github.com/james/lodo/media/ripple/D1_deep.wav"
const D2deep = "/root/gopath/src/github.com/james/lodo/media/ripple/D2_deep.wav"
const E1deep = "/root/gopath/src/github.com/james/lodo/media/ripple/E1_deep.wav"
const E2deep = "/root/gopath/src/github.com/james/lodo/media/ripple/E2_deep.wav"
const G1deep = "/root/gopath/src/github.com/james/lodo/media/ripple/G1_deep.wav"
const G2deep = "/root/gopath/src/github.com/james/lodo/media/ripple/G2_deep.wav"
const A1ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_A1.wav"
const A2ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_A2.wav"
const C1ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_C1.wav"
const C2ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_C2.wav"
const D1ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_D1.wav"
const D2ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_D2.wav"
const E1ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_E1.wav"
const E2ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_E2.wav"
const G1ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_G1.wav"
const G2ripple = "/root/gopath/src/github.com/james/lodo/media/ripple/ripple_G2.wav"

const Selectgame = "/root/gopath/src/github.com/james/lodo/media/voice/select_game.wav"
const Bounce1 = "/root/gopath/src/github.com/james/lodo/media/wav/bounce_1.wav"
const Bounce5 = "/root/gopath/src/github.com/james/lodo/media/wav/bounce_5.wav"
const Pong = "/root/gopath/src/github.com/james/lodo/media/wav/Pong.wav"
const Glass = "/root/gopath/src/github.com/james/lodo/media/wav/glass_breaking_1.wav"
const GameOver = "/root/gopath/src/github.com/james/lodo/media/wav/gameover.wav"
const Pewpewpew = "/root/gopath/src/github.com/james/lodo/media/wav/PewPewPew.wav"

const Rip1 = "/root/gopath/src/github.com/james/lodo/media/rip/1.wav"
const Rip2 = "/root/gopath/src/github.com/james/lodo/media/rip/2.wav"
const Rip3 = "/root/gopath/src/github.com/james/lodo/media/rip/3.wav"
const Rip4 = "/root/gopath/src/github.com/james/lodo/media/rip/4.wav"
const Rip5 = "/root/gopath/src/github.com/james/lodo/media/rip/5.wav"
const Rip6 = "/root/gopath/src/github.com/james/lodo/media/rip/6.wav"
const Rip7 = "/root/gopath/src/github.com/james/lodo/media/rip/7.wav"
const Rip8 = "/root/gopath/src/github.com/james/lodo/media/rip/8.wav"
const Rip9 = "/root/gopath/src/github.com/james/lodo/media/rip/9.wav"
const Rip10 = "/root/gopath/src/github.com/james/lodo/media/rip/10.wav"

const Rip11 = "/root/gopath/src/github.com/james/lodo/media/rip/11.wav"
const Rip12 = "/root/gopath/src/github.com/james/lodo/media/rip/12.wav"
const Rip13 = "/root/gopath/src/github.com/james/lodo/media/rip/13.wav"
const Rip14 = "/root/gopath/src/github.com/james/lodo/media/rip/14.wav"
const Rip15 = "/root/gopath/src/github.com/james/lodo/media/rip/15.wav"
const Rip16 = "/root/gopath/src/github.com/james/lodo/media/rip/16.wav"
const Rip17 = "/root/gopath/src/github.com/james/lodo/media/rip/17.wav"
const Rip18 = "/root/gopath/src/github.com/james/lodo/media/rip/18.wav"
const Rip19 = "/root/gopath/src/github.com/james/lodo/media/rip/19.wav"
const Rip20 = "/root/gopath/src/github.com/james/lodo/media/rip/20.wav"

const Rip21 = "/root/gopath/src/github.com/james/lodo/media/rip/21.wav"
const Rip22 = "/root/gopath/src/github.com/james/lodo/media/rip/22.wav"
const Rip23 = "/root/gopath/src/github.com/james/lodo/media/rip/23.wav"
const Rip24 = "/root/gopath/src/github.com/james/lodo/media/rip/24.wav"
const Rip25 = "/root/gopath/src/github.com/james/lodo/media/rip/25.wav"
const Rip26 = "/root/gopath/src/github.com/james/lodo/media/rip/26.wav"
const Rip27 = "/root/gopath/src/github.com/james/lodo/media/rip/27.wav"
const Rip28 = "/root/gopath/src/github.com/james/lodo/media/rip/28.wav"
const Rip29 = "/root/gopath/src/github.com/james/lodo/media/rip/29.wav"
const Rip30 = "/root/gopath/src/github.com/james/lodo/media/rip/30.wav"


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
		s.cmd = exec.Command("/usr/bin/mpg321", "-l", "0", "-a", "default:CARD=Set", "-g", fmt.Sprintf("%d",s.volume), soundfile)	
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
    defer func() {
        if e := recover(); e != nil {
        }
    } ()
	s.cmd.Process.Kill()
}

func (s *Sound) Pause(){
	s.cmd.Process.Kill()
}

func (s *Sound) Resume(){
	s.cmd.Process.Kill()
}

func (s *Sound) Print(){
		fmt.Printf("PID: %d",s.cmd.Process.Pid)
}