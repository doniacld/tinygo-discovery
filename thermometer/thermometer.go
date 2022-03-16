package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

func main() {
	// configure the data pin
	pin := machine.D6
	dhtSensor := dht.New(pin, dht.DHT22)
	for {
		// call the method asking the sensor for the data
		temp, hum, err := dhtSensor.Measurements()
		if err != nil {
			fmt.Printf("Measurements failed: %s\n", err.Error())
		} else {
			// print data with current time
			now := time.Now()
			fmt.Printf("%02d:%02d:%02d, ", now.Hour(), now.Minute(), now.Second())
			fmt.Printf("Temperature: %02d.%d°C, ", temp/10, temp%10)
			fmt.Printf("Humidity: %02d.%d%%\n", hum/10, hum%10)
		}

		// measurements should be checked after 2 seconds with this sensor
		time.Sleep(time.Second * 2)
	}
}
