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

/*
import (
	"bufio"
	"fmt"
	"machine"
	"runtime"
	"strings"
	"time"

	"tinygo.org/x/drivers/dht"
	"tinygo.org/x/drivers/net/http"
	"tinygo.org/x/drivers/wifinina"
)

// access point info
const (
	ssid = "myfree"
	pass = "l1erjdr2mv"

	// IP address of the server aka "hub". Replace with your own info.
	// Can specify a URL starting with http or https
	url = "http://192.168.1.127/measure"
)

// these are the default pins for the Arduino Nano33 IoT.
// change these to connect to a different UART or pins for the ESP8266/ESP32
var (
	// these are the default pins for the Arduino Nano33 IoT.
	spi = machine.NINA_SPI

	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor *wifinina.Device

	buf [0x46a]byte
	//	lastRequestTime time.Time
	//	conn            net.Conn
)

// measure holds the temperature and the humidity of the sensor
// NB: json fields are useless for the moment cause the type.Name() is unimplemented for the moment
type measure struct {
	Temp int16  `json:"temperature"`
	Hum  uint16 `json:"humidity"`
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
	// setup the device
	setup()
	http.SetBuf(buf[:])

	waitSerial()
	connectToAP()

	// configure the data pin
	pin := machine.D6
	dhtSensor := dht.New(pin, dht.DHT22)

	// get measurements and send them to the server
	cnt := 0
	for {
		runtime.GC()

		fmt.Printf("-------- %d --------\r\n", cnt)
		temp, hum, err := measurements(dhtSensor.(dht.DummyDevice))
		if err != nil {
			fmt.Printf("Measurements failed: %s\n", err.Error())
		}

		postMeasure(measure{temp, hum})

		cnt++
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

func postMeasure(m measure) {
	temperature := float32(m.Temp) / 10
	humidity := float32(m.Hum) / 10

	body := fmt.Sprintf(`{"temperature":%f,"humidity":%f}`, temperature, humidity)
	fmt.Println("body: ", body)

	resp, err := http.Post(url, "application/json", strings.NewReader(body))
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
	defer resp.Body.Close()
}

// measurements sends a command to get temperature and humidity to the sensor
// and returns the values
func measurements(dhtSensor measurable) (int16, uint16, error) {
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

	return temp, hum, err
}

type measurable interface {
	Measurements() (int16, uint16, error)
}
*/
