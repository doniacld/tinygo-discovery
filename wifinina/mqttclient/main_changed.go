// This is a sensor station that uses a ESP8266 or ESP32 running on the device UART1.
// It creates an MQTT connection that publishes a message every second
// to an MQTT broker.
//
// In other words:
// Your computer <--> UART0 <--> MCU <--> UART1 <--> ESP8266 <--> Internet <--> MQTT broker.
//
// You must install the Paho MQTT package to build this program:
//
// 		go get -u github.com/eclipse/paho.mqtt.golang
//
package main

/*
import (
	"machine"
	"math/rand"
	"time"

	"tinygo.org/x/drivers/net/mqtt"
	"tinygo.org/x/drivers/wifinina"
)

const (
	// access point info
	ssid = "myfree"
	pass = "l1erjdr2mv"

	//ssid = "Donia's iPhone"
	//pass = "dearsugar"

	// IP address of the MQTT broker to use. Replace with your own info.
	//server = "tcp://test.mosquitto.org:1883"
	server = "tcp://192.168.0.26:1883"
	//server = "ssl://test.mosquitto.org:8883"

	// topic to publish to
	topic = "tinygo"
)

// these are the default pins for the Arduino Nano33 IoT.
// change these to connect to a different UART or pins for the ESP8266/ESP32
var (
	// these are the default pins for the Arduino Nano33 IoT.
	spi = machine.NINA_SPI

	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor *wifinina.Device
)

func main() {
	time.Sleep(3000 * time.Millisecond)
	println("Let's begin")

	rand.Seed(time.Now().UnixNano())

	// Configure SPI for 8Mhz, Mode 0, MSB First
	spi.Configure(machine.SPIConfig{
		Frequency: 8 * 1e6,
		SDO:       machine.NINA_SDO,
		SDI:       machine.NINA_SDI,
		SCK:       machine.NINA_SCK,
	})

	// Init esp8266/esp32
	adaptor = wifinina.New(spi,
		machine.NINA_CS,
		machine.NINA_ACK,
		machine.NINA_GPIO0,
		machine.NINA_RESETN)
	adaptor.Configure()

	connectToAP()

	opts := mqtt.NewClientOptions()
	opts = opts.AddBroker(server).SetClientID("tinygo-client-" + randomString(10))

	println("Connecting to MQTT...")
	cl := mqtt.NewClient(opts)
	println("We have a new client")

	token := cl.Connect()
	if token.Wait() && token.Error() != nil {
		failMessage(token.Error().Error())
	}

	println("Connected...")

	for i := 0; ; i++ {
		println("Publishing MQTT message...")

		payload := "Hello World!"
		//data := []byte(fmt.Sprintf(`{"e":[{"n":"hello %d","v":101}]}`, i))
		token = cl.Publish(topic, 0, false, payload)
		println(payload, token)
		if err := token.Error(); err != nil {
			switch t := err.(type) {
			case wifinina.Error:
				println(t.Error(), "attempting to reconnect")
				if token = cl.Connect(); token.Wait() && token.Error() != nil {
					failMessage(token.Error().Error())
				}
			default:
				println(err.Error())
			}
		}
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
		println(err.Error())
		time.Sleep(1 * time.Second)
	}
	println(ip.String())
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func failMessage(msg string) {
	for {
		println(msg)
		time.Sleep(1 * time.Second)
	}
}


*/
