// This example opens a TCP connection using a device with WiFiNINA firmware
// and sends a HTTP request to retrieve a webpage, based on the following
// Arduino example:
//
// https://github.com/arduino-libraries/WiFiNINA/blob/master/examples/WiFiWebClientRepeating/
//
package main

/*
import (
	"bufio"
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/wifinina"
)

// access point info
const (
	//	ssid = "NETGEAR0E635F"
	//	pass = "blacktuba577"

	ssid = "myfree"
	pass = "l1erjdr2mv"

	// IP address of the server aka "hub". Replace with your own info.
	server = "192.168.1.127"
)

// these are the default pins for the Arduino Nano33 IoT.
// change these to connect to a different UART or pins for the ESP8266/ESP32
var (
	// these are the default pins for the Arduino Nano33 IoT.
	spi = machine.NINA_SPI

	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor *wifinina.Device

	buf [256]byte

	lastRequestTime time.Time
	conn            net.Conn
)

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

	waitSerial()

	connectToAP()

	for {
		readConnection()
		if time.Now().Sub(lastRequestTime).Milliseconds() >= 10000 {
			makeHTTPRequest()
		}
	}

}

// Wait for user to open serial console
func waitSerial() {
	for !machine.Serial.DTR() {
		time.Sleep(100 * time.Millisecond)
	}
}

func readConnection() {
	if conn != nil {
		for n, err := conn.Read(buf[:]); n > 0; n, err = conn.Read(buf[:]) {
			if err != nil {
				println("Read error: " + err.Error())
			} else {
				print(string(buf[0:n]))
			}
		}
	}
}

func makeHTTPRequest() {

	var err error
	defer conn.Close()

	// make TCP connection
	ip := net.ParseIP(server)
	raddr := &net.TCPAddr{IP: ip}
	laddr := &net.TCPAddr{Port: 8080}

	message("\r\n---------------\r\nDialing TCP connection")
	for ; err != nil; conn, err = net.DialTCP("tcp", laddr, raddr) {
		message("Connection failed: " + err.Error())
		time.Sleep(10 * time.Second)
	}
	println("Connected!\r")

	println("Sending HTTP request...")

	request := "GET " + server + "/hi HTTP/1.1"
	_, err = fmt.Fprintln(conn, request)
	if err != nil {
		fmt.Print("Message from server: " + err.Error())
		return
	}

	lastRequestTime = time.Now()
	println("Sent!\r\n\r")

	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + msg)

	fmt.Fprintln(conn, "GET /hi HTTP/1.1")
	fmt.Fprintln(conn, "Host:", server)
	fmt.Fprintln(conn, "User-Agent: TinyGo")
	fmt.Fprintln(conn, "Connection: close")
	fmt.Fprintln(conn)
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

*/
