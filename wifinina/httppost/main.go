// This example opens a TCP connection using a device with WiFiNINA firmware
// and sends a HTTP request to retrieve a webpage, based on the following
// Arduino example:
//
// https://github.com/arduino-libraries/WiFiNINA/blob/master/examples/WiFiWebClientRepeating/
//
// This example will not work with samd21 or other systems with less than 32KB
// of RAM.  Use the following if you want to run wifinina on samd21, etc.
//
// examples/wifinina/webclient
// examples/wifinina/tlsclient
//
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"machine"
	"strings"
	"time"

	"tinygo.org/x/drivers/dht"
	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/net/http"
	"tinygo.org/x/drivers/wifinina"
)

// access point info
const (
	ssid = "myfree"
	pass = "xxx"

	// IP address of the server aka "hub". Replace with your own info.
	// Can specify a URL starting with http or https
	url = "http://192.168.0.10/measure"
)

// these are the default pins for the Arduino Nano33 IoT.
// change these to connect to a different UART or pins for the ESP8266/ESP32
var (
	// these are the default pins for the Arduino Nano33 IoT.
	spi = machine.NINA_SPI

	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor *wifinina.Device

	// dhtSensor is the DHT sensor to measure temperature and humidity
	//	dhtSensor dht.DummyDevice

	buf             [0x400]byte
	lastRequestTime time.Time
	conn            net.Conn
)

type Measure struct {
	Temp int16  `json:"temp"`
	Hum  uint16 `json:"hum"`
}

func setup() {
	// Configure SPI for 8Mhz, Mode 0, MSB First
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	adaptor = wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()
}

func main() {

	setup()
	http.SetBuf(buf[:])

	waitSerial()
	connectToAP()

	// configure the data pin
	pin := machine.D6
	dhtSensor := dht.New(pin, dht.DHT22)

	cnt := 0
	for {
		// call the method asking the sensor for the data
		temp, hum, err := dhtSensor.Measurements()
		if err != nil {
			fmt.Printf("Measurements failed: %s\n", err.Error())
		} else {
			// print data with current time
			now := time.Now()
			fmt.Printf("%02d:%02d:%02d, ", now.Hour(), now.Minute(), now.Second())
			// received data is times 10
			fmt.Printf("Temperature: %02d.%d°C, ", temp/10, temp%10)
			fmt.Printf("Humidity: %02d.%d%%\n", hum/10, hum%10)
		}

		// To test the connection, send a request to the server
		//body := `{"temp": 270, "hum": 900}`
		//resp, err := http.Post(url, "application/json", strings.NewReader(body))
		//if err != nil {
		//	println(err)
		//	continue
		//}

		// TODO make it work
		//measure := Measure{Temp: temp, Hum: hum}
		//body, err := json.Marshal(measure)
		//if err != nil {
		//	fmt.Printf("%s\r\n", err.Error())
		//	continue
		//}

		//fmt.Println("body: ", string(body))

		body := fmt.Sprintf(`{"temp":%d,"hum":%d}`, temp, hum)

		resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
		if err != nil {
			fmt.Printf("%s\r\n", err.Error())
			continue
		}

		fmt.Printf("%s %s\r\n", resp.Proto, resp.Status)
		for k, v := range resp.Header {
			fmt.Printf("%s: %s\r\n", k, strings.Join(v, " "))
		}
		fmt.Printf("\r\n")

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			fmt.Printf("%s\r\n", scanner.Text())
		}
		resp.Body.Close()

		cnt++
		fmt.Printf("-------- %d --------\r\n", cnt)
		time.Sleep(10 * time.Second)
	}

}

// Wait for user to open serial console
func waitSerial() {
	for !machine.Serial.DTR() {
		time.Sleep(100 * time.Millisecond)
	}
}

// connect to access point
func connectToAP() {
	time.Sleep(2 * time.Second)
	println("Connecting to " + ssid)
	err := adaptor.ConnectToAccessPoint(ssid, pass, 10*time.Second)
	if err != nil { // error connecting to AP
		for {
			println(err)
			time.Sleep(1 * time.Second)
		}
	}

	println("Connected.")

	ip, _, _, err := adaptor.GetIP()
	for ; err != nil; ip, _, _, err = adaptor.GetIP() {
		message(err.Error())
		time.Sleep(1 * time.Second)
	}
	message(ip.String())
}

func message(msg string) {
	println(msg, "\r")
}

func postMeasure(measure Measure) {
	body, err := json.Marshal(measure)
	if err != nil {
		println(err)
		return
	}

	fmt.Println("body: ", string(body))

	// To test the connection, send a request to the server
	// body := `{"temp": 270, "hum": 900}`
	// resp, err := http.Post(url, "application/json", strings.NewReader(body))

	resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		return
	}

	fmt.Printf("%s %s\r\n", resp.Proto, resp.Status)
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\r\n", k, strings.Join(v, " "))
	}
	fmt.Printf("\r\n")

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Printf("%s\r\n", scanner.Text())
	}
	resp.Body.Close()
}
