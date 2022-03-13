package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

func main() {
	// Configure the data pin.
	pin := machine.D6
	dhtSensor := dht.New(pin, dht.DHT22)
	for {
		// Call the method asking the captor for the data.
		temp, hum, err := dhtSensor.Measurements()
		if err != nil {
			// error case
			fmt.Printf("Measurements failed: %s\n", err.Error())
		} else {
			// print result
			now := time.Now()
			fmt.Printf("%02d:%02d, ", now.Minute(), now.Second())
			fmt.Printf("Temperature: %02d.%d°C, ", temp/10, temp%10)
			fmt.Printf("Humidity: %02d.%d%%\n", hum/10, hum%10)
		}

		// Measurements cannot be updated only 2 seconds.
		// More frequent calls will return the same value.
		time.Sleep(time.Second * 2)
	}
}
