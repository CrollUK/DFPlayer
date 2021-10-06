package main

import (
	"github.com/CrollUK/DFPlayer/DFPlayerMini"
	"machine"
	"time"

)


func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	DFPlayerMini.Setup()
	DFPlayerMini.Play()

	time.Sleep(2 * time.Second)
	for {
		led.Set(!led.Get())
		time.Sleep(time.Millisecond * 500)
	}
}