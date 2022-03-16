package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"machine"
	"time"
)

func main() {
	// setup the LED as an output
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		// turn the light on
		led.Low()
		time.Sleep(time.Millisecond * 500)

		// turn the light off
		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
