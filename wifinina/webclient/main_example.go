// This example opens a TCP connection using a device with WiFiNINA firmware
// and sends a HTTP request to retrieve a webpage, based on the following
// Arduino example:
//
// https://github.com/arduino-libraries/WiFiNINA/blob/master/examples/WiFiWebClientRepeating/
//
package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/net"
	"tinygo.org/x/drivers/net/tls"
	"tinygo.org/x/drivers/wifinina"
)

// access point info
const (
	ssid = "myfree"
	pass = "l1erjdr2mv"

	// IP address of the server aka "hub". Replace with your own info.
	server = "192.168.1.127:80"
)

// these are the default pins for the Arduino Nano33 IoT.
// change these to connect to a different UART or pins for the ESP8266/ESP32
var (

	// these are the default pins for the Arduino Nano33 IoT.
	spi = machine.NINA_SPI

	// this is the ESP chip that has the WIFININA firmware flashed on it
	adaptor *wifinina.Device
)

var buf [256]byte

var lastRequestTime time.Time

//var conn net.Conn
var conn *net.TCPSerialConn

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

	time.Sleep(10 * time.Second)
	for {
		readConnection()
		makeHTTPRequest()
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
	if conn != nil {
		conn.Close()
	}

	//// make TCP connection
	//ip := net.ParseIP(server)
	//raddr := &net.TCPAddr{IP: ip, Port: 80}
	//laddr := &net.TCPAddr{Port: 8080}

	conn, err = tls.Dial("tcp", server, &tls.Config{})
	for err != nil {
		conn, err = tls.Dial("tcp", server, &tls.Config{})
		println("connection failed: " + err.Error())
		time.Sleep(30 * time.Second)
	}
	println("Connected!\r")

	//message("\r\n---------------\r\nDialing TCP connection")
	//conn, err = net.DialTCP("tcp", laddr, raddr)
	//for ; err != nil; conn, err = net.DialTCP("tcp", laddr, raddr) {
	//	message("Connection failed: " + err.Error())
	//	time.Sleep(5 * time.Second)
	//}
	//println("Connected!\r")

	print("Sending HTTP request...")
	fmt.Fprintln(conn, "GET /hi HTTP/1.1")
	fmt.Fprintln(conn, "Host:", server)
	fmt.Fprintln(conn, "User-Agent: TinyGo")
	fmt.Fprintln(conn, "Connection: close")
	fmt.Fprintln(conn)
	println("Sent!\r\n\r")

	lastRequestTime = time.Now()
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

	//ip, _, _, err := adaptor.GetIP()
	//for ; err != nil; ip, _, _, err = adaptor.GetIP() {
	//	message(err.Error())
	//	time.Sleep(1 * time.Second)
	//}
	//message(ip.String())
}

func message(msg string) {
	println(msg, "\r")
}
